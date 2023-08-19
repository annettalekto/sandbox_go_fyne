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
	// l1 := widget.NewLabel("Objects")
	l2 := widget.NewLabel("buttons")

	// bar := widget.NewProgressBar()
	// bar.Max = 100
	// bar.Min = 1
	// bar.SetValue(10)
	// box := container.NewVBox(l1, bar)

	var goal1, goal2, goal3 goalType
	goal1.Create("Читать ITM", 300)
	goal2.Create("Читать ENG", 1300)
	goal3.Create("Перебрать тетради", 15)

	barBox := container.NewVBox(goal1.box, goal2.box, goal3.box)

	split := container.NewHSplit(barBox, l2)
	return container.NewBorder(nil, nil, nil, nil, split)
}

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

	g.button = widget.NewButton("Прогресс", nil)

	boxH := container.NewBorder(nil, nil, nil, g.button, g.progressBar)

	g.box = container.NewVBox(label, boxH)
}
