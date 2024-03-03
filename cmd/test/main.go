package main

import (
	"fmt"
	"os"

	md "github.com/sgosiaco/md/pkg"
)

func main() {
	var c md.Column

	t := md.NewTable("Field", "Value")

	t.Add(
		[]string{"Name", "Test"},
		[]string{"Age", "-1"},
		[]string{"Hello"},
	)

	c.Add(
		md.H1("Testing"),
		t,
		md.Bold("Very bold"),
		md.Italic("Emphasis"),
		md.Text("Lorem ipsum etc"),
		md.List{"Apple", "Banana", "Orange"},
		md.OrderedList{"Item 1", "Item 2", "Item 3"},
		md.Row{
			t,
			t,
		},
	)

	fmt.Print(c.String())

	f, err := os.Create("test.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(c.String())
}
