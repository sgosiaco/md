package md

import (
	"fmt"
	"strings"
)

func (c Column) String() string {
	var sb strings.Builder

	for _, item := range c {
		sb.WriteString(item.String())
		// one newline for gap
		sb.WriteString("\n")
	}

	return sb.String()
}

func (r Row) String() string {
	var sb strings.Builder
	var cols strings.Builder

	cols.WriteString("<tr>\n")
	for _, column := range r {
		cols.WriteString(fmt.Sprintf("<td>\n\n%s</td>\n", column))
	}
	cols.WriteString("</tr>\n")

	sb.WriteString("<table>\n")
	sb.WriteString(cols.String())
	sb.WriteString("</table>\n")

	return sb.String()
}

func (h HeaderRow) String() string {
	var sb strings.Builder
	var header strings.Builder
	var cols strings.Builder

	header.WriteString("<tr>\n")
	for _, h := range h.Headers {
		// adding /n after %s since str only
		header.WriteString(fmt.Sprintf("<th>\n\n%s\n</th>\n", h))
	}
	header.WriteString("</tr>\n")

	cols.WriteString("<tr>\n")
	for _, row := range h.Columns {
		// convert each renderable to html in order to ensure maximum compatibility
		cols.WriteString(fmt.Sprintf("<td>\n\n%s</td>\n", ToHTML(row.String())))
	}
	cols.WriteString("</tr>\n")

	sb.WriteString("<table>\n")
	sb.WriteString("<thead>\n")
	sb.WriteString(header.String())
	sb.WriteString("</thead>\n")
	sb.WriteString("<tbody>\n")
	sb.WriteString(cols.String())
	sb.WriteString("</tbody>\n")
	sb.WriteString("</table>\n")

	return sb.String()
}

func (t Text) String() string {
	return string(t)
}

func (h H1) String() string {
	return fmt.Sprintf("# %s\n", string(h))
}

func (h H2) String() string {
	return fmt.Sprintf("## %s\n", string(h))
}

func (h H3) String() string {
	return fmt.Sprintf("### %s\n", string(h))
}

func (b Bold) String() string {
	return fmt.Sprintf("**%s**\n", string(b))
}

func (i Italic) String() string {
	return fmt.Sprintf("*%s*\n", string(i))
}

func (s Strikethrough) String() string {
	return fmt.Sprintf("~~%s~~\n", string(s))
}

func (b BlockQuote) String() string {
	return "> " + strings.Join(strings.Split(string(b), "\n"), "\n> ")
}

func (o OrderedList) String() string {
	var sb strings.Builder

	for i, item := range o {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
	}

	return sb.String()
}

func (l List) String() string {
	var sb strings.Builder

	for _, item := range l {
		sb.WriteString(fmt.Sprintf("- %s\n", item))
	}

	return sb.String()
}

func (c Code) String() string {
	return fmt.Sprintf("`%s`\n", string(c))
}

func (c CodeBlock) String() string {
	return fmt.Sprintf("```\n%s\n```\n", string(c))
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
