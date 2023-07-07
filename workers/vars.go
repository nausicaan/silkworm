package workers

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
	Cal       string `json:"calendar"`
	Tickets   string `json:"tickets"`
	ACF       string `json:"acf"`
	ECBase    string `json:"ecbase"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	WordPress string `json:"wordpress"`
	WPExport  string `json:"wpexport"`
	Special   string `json:"special"`
	Spotlight string `json:"spotlight"`
}

const (
	scraped, grepped, header string = "temp/scrape.txt", "temp/grep.txt", "h2. Changelog\n"
)

var (
	deletions    = []string{"<br />", "</h1>", "</h2>", "</h3>", "</h4>", "</li>", "<ul>", "</ul>", "<code>", "</code>", "</div>", "</div>", "<p>", "</p>", "<span>", "</span>", "<footer>", "</footer>", "<header>", "</header>"}
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
	versions = [1][2]string{{".", "-"}}
	filter   Filters
	version  string
)
