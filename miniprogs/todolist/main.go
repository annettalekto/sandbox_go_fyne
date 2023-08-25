package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO List")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.SetMaster()

	w.SetContent(mainForm())
	w.ShowAndRun()
}

func mainForm() *fyne.Container {

	var goal1, goal2, goal3 goalType
	goal1.Create("Читать ITM:", 300)
	goal2.Create("Читать ENG:", 1300)
	goal3.Create("Перебрать тетради:", 15)

	var task1, task2 taskType
	task1.Create("Уборка")
	task2.Create("Йога")
	barBox := container.NewVBox(goal1.Box, goal2.Box, goal3.Box, task1.Check, task2.Check)

	l2 := widget.NewLabel("buttons")
	split := container.NewHSplit(barBox, l2)
	return container.NewBorder(nil, nil, nil, nil, split)
}

// ----------------------------------------------------------------------------
// 										goal
// ----------------------------------------------------------------------------

// var goalSlice []goalType

// goalType data
type goalType struct {
	Name        string
	Max         float64
	Value       float64
	ProgressBar *widget.ProgressBar
	Button      *widget.Button
	Box         *fyne.Container
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

	boxH := container.NewBorder(nil, nil, nil, g.Button, g.ProgressBar)

	g.Box = container.NewVBox(label, boxH)
}

// ----------------------------------------------------------------------------
// 										todo
// ----------------------------------------------------------------------------
// var todoSlice []goalType

// taskType data
type taskType struct {
	Name string

	Check *widget.Check
	// Button *widget.Button
	// Box    *fyne.Container
}

func (t *taskType) Create(name string) {
	t.Name = name
	t.Check = widget.NewCheck(name, nil)
}

// ----------------------------------------------------------------------------
// 										notes
// ----------------------------------------------------------------------------
