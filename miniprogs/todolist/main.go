package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
с цветом задач ошибка
Разделить на вкладки. 3 вкладки цели, задаци, заметки.
Добавить ежедневные задачи, расписание.
цвет/приоритет map?
Сортировать по приоритету.
Заметки - 2-3 блока для заметок. Просто квадрат многострочного поля ввода.
При сохранении в файл ставить дату(?)
добавить напоминалку (сообщение по дате)
будильник?
Фон - цветом зоны и заголовок выделить
*/

func mainForm() *fyne.Container {

	goalBox := goalForm()
	taskBox := taskForm()

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
	mainBox := container.NewVBox(goalBox, taskBox /*, noteBox,, debug*/)

	/*go func() {
		sec := time.NewTicker(3 * time.Second)
		for range sec.C {
			// отладить
			debug.SetText("")
			for i, g := range goalSlice {
				s := fmt.Sprintf("%d: %v ", i, g.Name)
				debug.Append(s)
			}
			// обновление элементов
			goalsBox = getGoalsBox(goalSlice)
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
// 										notes
// ----------------------------------------------------------------------------
// var notesSlice []notesType
/*
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
}*/

// ----------------------------------------------------------------------------
// 										общее
// ----------------------------------------------------------------------------
