package workers

// Links builds
type Links struct {
	Cal       string `json:"calendar"`
	Tickets   string `json:"tickets"`
	ACF       string `json:"acf"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	WordPress string `json:"wordpress"`
	WPExport  string `json:"wpexport"`
}

// Secrets builds
type Secrets struct {
	Issue   string `json:"issue"`
	Token   string `json:"token"`
	Myself  string `json:"myself"`
	MyselfA string `json:"myselfA"`
}

// Filters builds the parameters for sed to execute on the scrapped.txt file
type Filters struct {
	OPH1      string `json:"oph1"`
	OPH2      string `json:"oph2"`
	OPH3      string `json:"oph3"`
	OPH4      string `json:"oph4"`
	CLH1      string `json:"clh1"`
	CLH2      string `json:"clh2"`
	CLH3      string `json:"clh3"`
	CLH4      string `json:"clh4"`
	End       string `json:"end"`
	Event     string `json:"event"`
	Spotlight string `json:"spotlight"`
}

// Post contains the JSON parameters for a new Jira ticket
type Post struct {
	Issues []struct {
		Fields struct {
			Assignee struct {
				Key string `json:"key"`
			} `json:"assignee"`
			Issuetype struct {
				ID string `json:"id"`
			} `json:"issuetype"`
			Creator struct {
				Key string `json:"key"`
			} `json:"creator"`
			Labels   []string `json:"labels"`
			Reporter struct {
				Key string `json:"key"`
			} `json:"reporter"`
			Project struct {
				ID  string `json:"id"`
				Key string `json:"key"`
			} `json:"project"`
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}

const (
	scraped, grepped, header string = "temp/scrape.txt", "temp/grep.txt", "h2. Changelog\n"
)

var (
	deletions    = []string{"<br />", "</h1>", "</h2>", "</h3>", "</h4>", "</li>", "<ul>", "</ul>", "<code>", "</code>", "</div>", "</div>", "<p>", "</p>", "<span>", "</span>", "<footer>", "</footer>", "<header>", "</header>"}
	jsons        = []string{local + "defaults/body.json", local + "defaults/filters.json", local + "defaults/links.json", local + "defaults/secret.json"}
	replacements = [9][2]string{
		{"<h1>", "h1. "},
		{"<h2>", "h2. "},
		{"<h3>", "h3. "},
		{"<h4>", "h3. "},
		{"<li>", "- "},
		{"<strong>", "*"},
		{"</strong>", "*"},
		{"<em>", "**"},
		{"</em>", "**"},
	}
	local    = hd + "/Documents/github/silkworm/"
	versions = [1][2]string{{".", "-"}}
	content  []byte
	version  string
	filter   Filters
	post     Post
	link     Links
	secret   Secrets
)
