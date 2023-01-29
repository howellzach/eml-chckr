package util

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func GenerateTable(title string, fileName string, data map[string]string) {
	t := table.NewWriter()
	tableTitle := title + " - " + fileName
	t.AppendHeader(
		table.Row{tableTitle, tableTitle},
		table.RowConfig{AutoMerge: true},
	)
	for key, value := range data {
		t.AppendRow(
			table.Row{key, value},
			table.RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignLeft},
		)
	}
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.SetColumnConfigs(
		[]table.ColumnConfig{
			{Number: 1, WidthMax: 90},
			{Number: 2, WidthMax: 70},
		})
	t.Render()
}
