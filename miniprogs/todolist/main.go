package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

/*
todo:
// добавить напоминалку (сообщение по дате)
// будильник?
*/

func mainForm() *fyne.Container {
	var Goals []goalType

	Goals = append(Goals, getGoalsFromFile()...)
	addGoal := widget.NewButton("New goal", func() {
		newGoalForm()
	})
	goalBox := container.NewBorder(getGoalsBox(Goals), nil, nil, addGoal)

	var task1, task2 taskType
	done := binding.NewFloat()
	task1.Create("Go test", "slice разбор", ComputerStuff)
	task2.Create("Йога", "3 упр", Housework)
	pbar := widget.NewProgressBarWithData(done)
	pbar.Max = 2 // количество задач на сегодня
	pbar.Min = 1
	pbar.SetValue(0)
	box := container.NewVBox(widget.NewLabel("Задачи на сегдня:"), task1.Box, task2.Box, pbar) // todo: задачи label ярче
	addTask := widget.NewButton("New task", nil)
	cleanTask := widget.NewButton("Clean", nil)
	buttonBox := container.NewHBox(addTask, cleanTask)
	taskBox := container.NewBorder(box, nil, nil, buttonBox)
	// см сколько задач -> добавить прогресс бар на сегодня, прибавлять по завершению задач

	var note1, note2 noteType // note: разделиетельные лайблы выделить полосой
	note1.Create("Незабыть про голицина")
	note2.Create("вычесать кошку")
	box = container.NewVBox(widget.NewLabel("Заметки:"), note1.TextWidget, note2.TextWidget)
	addNote := widget.NewButton("New note", nil)
	cleanAll := widget.NewButton("Clean all", nil) // todo: заменить на удаление по одной
	buttonBox = container.NewHBox(addNote, cleanAll)
	noteBox := container.NewBorder(box, nil, nil, buttonBox)
	// todo: или сделать задача - заметка и тд.
	// придется добавить прокрутку

	return container.NewVBox(goalBox, taskBox, noteBox)
}

// ----------------------------------------------------------------------------
// 										goal
// ----------------------------------------------------------------------------

func getGoalsFromFile() []goalType { // todo: File!
	var goal1, goal2, goal3 goalType
	goal1.Create("Читать ITM:", "", 300)
	goal2.Create("Читать ENG:", "", 1300)
	goal3.Create("Перебрать тетради:", "", 15)
	var goals []goalType
	goals = append(goals, goal1, goal2, goal3)
	return goals
}

func newGoalForm() {
	w := fyne.CurrentApp().NewWindow("Создать") // CurrentApp!
	w.Resize(fyne.NewSize(400, 150))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	label := widget.NewLabel("хм")

	w.SetContent(label)
	w.Show() // ShowAndRun -- panic!
}

func getGoalsBox(goals []goalType) *fyne.Container {

	// note: при выводе сортировать как то?
	box := container.NewVBox()
	for _, g := range goals {
		box.Add(g.Box)
	}
	return box
}

// ----------------------------------------------------------------------------
//										todo
// ----------------------------------------------------------------------------

// var todoSlice []taskType
// todo: разобрать на файлы?

type taskStatus int

// taskType data
type taskType struct {
	Name       string
	Note       string
	Status     taskStatus
	Check      *widget.Check
	NameWidget *canvas.Text
	NoteWidget *canvas.Text
	Box        *fyne.Container
	// Button *widget.Button
}

func (t *taskType) Create(name, note string, status taskStatus) {
	t.Name = name
	t.Note = note
	t.Status = status
	cl := GetColorOfStatus(t.Status)
	t.Check = widget.NewCheck("", nil)

	t.NameWidget = canvas.NewText(name, cl)
	t.NameWidget.TextSize = 14
	t.NameWidget.TextStyle.Monospace = true

	t.NoteWidget = canvas.NewText("    ("+note+")", color.Black)
	t.NoteWidget.TextSize = 10
	t.NameWidget.TextStyle.Italic = true

	// t.Box = container.NewVBox(container.NewHBox(t.Check, t.NameWidget), t.NoteWidget) // пояснение к задаче снизу
	t.Box = container.NewHBox(t.Check, t.NameWidget, t.NoteWidget) // пояснение к задаче снизу
}

// ----------------------------------------------------------------------------
// 										notes
// ----------------------------------------------------------------------------
// var notesSlice []notesType

// noteType data
type noteType struct {
	Name string
	// Status     taskStatus
	TextWidget *canvas.Text
	// todo: добавить удаление или завершение + удаление завершенных
	// Button *widget.Button
	// Box    *fyne.Container
}

func (t *noteType) Create(name string) {
	t.Name = name
	// t.Status = status
	// cl := GetColorOfStatus(t.Status)

	t.TextWidget = canvas.NewText(name, color.Black)
	t.TextWidget.TextSize = 14
	t.TextWidget.TextStyle.Italic = true
}

// ----------------------------------------------------------------------------
// 										общее
// ----------------------------------------------------------------------------

var ( // todo: без приоритета - черный, для заметок
	red    = color.NRGBA{R: 255, G: 0, B: 0, A: 255}    // 0: очень срочно!
	purple = color.NRGBA{R: 184, G: 15, B: 200, A: 255} // 1: срочно
	orange = color.NRGBA{R: 255, G: 50, B: 20, A: 255}  // 2: в приоритете
	jellow = color.NRGBA{R: 255, G: 230, B: 5, A: 255}  // 3: другое
	green  = color.NRGBA{R: 0, G: 255, B: 0, A: 255}    // 4: домашние дела
	blue   = color.NRGBA{R: 0, G: 0, B: 255, A: 255}    // 5: дела за компом (обучение, работа)
)

const (
	veryImpotant taskStatus = iota
	Impotant
	Priority
	AnotherOne
	Housework
	ComputerStuff
)

func GetColorOfStatus(status taskStatus) color.NRGBA {
	var cl color.NRGBA

	switch status {
	case veryImpotant:
		cl = red
	case Impotant:
		cl = purple
	case Priority:
		cl = orange
	case AnotherOne:
		cl = jellow
	case ComputerStuff:
		cl = blue
	case Housework:
		cl = green
	default:
		// cl = color.Black
	}
	return cl
}
