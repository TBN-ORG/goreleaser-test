// main.go
package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func main() {
	println("Ba dum, tss!")
	fmt.Println("Do Dah")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Network", "Start Time", "Media ID", "Title"})
	t.AppendRow([]interface{}{"TBN", "now", "M1001", "Praise"})
	t.AppendFooter(table.Row{"", "", "Events Count", 1})
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.Render()
}
