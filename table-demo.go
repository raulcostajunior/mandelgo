package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{{"top left", "top right"},
	{"bottom left", "bottom right"}}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content to reserve space for cells")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	myWindow.SetContent(list)
	myWindow.Resize(fyne.NewSize(650, 220))
	myWindow.ShowAndRun()
}
