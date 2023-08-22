package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ----------------------------------------------------------------------------
// 										goal
// ----------------------------------------------------------------------------

// var goalSlice []goalType

// goalType data
type goalType struct {
	Name, Note  string
	Max, Value  float64
	ProgressBar *widget.ProgressBar
	Button      *widget.Button
	Box         *fyne.Container
	// note: добавить цельное название / описание
}

// Create for goalType's progressBar
func (g *goalType) Create(name string, max float64) {
	g.Name = name
	g.Max = max
	g.Value = 0

	label := widget.NewLabel(g.Name)

	g.ProgressBar = widget.NewProgressBar()
	g.ProgressBar.Max = g.Max
	g.ProgressBar.Min = 1
	g.ProgressBar.SetValue(0)

	g.Button = widget.NewButton("  +  ", nil)

	g.Box = container.NewBorder(nil, nil, label, g.Button, g.ProgressBar)
}

func NewGoalForm() {
	w := fyne.CurrentApp().NewWindow("Создать") // CurrentApp!
	w.Resize(fyne.NewSize(400, 150))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	label := widget.NewLabel("хм")

	w.SetContent(label)
	w.Show() // ShowAndRun -- panic!
}

func ChangeGoalForm() {

}
