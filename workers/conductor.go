package workers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var content []byte

// Quarterback is in charge of directing the program
func Quarterback() {
	jsoner()
	sifter()
}

func jsoner() {
	// Read the JSON file
	data, err := os.ReadFile("defaults/vars.json")
	inspect(err)

	// Unmarshal the JSON data into the struct
	json.Unmarshal([]byte(data), &filter)
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
		description := append([]byte("h2. Changelog\n"), content...)
		/* TODO Create Jira ticket using description */
		fmt.Println(string(description))
	}
}

// Sort the query based on the repo name
func sorter(repo, label string) {
	switch repo {
	case "bcgov-plugin":
		premium(label)
	case "freemius":
		finder(filter.WordPress+label+"/", "/Changelog"+filter.Spotlight)
	case "wpengine":
		finder(filter.ACF, "/Changelog"+filter.CLH1)
	default:
		finder(filter.WordPress+label+"/", "/Changelog"+filter.CLH2)
		content = capture("sed", "1d", grepped)
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
		content = capture("sed", "1,4d", grepped)
	case "polylang-pro":
		finder(filter.Poly, filter.OPH4+version+filter.End)
	case "event-tickets-plus":
		finder(filter.Tickets+string(v)+"/", "/"+version+filter.Special)
	case "wp-all-export-pro":
		finder(filter.WPExport, "/"+version+filter.CLH4)
		content = capture("sed", "${/h3./d;}", grepped)
	}
}

// Find and replace/delete html tags
func finder(link, filter string) {
	execute("curl", "-s", link, "-o", scraped)
	grep := capture("sed", "-n", filter, scraped)
	for _, v := range deletions {
		r := bytes.ReplaceAll(grep, []byte(v), []byte(""))
		grep = r
	}
	for i := 0; i < len(replacements); i++ {
		r := bytes.ReplaceAll(grep, []byte(replacements[i][0]), []byte(replacements[i][1]))
		grep = r
	}
	scribe(grepped, grep)
	content = capture("sed", "/^$/d", grepped)
	scribe(grepped, content)
}
