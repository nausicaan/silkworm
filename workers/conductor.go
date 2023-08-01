package workers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var hd, _ = os.UserHomeDir()

// Quarterback is in charge of directing the program
func Quarterback() {
	jsoner()
	sifter()
}

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func jsoner() {
	for _, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch element {
		case "defaults/body.json":
			json.Unmarshal([]byte(data), &post)
		case "defaults/filters.json":
			json.Unmarshal([]byte(data), &filter)
		case "defaults/links.json":
			json.Unmarshal([]byte(data), &link)
		case "defaults/secret.json":
			json.Unmarshal([]byte(data), &secret)
		}
	}
}

// Iterate through the Args array and assign plugin and ticket values
func sifter() {
	for i := 1; i < len(os.Args); i++ {
		firstsplit := strings.Split(os.Args[i], "/")
		repo := firstsplit[0]
		secondsplit := strings.Split(firstsplit[1], ":")
		label := secondsplit[0]
		version = secondsplit[1]

		sorter(repo, label)
		changelog := append([]byte("h2. Changelog\n"), content...)

		/* TODO Create Jira ticket using Description & Summary */
		post.Issues[0].Fields.Description = os.Args[1]
		post.Issues[0].Fields.Summary = string(changelog)
		// body, _ := json.Marshal(post)
		// jira(body)

		fmt.Println(string(changelog))
	}
}

// Sort the query based on repository name
func sorter(repo, label string) {
	switch repo {
	case "bcgov-plugin":
		premium(label)
	case "freemius":
		finder(link.WordPress+"spotlight-social-photo-feeds/#developers", "/Changelog"+filter.Spotlight)
		content = capture("sed", "1d", local+grepped)
	case "wpengine":
		finder(link.ACF, "/Changelog"+filter.CLH1)
	default:
		finder(link.WordPress+label+"/#developers", "/Changelog"+filter.CLH2)
		content = capture("sed", "1d", local+grepped)
	}
}

// Apply special conditions to the premium in-house plugins
func premium(label string) {
	v := bytes.ReplaceAll([]byte(version), []byte(versions[0][0]), []byte(versions[0][1]))
	switch label {
	case "events-calendar-pro":
		finder(link.Cal+string(v)+"/", "/"+version+filter.CLH2)
	case "gravityforms":
		finder(link.Gravity, filter.OPH3+version+filter.End)
	case "polylang-pro":
		finder(link.Poly, filter.OPH4+version+filter.End)
	case "event-tickets-plus":
		finder(link.Tickets+string(v)+"/", "/"+version+filter.Special)
	case "wp-all-export-pro":
		finder(link.WPExport, "/"+version+filter.CLH4)
		content = capture("sed", "${/h3./d;}", local+grepped)
	}
}

// Find and replace/delete html tags
func finder(link, filter string) {
	execute("curl", "-s", link, "-o", local+scraped)
	grep := capture("sed", "-n", filter, local+scraped)
	for _, v := range deletions {
		replace := bytes.ReplaceAll(grep, []byte(v), []byte(""))
		grep = replace
	}
	for i := 0; i < len(replacements); i++ {
		replace := bytes.ReplaceAll(grep, []byte(replacements[i][0]), []byte(replacements[i][1]))
		grep = replace
	}
	document(local+grepped, grep)
	content = capture("sed", "/^$/d", local+grepped)
	document(local+grepped, content)
	content = capture("sed", "s/	//g", local+grepped)
	document(local+grepped, content)
}

func jira(body []byte) {
	posturl := secret.API
	request, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	inspect(err)

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	inspect(err)

	defer response.Body.Close()

	// derr := json.NewDecoder(response.Body).Decode(post)
	// inspect(derr)

	if response.StatusCode != http.StatusCreated {
		panic(response.Status)
	}
}
