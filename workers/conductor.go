package workers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	hd, _ = os.UserHomeDir()
	local = hd + "/Documents/github/silkworm/"
)

// Quarterback is in charge of directing the program
func Quarterback() {
	jsoner("defaults/body.json")
	jsoner("defaults/vars.json")
	sifter()
}

// Read the vars.json file and Unmarshal the data into a go structure
func jsoner(target string) {
	data, err := os.ReadFile(local + target)
	inspect(err)
	if strings.Contains(target, "body") {
		json.Unmarshal([]byte(data), &post)
	} else {
		json.Unmarshal([]byte(data), &filter)
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

		fmt.Println(string(changelog))
	}
}

// Sort the query based on the repo name
func sorter(repo, label string) {
	switch repo {
	case "bcgov-plugin":
		premium(label)
	case "freemius":
		finder(filter.WordPress+"spotlight-social-photo-feeds/#developers", "/Changelog"+filter.Spotlight)
		content = capture("sed", "1d", local+grepped)
	case "wpengine":
		finder(filter.ACF, "/Changelog"+filter.CLH1)
	default:
		finder(filter.WordPress+label+"/#developers", "/Changelog"+filter.CLH2)
		content = capture("sed", "1d", local+grepped)
	}
}

// Apply special conditions to the premium in-house plugins
func premium(label string) {
	v := bytes.ReplaceAll([]byte(version), []byte(versions[0][0]), []byte(versions[0][1]))
	switch label {
	case "events-calendar-pro":
		finder(filter.Cal+string(v)+"/", "/"+version+filter.CLH2)
	case "gravityforms":
		finder(filter.Gravity, filter.OPH3+version+filter.End)
	case "polylang-pro":
		finder(filter.Poly, filter.OPH4+version+filter.End)
	case "event-tickets-plus":
		finder(filter.Tickets+string(v)+"/", "/"+version+filter.Special)
	case "wp-all-export-pro":
		finder(filter.WPExport, "/"+version+filter.CLH4)
		content = capture("sed", "${/h3./d;}", local+grepped)
	}
}

// Find and replace/delete html tags
func finder(link, filter string) {
	execute("curl", "-s", link, "-o", local+scraped)
	grep := capture("sed", "-n", filter, local+scraped)
	for _, v := range deletions {
		r := bytes.ReplaceAll(grep, []byte(v), []byte(""))
		grep = r
	}
	for i := 0; i < len(replacements); i++ {
		r := bytes.ReplaceAll(grep, []byte(replacements[i][0]), []byte(replacements[i][1]))
		grep = r
	}
	document(local+grepped, grep)
	content = capture("sed", "/^$/d", local+grepped)
	document(local+grepped, content)
}
