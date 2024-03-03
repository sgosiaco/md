package md

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
func NewHeaderRow(headers ...string) HeaderRow {
	return HeaderRow{
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
func NewTable(columns ...string) Table {
	return Table{
		Columns: columns,
	}
}

// Add adds row to table
// Returns itself for composability
func (t *Table) Add(row ...[]string) Renderable {
	t.Rows = append(t.Rows, row...)
	return t
}
