package md

import (
	"fmt"
	"strings"
)

type Renderable interface {
	String() string
}

type Column []Renderable

func (d *Column) Add(item ...Renderable) {
	*d = append(*d, item...)
}

func (d Column) String() string {
	var sb strings.Builder

	for _, item := range d {
		sb.WriteString(item.String())
		// one newline for gap
		sb.WriteString("\n")
	}

	return sb.String()
}

type Row []Renderable

func (r *Row) Add(item ...Renderable) {
	*r = append(*r, item...)
}

func (r Row) String() string {
	var sb strings.Builder
	var header strings.Builder
	var cols strings.Builder

	cols.WriteString("\t<tr>\n")

	for i, row := range r {
		header.WriteString(fmt.Sprintf("\t\t<th> %d </th>\n", i))
		cols.WriteString(fmt.Sprintf("\t\t<td>\n%s\t\t</td>\n", row))
	}
	cols.WriteString("\t</tr>\n")

	sb.WriteString("<table>\n")
	sb.WriteString("\t<tr>\n")
	sb.WriteString(header.String())
	sb.WriteString("\t</tr>\n")
	sb.WriteString(cols.String())
	sb.WriteString("</table>\n")

	return sb.String()
}

type Text string

func (t Text) String() string {
	return string(t)
}

type H1 string

func (h H1) String() string {
	return fmt.Sprintf("# %s\n", string(h))
}

type H2 string

func (h H2) String() string {
	return fmt.Sprintf("## %s\n", string(h))
}

type H3 string

func (h H3) String() string {
	return fmt.Sprintf("### %s\n", string(h))
}

type Bold string

func (b Bold) String() string {
	return fmt.Sprintf("**%s**\n", string(b))
}

type Italic string

func (i Italic) String() string {
	return fmt.Sprintf("*%s*\n", string(i))
}

type OrderedList []string

func (o OrderedList) String() string {
	var sb strings.Builder

	for i, item := range o {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
	}

	return sb.String()
}

type List []string

func (l List) String() string {
	var sb strings.Builder

	for _, item := range l {
		sb.WriteString(fmt.Sprintf("- %s\n", item))
	}

	return sb.String()
}

type Code string

func (c Code) String() string {
	return fmt.Sprintf("```\n%s\n```\n", string(c))
}

type Table struct {
	Columns []string
	Rows    [][]string
}

func NewTable(columns ...string) Table {
	return Table{
		Columns: columns,
	}
}

func (t *Table) Add(row ...[]string) {
	t.Rows = append(t.Rows, row...)
}

func (t Table) String() string {
	var sb strings.Builder

	// determine length of each column title to get padding level
	var lengths []int
	for _, col := range t.Columns {
		lengths = append(lengths, len(col))
	}

	sb.WriteString(fmt.Sprintf("| %s |\n", strings.Join(t.Columns, " | ")))

	sb.WriteString("| ")
	for _, len := range lengths {
		sb.WriteString(fmt.Sprintf("%s |", strings.Repeat("-", len)))
	}
	sb.WriteString("\n")

	for _, row := range t.Rows {
		if len(row) < len(t.Columns) {
			// pad row
			row = append(row, make([]string, len(t.Columns)-len(row))...)
		}

		sb.WriteString(fmt.Sprintf("| %s |\n", strings.Join(row, " | ")))
	}

	return sb.String()
}

/*

| Syntax | Description |
| ----------- | ----------- |
| Header | Title |
| Paragraph | Text |
*/
