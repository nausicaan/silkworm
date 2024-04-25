package main

// Links builds a collection of urls to target changlogs
type Links struct {
	ACF       string `json:"acf"`
	Calendar  string `json:"calendar"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	Spotlight string `json:"spotlight"`
	Tickets   string `json:"tickets"`
	Virtual   string `json:"virtual"`
	WordPress string `json:"wordpress"`
	WPExport  string `json:"wpexport"`
}

// Atlassian builds a list of jira tokens and api addresses
type Atlassian struct {
	Base   string `json:"base"`
	Token  string `json:"token"`
	Source string `json:"source"`
}

// Filters builds the parameters for sed to execute on the scrapped.txt file
type Filters struct {
	OPH2  string `json:"oph2"`
	OPH3  string `json:"oph3"`
	OPH4  string `json:"oph4"`
	CLH2  string `json:"clh2"`
	CLH3  string `json:"clh3"`
	CLH4  string `json:"clh4"`
	End   string `json:"end"`
	ESP   string `json:"esp"`
	Event string `json:"event"`
}

// Desso holds the value of the DESSO-XXXX identifier
type Desso struct {
	Issues []Issue `json:"issues"`
}

// Issue is a sub-structure of Desso
type Issue struct {
	Key string `json:"key"`
}

// Post contains the JSON parameters for a new Jira ticket
type Post struct {
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
}

const (
	header string = "\nh2. Changelog\n"
	bv     string = "2.0"
	halt   string = "program halted"
)

var (
	content   []byte
	label     string
	repo      string
	version   string
	filter    Filters
	jira      Atlassian
	link      Links
	post      Post
	title     Desso
	versions  = [1][2]string{{".", "-"}}
	common    = hmdr + "/Documents/common/"
	self      = hmdr + "/Github/silkworm/"
	jsons     = []string{self + "jsons/body.json", self + "jsons/filters.json", self + "jsons/links.json", self + "jsons/jira.json"}
	temp      = []string{common + "temp/grep.txt", common + "temp/scrape.txt"}
	deletions = []string{
		"<br />", "</h1>", "</h2>",
		"</h3>", "</h4>", "</li>",
		"<ul>", "</ul>", "</div>",
		"</div>", "<p>", "</p>",
		"<span>", "<entry>", "</entry>",
		"</span>", "<footer>", "</footer>",
		"<header>", "</header>",
	}
	replacements = [12][2]string{
		{"<h1>", "h1. "},
		{"<h2>", "h2. "},
		{"<h3>", "h3. "},
		{"<h4>", "h3. "},
		{"<li>", "- "},
		{"<strong>", "*"},
		{"</strong>", "*"},
		{"<em>", "*"},
		{"</em>", "*"},
		{"<code>", "*"},
		{"</code>", "*"},
		{"<li class=\"free\">", "- "},
	}
)
