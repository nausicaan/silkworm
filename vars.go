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
	Creds  string `json:"creds"`
}

// Filters builds the parameters for sed to execute on the scrapped.txt file
type Filters struct {
	OPH1  string `json:"oph1"`
	OPH2  string `json:"oph2"`
	OPH3  string `json:"oph3"`
	OPH4  string `json:"oph4"`
	CLH1  string `json:"clh1"`
	CLH2  string `json:"clh2"`
	CLH3  string `json:"clh3"`
	CLH4  string `json:"clh4"`
	End   string `json:"end"`
	ESP   string `json:"esp"`
	Event string `json:"event"`
}

// Ticket does some shit
type Ticket struct {
	Issues []struct {
		Key string `json:"key"`
	} `json:"issues"`
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
	header string = "h2. Changelog\n\n"
	bv     string = "2.0"
	reset  string = "\033[0m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	red    string = "\033[41m"
	halt   string = "program halted"
)

var (
	deletions    = []string{"<br />", "</h1>", "</h2>", "</h3>", "</h4>", "</li>", "<ul>", "</ul>", "</div>", "</div>", "<p>", "</p>", "<span>", "<entry>", "</entry>", "</span>", "<footer>", "</footer>", "<header>", "</header>"}
	jsons        = []string{gitpath + "jsons/body.json", gitpath + "jsons/filters.json", gitpath + "jsons/links.json", gitpath + "jsons/jira.json"}
	temp         = []string{hmdr + "/grep.txt", hmdr + "/scrape.txt"}
	replacements = [11][2]string{
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
	}
	gitpath  = hmdr + "/Documents/github/silkworm/"
	versions = [1][2]string{{".", "-"}}
	content  []byte
	version  string
	repo     string
	label    string
	post     Post
	link     Links
	filter   Filters
	jira     Atlassian
	title    Ticket
)
