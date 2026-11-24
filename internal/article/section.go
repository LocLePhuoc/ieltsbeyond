package article

// determines if section should be returned to web
// in markdown file, "## Ghi chú {#notes .internal}"
type Visibility string

const (
	VisibilityPublic   Visibility = "public"
	VisibilityInternal Visibility = "internal"
)

const ClassInternal = "internal"

// section extracted from file, not yet processed
type RawSection struct {
	Key        string     // ID of heading 2: "## Tiêu đề {#key}"
	Title      string     // heading text
	Classes    []string   // heading classes
	Body       string     // markdown below heading, until the next heading
	Visibility Visibility // inferred from classes
}

// processed Section
type Section struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Title string `json:"title"`
	HTML  string `json:"html,omitempty"`
	Data  any    `json:"data,omitempty"`
}
