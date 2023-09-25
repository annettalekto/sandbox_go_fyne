package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type taskStatus int

// taskType data
type taskType struct {
	Name       string
	Status     taskStatus
	Check      *widget.Check
	NameWidget *canvas.Text
	Box        *fyne.Container
}

var Tasks []taskType
var TasksDone binding.Float

const (
	veryImpotant taskStatus = iota
	Impotant
	Priority
	AnotherOne
	Housework
	ComputerStuff
)

func (t *taskType) Init(name string, status taskStatus) {
	t.Name = name
	t.Status = status
	cl := getColorOfStatus(t.Status)
	t.Check = widget.NewCheck("", func(b bool) {
		v, _ := TasksDone.Get()
		if b {
			TasksDone.Set(v + 1)
		} else {
			TasksDone.Set(v - 1)
		}
	})

	t.NameWidget = canvas.NewText(name, cl)
	t.NameWidget.TextSize = 14
	t.NameWidget.TextStyle.Monospace = true

	t.Box = container.NewHBox(t.Check, t.NameWidget)
}

// ----------------------------------------------------------------------------
// 									todo form
// ----------------------------------------------------------------------------

func taskForm() *fyne.Container {

	Tasks = readTasksFromFile()
	TasksDone = binding.NewFloat()

	pbar := widget.NewProgressBarWithData(TasksDone)
	pbar.Max = float64(len(Tasks))
	pbar.Min = 0
	pbar.SetValue(0)

	tasksBox := container.NewGridWithColumns(2)
	for _, t := range Tasks { // + сортировку и вынести в отд. ф.
		tasksBox.Add(t.Box)
	}

	addTask := widget.NewButton("New task", func() {
		addTaskForm(tasksBox)
	})
	cleanTask := widget.NewButton("Clean", func() {

	})

	buttonBox := container.NewHBox(addTask, cleanTask)

	box := container.NewVBox(tasksBox, pbar, buttonBox)
	return box
}

func addTaskForm(tb *fyne.Container) { // или расположить на главной форме entry
	w := fyne.CurrentApp().NewWindow("Создать")
	w.Resize(fyne.NewSize(400, 100))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	nameEntry := widget.NewEntry()
	nameBox := container.NewBorder(nil, nil, widget.NewLabel("Название: "), nil, nameEntry)

	// widget select какой то баг: обрезаются слова в версии v2.4.0, но в v2.3.4 этого нет
	priority := []string{"комп    ", "дом     ", "идти     ", "другое     ", "важно     ", "срочно     "}
	selectPriority := widget.NewSelect(priority, func(s string) {})
	selectPriority.SetSelected(priority[0])

	okButton := widget.NewButton("Ok", func() {
		if nameEntry.Text == "" {
			nameEntry.SetPlaceHolder("Сюда название, пожалуйста")
			return
		}
		var t taskType
		t.Init(nameEntry.Text, getStatus(selectPriority.SelectedIndex()))
		Tasks = append(Tasks, t)
		tb.Add(t.Box)
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 35)), okButton)
	selectBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 35)), selectPriority)
	box := container.NewBorder(nameBox, nil, nil, buttonBox, selectBox)

	w.SetContent(box)
	w.Show()
}

// func readTasksFromFile() []taskType {
func readTasksFromFile() []taskType {
	var tasks []taskType

	for i := 0; i <= 10; i++ {
		var temp taskType
		temp.Init("aaa", Impotant)
		tasks = append(tasks, temp)
	}
	for i := 0; i <= 10; i++ {
		var temp taskType
		temp.Init("bbb", ComputerStuff)
		tasks = append(tasks, temp)
	}

	return tasks
}

// ----------------------------------------------------------------------------
// 										приоритет
// ----------------------------------------------------------------------------

// todo: все цвета в отдельный файл (библ...типо)
var ( // todo: без приоритета - черный, для заметок
	red    = color.NRGBA{R: 255, G: 0, B: 0, A: 255}    // 0: очень срочно!
	purple = color.NRGBA{R: 184, G: 15, B: 200, A: 255} // 1: срочно
	orange = color.NRGBA{R: 255, G: 50, B: 20, A: 255}  // 2: в приоритете
	jellow = color.NRGBA{R: 255, G: 230, B: 5, A: 255}  // 3: другое
	green  = color.NRGBA{R: 0, G: 255, B: 0, A: 255}    // 4: домашние дела
	blue   = color.NRGBA{R: 0, G: 0, B: 255, A: 255}    // 5: дела за компом (обучение, работа)
)

// todo: map?
func getColorOfStatus(status taskStatus) color.NRGBA {
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

func getStatus(n int) taskStatus {
	var st taskStatus

	switch n {
	case 0:
		st = Housework
	case 1:
		st = ComputerStuff
	case 2:
		st = AnotherOne
	case 3:
		st = Priority
	case 4:
		st = Impotant
	case 5:
		st = veryImpotant
	default:
		// cl = color.Black
	}
	return st
}
