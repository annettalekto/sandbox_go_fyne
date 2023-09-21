package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
