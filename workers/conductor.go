package workers

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			json.Unmarshal([]byte(data), &post)
		case 1:
			json.Unmarshal([]byte(data), &filter)
		case 2:
			json.Unmarshal([]byte(data), &link)
		case 3:
			json.Unmarshal([]byte(data), &secret)
		}
	}
}

// Iterate through the Args array and assign plugin and ticket values
func sifter() {
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) > 25 {
			firstsplit := strings.Split(os.Args[i], "/")
			repo := firstsplit[0]
			secondsplit := strings.Split(firstsplit[1], ":")
			label := secondsplit[0]
			version = secondsplit[1]

			sorter(repo, label)
			changelog := append([]byte(header), content...)

			/* TODO Create Jira ticket using Description & Summary */
			post.Issues[0].Fields.Description = os.Args[1]
			post.Issues[0].Fields.Summary = string(changelog)
			// body, _ := json.Marshal(post)
			// response := capture("curl", "-D-", "-X", "POST", "-d", string(body), "-H", "Authorization: Bearer "+secret.Token, "-H", "Content-Type: application/json", secret.Issue)
			fmt.Println(string(changelog))
		}
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
		finder(link.Cal+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "event-tickets-plus":
		finder(link.Tickets+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "gravityforms":
		finder(link.Gravity, filter.OPH3+version+filter.End)
	case "polylang-pro":
		finder(link.Poly, filter.OPH4+version+filter.End)
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
	content = capture("sed", "/^$/d ; s/	//g", local+grepped)
	document(local+grepped, content)
}

func eventfilter() {
	content = capture("grep", "-v", "<", local+grepped)
	document(local+grepped, content)
	content = capture("sed", "1,3d", local+grepped)
	content = append([]byte("h3. "+version+"\n"), content...)
}
