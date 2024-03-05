package md

import (
	"fmt"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// Renderable defines renerable items
type Renderable interface {
	String() string
}

// Column formats items in a column
// i.e items just separated by newlines
type Column []Renderable

// Add adds renderable item to column
// Returns itself for composability
func (c *Column) Add(item ...Renderable) Renderable {
	*c = append(*c, item...)
	return c
}

// Row formats items in a row
// i.e items are a column in a table
type Row []Renderable

// Add adds renderable item to row
// Returns itself for composability
func (r *Row) Add(item ...Renderable) Renderable {
	*r = append(*r, item...)
	return r
}

// HeaderRow formats items in a row with header row
// i.e items are a column in a table
type HeaderRow struct {
	Headers []string
	Columns []Renderable
}

// NewHeaderRow creates new header row with given headers
func NewHeaderRow(headers ...string) *HeaderRow {
	return &HeaderRow{
		Headers: headers,
	}
}

// Add adds renderable item to row
// Returns itself for composability
func (h *HeaderRow) Add(item ...Renderable) Renderable {
	h.Columns = append(h.Columns, item...)
	return h
}

// Text basic text passthrough
// Note: does not add newline
type Text string

// H1 #
type H1 string

// H2 ##
type H2 string

// H3 ###
type H3 string

// Bold ** **
type Bold string

// Italic * *
type Italic string

// Strikethrough ~~ ~~
type Strikethrough string

// Divider ---
const Divider = Text("---\n")

// Blockquote >
// Note attempts to split the given string by newlines for
// multiline block quotes
type BlockQuote string

// OrderedList
// 1.
// 2.
// 3.
type OrderedList []string

// List
// -
// -
// -
type List []string

// Code ` `
type Code string

// CodeBlock
// ```
//
// ```
type CodeBlock string

// Table
// | Field | Value |
// | ----- | ----- |
// |       |       |
// |       |       |
type Table struct {
	Columns []string
	Rows    [][]string
}

// NewTable creates new table with given columns
func NewTable(columns ...string) *Table {
	return &Table{
		Columns: columns,
	}
}

// Add adds rows to table
// Returns itself for composability
func (t *Table) Add(row ...[]string) Renderable {
	t.Rows = append(t.Rows, row...)
	return t
}

// AddAny adds rows to table
// Calls fmt.Sprint on each item in each row to convert to string
// Returns itself for composability
func (t *Table) AddAny(row ...[]any) Renderable {
	var converted [][]string
	for _, r := range row {
		convertedRow := make([]string, 0, len(r))
		for _, item := range r {
			convertedRow = append(convertedRow, fmt.Sprint(item))
		}
		converted = append(converted, convertedRow)
	}

	t.Rows = append(t.Rows, converted...)
	return t
}

// AddRow adds items as row to table
// Returns itself for composability
func (t *Table) AddRow(items ...string) Renderable {
	t.Rows = append(t.Rows, items)
	return t
}

// AddRowAny adds items as row to table
// Calls fmt.Sprint on each item to convert to string
// Returns itself for composability
func (t *Table) AddRowAny(items ...any) Renderable {
	row := make([]string, 0, len(items))
	for _, item := range items {
		row = append(row, fmt.Sprint(item))
	}
	t.Rows = append(t.Rows, row)
	return t
}

// ToHTML converts given markdown to html string
func ToHTML(input string) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(input))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}
