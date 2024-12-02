package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Score", "Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for i := 0; i < 10; i++ {
		tbl.AddRow(i, i, i, i)
	}

	tbl.Print()
}
