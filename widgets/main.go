package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne test")
	w.Resize(fyne.NewSize(800, 600))

	// варианты label
	l1 := widget.NewLabel("Test label: Italic style")
	l1.TextStyle.Italic = true
	l2 := widget.NewLabel("Test label: Monospace style")
	l2.TextStyle.Monospace = true
	l3 := widget.NewLabel("Test label: Bold style")
	l3.TextStyle.Bold = true
	l4 := widget.NewLabel("Test label: Symbol style")
	l4.TextStyle.Symbol = true
	l5 := widget.NewLabel("Test label: TabWidth=1")
	l5.TextStyle.TabWidth = 1
	l6 := widget.NewLabel("Test label: TabWidth=5")
	l6.TextStyle.TabWidth = 5

	labelBox := container.NewVBox(l1, l2, l3, l4, l5, l6)
	w.SetContent(labelBox)
	w.ShowAndRun()
}
