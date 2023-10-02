package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// taskType data
type taskType struct {
	Check *widget.Check
	Box   *fyne.Container
}

var Tasks []taskType
var TasksDone binding.Float

func (t *taskType) Init(name string, priotity taskPriority) {

	t.Check = widget.NewCheck("", func(b bool) {
		v, _ := TasksDone.Get()
		if b {
			TasksDone.Set(v + 1)
		} else {
			TasksDone.Set(v - 1)
			if v-1 < 0 {
				fmt.Println(v)
			}
		}
	})

	nameWidget := canvas.NewText(name, getColorOfPriority(priotity))
	nameWidget.TextSize = 14
	nameWidget.TextStyle.Monospace = true

	t.Box = container.NewHBox(t.Check, nameWidget)
}

// ----------------------------------------------------------------------------
// 									todo form
// ----------------------------------------------------------------------------

func taskForm() *fyne.Container {

	Tasks = readTasksFromFile()
	TasksDone = binding.NewFloat()

	pbar := widget.NewProgressBarWithData(TasksDone)
	pbar.Max = 0.1     // вместо 0,
	pbar.Min = 1       // так не будет ошибки отображения %
	TasksDone.Set(0.1) // (если поставить нули процент в минус уходит -9465465843%)

	tasksBox := container.NewGridWithColumns(2)
	for _, t := range Tasks { // + сортировку и вынести в отд. ф.
		tasksBox.Add(t.Box)
	}

	addTask := widget.NewButton("Новая задача", func() {
		addTaskForm(tasksBox, pbar)
	})

	cleanTask := widget.NewButton("Удалить отмеченные", func() {
		for i := 0; i < len(Tasks); {
			t := Tasks[i]
			if t.Check.Checked { // если пункт отмечен, то удалить
				Tasks = removeTask(Tasks, i) // удалить из среза
				tasksBox.Remove(t.Box)       // удалить с формы
				// удалить из файла
			} else {
				i++
			}
		}
		if len(Tasks) == 0 {
			pbar.Max = 0.1
			TasksDone.Set(0.1)

		} else {
			pbar.Max = float64(len(Tasks))
			TasksDone.Set(0)
		}
		pbar.Refresh()
	})

	buttonBox := container.NewBorder(nil, nil, cleanTask, addTask)
	box := container.NewVBox(tasksBox, pbar, buttonBox)
	return box
}

func removeTask(slice []taskType, i int) []taskType {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func addTaskForm(tb *fyne.Container, pbar *widget.ProgressBar) { // или расположить на главной форме entry
	w := fyne.CurrentApp().NewWindow("Создать")
	w.Resize(fyne.NewSize(400, 130))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	nameEntry := widget.NewEntry()
	nameBox := container.NewBorder(nil, nil, widget.NewLabel("Название: "), nil, nameEntry)
	nameBox = container.NewVBox(nameBox, widget.NewLabel(""))

	priority := getPrioritySlice()
	selectPriority := widget.NewSelect(priority, func(s string) {})
	selectPriority.SetSelected(priority[0])

	okButton := widget.NewButton("Ok", func() {
		if nameEntry.Text == "" {
			nameEntry.SetPlaceHolder("Сюда название, пожалуйста")
			return
		}
		var t taskType
		t.Init(nameEntry.Text, taskPriority(selectPriority.SelectedIndex()))
		Tasks = append(Tasks, t)
		tb.Add(t.Box)
		// file
		pbar.Max = pbar.Max + 1
		pbar.Refresh()
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 35)), okButton)
	selectBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 35)), selectPriority)
	box := container.NewBorder(nameBox, nil, nil, buttonBox, selectBox)

	w.SetContent(box)
	w.Show()
}

func readTasksFromFile() []taskType {
	var tasks []taskType

	// for i := 0; i <= 2; i++ {
	// 	var temp taskType
	// 	temp.Init("aaa", Impotant)
	// 	tasks = append(tasks, temp)
	// }
	// for i := 0; i <= 2; i++ {
	// 	var temp taskType
	// 	temp.Init("bbb", ComputerStuff)
	// 	tasks = append(tasks, temp)
	// }

	return tasks
}
