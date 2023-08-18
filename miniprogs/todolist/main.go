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

	pbar1 := CreateGoal("Читать ITM", 300)
	pbar2 := CreateGoal("Читать ENG", 1300)
	pbar3 := CreateGoal("Перебрать тетради", 15)

	barBox := container.NewVBox(pbar1, pbar2, pbar3)

	split := container.NewHSplit(barBox, l2)
	return container.NewBorder(nil, nil, nil, nil, split)
}

func CreateGoal(name string, max float64) *fyne.Container {
	label := widget.NewLabel(name)

	pbar := widget.NewProgressBar()
	pbar.Max = max
	pbar.Min = 1
	pbar.SetValue(0)

	addProgress := widget.NewButton("Прогресс", nil)

	boxH := container.NewBorder(nil, nil, nil, addProgress, pbar)

	return container.NewVBox(label, boxH)
}
