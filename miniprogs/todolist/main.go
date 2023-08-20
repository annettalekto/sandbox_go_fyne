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

	check1 := CreateCheck("Уборка")
	check2 := CreateCheck("Йога")
	barBox := container.NewVBox(goal1.box, goal2.box, goal3.box, check1, check2)

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
	name        string
	max         float64
	value       float64
	progressBar *widget.ProgressBar
	button      *widget.Button
	box         *fyne.Container
}

// Create for goalType's progressBar
func (g *goalType) Create(name string, max float64) {
	g.name = name
	g.max = max
	g.value = 0

	label := widget.NewLabel(g.name)

	g.progressBar = widget.NewProgressBar()
	g.progressBar.Max = g.max
	g.progressBar.Min = 1
	g.progressBar.SetValue(0)

	g.button = widget.NewButton("Изменить", nil)

	boxH := container.NewBorder(nil, nil, nil, g.button, g.progressBar)

	g.box = container.NewVBox(label, boxH)
}

// ----------------------------------------------------------------------------
// 										todo
// ----------------------------------------------------------------------------
// var todoSlice []goalType

// todoType data
// type goalType struct {
// 	name        string
// 	max         float64
// 	value       float64
// 	progressBar *widget.ProgressBar
// 	button      *widget.Button
// 	box         *fyne.Container
// }

func CreateCheck(name string) *widget.Check {
	check := widget.NewCheck(name, nil)
	return check
}

// ----------------------------------------------------------------------------
// 										notes
// ----------------------------------------------------------------------------
