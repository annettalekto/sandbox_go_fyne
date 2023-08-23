package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// var goalSlice []goalType

// goalType data
type goalType struct {
	Name, Note  string
	Max         float64
	ProgressBar *widget.ProgressBar
	Button      *widget.Button
	Box         *fyne.Container
	// note: добавить цельное название / описание
}

// Create for goalType's progressBar
func (g *goalType) Create(name, note string, max float64) {
	g.Name = name
	g.Note = note
	g.Max = max

	label := widget.NewLabel(g.Name)

	g.ProgressBar = widget.NewProgressBar()
	g.ProgressBar.Max = g.Max
	g.ProgressBar.Min = 1
	g.ProgressBar.SetValue(0)

	g.Button = widget.NewButton("  +  ", func() {
		// changeGoalForm()
		g.ChangeValue()
	})

	g.Box = container.NewBorder(nil, nil, label, g.Button, g.ProgressBar)
}

// ChangeValue прибавить прогресс
func (g *goalType) ChangeValue() {
	g.ProgressBar.Value++
	g.ProgressBar.Refresh()
	// изменение кнопками + -?
	// изменение через поле ввода (Сделано: )
	// Изменить описание?
	// кнопка завершить (должна делать запись в файл и удалять виджит)
	// удалить (удалить везде)
}

func changeGoalForm() {
	w := fyne.CurrentApp().NewWindow("Изменить") // CurrentApp!
	w.Resize(fyne.NewSize(400, 150))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	label := widget.NewLabel("гм")

	w.SetContent(label)
	w.Show() // ShowAndRun -- panic!
}
