package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// var goalSlice []goalType

// goalType data
type goalType struct {
	Name, Note   string
	Max          float64
	ProgressBar  *widget.ProgressBar // тут value
	PlusButton   *widget.Button      // todo: проредить
	ChangeButton *widget.Button
	Box          *fyne.Container
	// note: добавить цельное название / описание
}

// Create for goalType's progressBar
func (g *goalType) Create(name, note string, max float64) {
	g.Name = name
	g.Note = note
	g.Max = max

	// label := widget.NewLabel(g.Name)
	text := canvas.NewText("     "+g.Name, color.Black) // без пробелов выходит за прогресс бар слева
	text.TextStyle.Italic = true
	textBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(0, 30)), text)

	g.ProgressBar = widget.NewProgressBar()
	g.ProgressBar.Max = g.Max
	g.ProgressBar.Min = 0
	g.ProgressBar.SetValue(0)

	g.PlusButton = widget.NewButton("  +  ", func() {
		g.ProgressBar.Value++
		g.ProgressBar.Refresh()
	})
	g.ChangeButton = widget.NewButton("  ...  ", func() {
		g.ChangeGoalForm()
	})
	buttonBox := container.NewHBox(g.PlusButton, g.ChangeButton)
	g.Box = container.NewBorder(nil, nil, textBox, buttonBox, g.ProgressBar)
}

// ChangeValue прибавить прогресс
func (g *goalType) ChangeValue() {
	g.ProgressBar.Value++
	g.ProgressBar.Refresh()
}

// ChangeGoalForm форма для изменения парамметров цели
func (g *goalType) ChangeGoalForm() {
	w := fyne.CurrentApp().NewWindow("Изменить") // CurrentApp!
	w.Resize(fyne.NewSize(400, 190))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	name := canvas.NewText(g.Name, color.Black)
	name.TextSize = 16
	name.TextStyle.Monospace = true
	nameBox := container.NewHBox(layout.NewSpacer(), name, layout.NewSpacer())
	noteEntry := widget.NewEntry()
	if g.Note == "" {
		noteEntry.SetPlaceHolder("Введите примечание...")
	} else {
		noteEntry.SetText(g.Note)
	}

	maxValueEntry := newNumericalEntry() // установка по нажатию
	maxValueEntry.SetPlaceHolder(fmt.Sprintf("%v", g.ProgressBar.Value))
	boxValue := container.NewBorder(nil, nil, widget.NewLabel("Сделано: "), nil, maxValueEntry)

	doneButton := widget.NewButton("Завершить", func() {
		// окно с вопросом если не сделано 100%
		// сохранить в отдел завершенные в файле
	})
	deleteButton := widget.NewButton("Удалить", func() {
		// окно с вопросом
		// удалить из слайса, файла и формы
	})
	okButton := widget.NewButton("Ok", func() {

	})
	buttonBox := container.NewHBox(deleteButton, doneButton, layout.NewSpacer(), okButton)
	// buttonBox = container.NewBorder(nil, nil, nil, buttonBox)

	box := container.NewVBox(nameBox, noteEntry, boxValue, widget.NewLabel(""), buttonBox)
	w.SetContent(box)
	w.Show() // ShowAndRun -- panic!
}
