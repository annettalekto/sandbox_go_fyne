package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// goalType data
type goalType struct { // todo: удалить лишнее
	Name, Note   string
	Max          float64
	ProgressBar  *widget.ProgressBar // тут value
	PlusButton   *widget.Button      // todo: проредить
	ChangeButton *widget.Button
	Box          *fyne.Container
	// note: добавить цельное название / описание
}

var goalSlice []goalType

// Init for goalType's progressBar
func (g *goalType) Init(name, note string, max float64) {
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

// IncrementProgress прибавить прогресс
func (g *goalType) IncrementProgress() {
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
	boxValue := container.NewBorder(nil, nil, widget.NewLabel("Сделано: "),
		widget.NewLabel(fmt.Sprintf("(из %v)", g.ProgressBar.Max)), maxValueEntry)

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

// ----------------------------------------------------------------------------
// 									goal form
// ----------------------------------------------------------------------------

func goalForm() *fyne.Container {

	goalSlice = append(goalSlice, readGoalsFromFile()...)
	goalsBox := createGoalsBox(goalSlice)
	addGoalButton := widget.NewButton("Новая цель", func() {
		newGoalForm(goalsBox)
	})

	button := container.NewBorder(nil, nil, nil, addGoalButton)

	notesEntry := widget.NewMultiLineEntry()
	notesEntry.Wrapping = fyne.TextWrapWord

	box := container.NewVBox(goalsBox, button)
	return container.NewBorder(box, nil, nil, nil, notesEntry)
	// return container.NewBorder(nil, nil, nil, nil, notesEntry)
}

func newGoalForm(goalsBox *fyne.Container) {
	w := fyne.CurrentApp().NewWindow("Создать") // CurrentApp!
	w.Resize(fyne.NewSize(500, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	var err error
	var name, note string // todo: как передать данные
	var max int
	errorLabel := widget.NewLabel("...") // вывод ошибок

	nameStr := "Название"
	nameEntry := widget.NewEntry()
	noteStr := "Примечание"
	noteEntry := widget.NewEntry()
	maxValueStr := "Максимальноe число задач"
	maxValueEntry := newNumericalEntry()

	grid := container.NewGridWithColumns(2,
		widget.NewLabel(nameStr+": "), nameEntry,
		widget.NewLabel(noteStr+": "), noteEntry,
		widget.NewLabel(maxValueStr+": "), maxValueEntry,
	)

	buttonOk := widget.NewButton("OK", func() {
		name = nameEntry.Text
		if name == "" {
			err = fmt.Errorf(fmt.Sprintf("Поле ввода \"%s\" не может быть пустым", nameStr))
			errorLabel.Text = err.Error()
			errorLabel.Refresh()
			return
		}
		note = noteEntry.Text
		maxStr := maxValueEntry.Text
		if maxStr == "" {
			err = fmt.Errorf("поле ввода \"%s\" не может быть пустым", maxValueStr)
			errorLabel.Text = err.Error()
			errorLabel.Refresh()
			return
		}
		max, err = strconv.Atoi(maxStr)
		if err != nil {
			err = fmt.Errorf("ошибка в поле ввода \"%s\"", maxValueStr)
			errorLabel.Text = err.Error()
			errorLabel.Refresh()
			return
		}
		if max <= 0 {
			err = fmt.Errorf("\"%s\" должно быть меньше нуля", maxValueStr)
			errorLabel.Text = err.Error()
			errorLabel.Refresh()
			return
		}
		if max > 1000000 {
			err = fmt.Errorf("\"%s\" слишком большое (более 1 000 000)", maxValueStr)
			errorLabel.Text = err.Error()
			errorLabel.Refresh()
			return
		}
		errorLabel.Text = "ок"
		var g goalType
		g.Init(name, note, float64(max))
		goalSlice = append(goalSlice, g)
		goalsBox.Add(g.Box)
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 30)), buttonOk) // size
	buttonBox = container.NewBorder(nil, nil, nil, buttonBox, nil)                       // left
	box := container.NewVBox(grid, buttonBox, errorLabel)

	w.SetContent(box)
	w.Show() // ShowAndRun -- panic!
}

func createGoalsBox(goals []goalType) *fyne.Container {

	// note: при выводе сортировать как то?
	box := container.NewVBox()
	for _, g := range goals {
		box.Add(g.Box)
	}
	//box = container.New(layout.NewGridWrapLayout(fyne.NewSize(780, 200)), container.NewVScroll(box))
	return box
}

func readGoalsFromFile() []goalType { // todo: File!
	var goals []goalType
	var goal1, goal2, goal3 goalType
	goal1.Init("Читать ITM:", "", 300)
	goal2.Init("Читать ENG:", "", 1300)
	goal3.Init("Перебрать тетради:", "", 15)
	goals = append(goals, goal1, goal2, goal3)
	return goals
}
