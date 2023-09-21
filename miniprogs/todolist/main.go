package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO List")
	w.Resize(fyne.NewSize(600, 600))
	w.CenterOnScreen()
	w.SetMaster()

	w.SetContent(mainForm())
	w.ShowAndRun()
}

/*
todo:
цели - заменить на текс, с подходящим цветом, сделать прогресс бар на всю длинну, масштабируемым
кнопку новая цель на верх, добавить заголовок?
Задачи -  можно ли сделать в 2 ряда?
Сортировать по приоритету.
Убрать заметки к каждому пункту
Заметки - 2-3 блока для заметок. Просто квадрат многострочного поля ввода.
добавить напоминалку (сообщение по дате)
будильник?
*/

type mainFormType struct {
	Goals []goalType
}

var mainFormData mainFormType

func mainForm() *fyne.Container {
	// var err error

	text := canvas.NewText("ToDoList", color.Black)
	text.TextStyle.Monospace = true

	mainFormData.Goals = append(mainFormData.Goals, getGoalsFromFile()...)
	goalsBox := getGoalsBox(mainFormData.Goals)
	addGoalButton := widget.NewButton("New goal", func() {
		newGoalForm(goalsBox)
	})
	b := container.NewBorder(nil, nil, nil, addGoalButton, text)
	goalsAllBox := container.NewVBox(b, goalsBox)

	// todo: форма для изменения goals ...
	//-----------------

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

	/*	var note1, note2 noteType // note: разделительные лайблы выделить полосой
		note1.Create("Незабыть про голицина")
		note2.Create("вычесать кошку")
		box = container.NewVBox(widget.NewLabel("Заметки:"), note1.TextWidget, note2.TextWidget)
		addNote := widget.NewButton("New note", nil)
		cleanAll := widget.NewButton("Clean all", nil) // todo: заменить на удаление по одной
		buttonBox = container.NewHBox(addNote, cleanAll)
		noteBox := container.NewBorder(box, nil, nil, buttonBox)
		// todo: или сделать задача - заметка и тд.
		// придется добавить прокрутку
	*/
	//debug := widget.NewMultiLineEntry()
	mainBox := container.NewVBox(goalsAllBox, taskBox /*, noteBox,, debug*/)

	/*go func() {
		sec := time.NewTicker(3 * time.Second)
		for range sec.C {
			// отладить
			debug.SetText("")
			for i, g := range mainFormData.Goals {
				s := fmt.Sprintf("%d: %v ", i, g.Name)
				debug.Append(s)
			}
			// обновление элементов
			goalsBox = getGoalsBox(mainFormData.Goals)
			// goalsBox.Refresh()
			// goalsAllBox = container.NewBorder(goalsBox, nil, nil, addGoalButton)
			// goalsAllBox.Refresh()
			// mainBox = container.NewVBox(goalsAllBox, taskBox, noteBox, debug)
			// mainBox.Refresh()
		}
	}()*/

	return mainBox
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
		if max > 1000000 { //
			err = fmt.Errorf("\"%s\" слишком большое (более 1 000 000)", maxValueStr)
			errorLabel.Text = err.Error()
			errorLabel.Refresh()
			return
		}
		errorLabel.Text = "ок"
		var g goalType
		g.Create(name, note, float64(max))
		mainFormData.Goals = append(mainFormData.Goals, g)
		goalsBox.Add(g.Box)
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 30)), buttonOk) // size
	buttonBox = container.NewBorder(nil, nil, nil, buttonBox, nil)                       // left
	box := container.NewVBox(grid, buttonBox, errorLabel)

	w.SetContent(box)
	w.Show() // ShowAndRun -- panic!
}

func getGoalsBox(goals []goalType) *fyne.Container {

	// note: при выводе сортировать как то?
	box := container.NewVBox()
	for _, g := range goals {
		box.Add(g.Box)
	}
	//box = container.New(layout.NewGridWrapLayout(fyne.NewSize(780, 200)), container.NewVScroll(box))
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
