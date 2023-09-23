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

var Tasks []taskType
var TasksDone binding.Float

type taskStatus int

const (
	veryImpotant taskStatus = iota
	Impotant
	Priority
	AnotherOne
	Housework
	ComputerStuff
)

// taskType data
type taskType struct {
	Name       string
	Status     taskStatus
	Check      *widget.Check
	NameWidget *canvas.Text
	Box        *fyne.Container
}

func (t *taskType) Create(name string, status taskStatus) {
	t.Name = name
	t.Status = status
	cl := GetColorOfStatus(t.Status)
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

func getTasksFromFile() []taskType {
	var tasks []taskType

	for i := 0; i <= 10; i++ {
		var temp taskType
		temp.Create("aaa", Impotant)
		tasks = append(tasks, temp)
	}
	for i := 0; i <= 10; i++ {
		var temp taskType
		temp.Create("bbb", ComputerStuff)
		tasks = append(tasks, temp)
	}

	return tasks
}

func taskForm() *fyne.Container {

	Tasks = getTasksFromFile()
	TasksDone = binding.NewFloat()

	pbar := widget.NewProgressBarWithData(TasksDone)
	pbar.Max = float64(len(Tasks))
	pbar.Min = 0
	pbar.SetValue(0)

	b := container.NewGridWithColumns(2)
	for _, t := range Tasks {
		b.Add(t.Box)
	}

	addTask := widget.NewButton("New task", func() {
		createTaskForm()
	})
	cleanTask := widget.NewButton("Clean", func() {

	})

	buttonBox := container.NewHBox(addTask, cleanTask)

	box := container.NewVBox(b, pbar, buttonBox)
	return box
}

func createTaskForm() { // или расположить на главной форме entry
	w := fyne.CurrentApp().NewWindow("Создать")
	w.Resize(fyne.NewSize(400, 100))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	nameEntry := widget.NewEntry()
	b := container.NewBorder(nil, nil, widget.NewLabel("Название: "), nil, nameEntry)

	priority := []string{"очень срочно!", "срочно", "в приоритете", "другое", "домашние дела", "дела за компом"}
	selectPriority := widget.NewSelect(priority, func(s string) {
	})
	selectPriority.SetSelectedIndex(5)

	okButton := widget.NewButton("Ok", func() {
		if nameEntry.Text == "" {
			nameEntry.SetPlaceHolder("Сюда название, пожалуйста")
			return
		}
		var t taskType
		t.Create(nameEntry.Text, ComputerStuff)
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 40)), okButton)
	selectBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 40)), selectPriority)

	box := container.NewBorder(b, nil, nil, buttonBox, selectBox)
	w.SetContent(box)
	w.Show()
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
