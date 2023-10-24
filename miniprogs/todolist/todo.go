package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// taskType data
type taskType struct {
	Name     string
	Done     bool
	Priority int
	Check    *widget.Check   `json:"-"`
	Box      *fyne.Container `json:"-"`
}

var TasksDone binding.Float
var Tasks = make([]taskType, 0, 10)

var taskJSON string = "data\\task.json"
var TasksNoteFile string = "data\\task_notes.txt"

func (t *taskType) Init(name string, priority taskPriority, done bool) {
	v, _ := TasksDone.Get()
	t.Name = name
	t.Priority = int(priority)
	t.Done = done

	t.Check = widget.NewCheck("", func(b bool) {
		TasksDone.Set(v)
		t.Done = b
		if b {
			v += 1
		} else {
			v -= 1
		}
	})

	nameWidget := canvas.NewText(name, getColorOfPriority(priority))
	nameWidget.TextSize = 14
	nameWidget.TextStyle.Monospace = true

	t.Box = container.NewHBox(t.Check, nameWidget)
}

// ----------------------------------------------------------------------------
//
//	task form
//
// ----------------------------------------------------------------------------

func taskForm(t *container.AppTabs) *fyne.Container {
	tb := container.NewGridWithColumns(2)
	TasksDone = binding.NewFloat()

	savedTasks, err := readTasksFromFile()
	if err != nil {
		fmt.Printf("ошибка получения данных json: %v", err)
	}

	for _, saved := range savedTasks {
		Tasks = append(Tasks, taskType{Name: ""})
		Tasks[len(Tasks)-1].Init(saved.Name, taskPriority(saved.Priority), saved.Done)
		if saved.Done {
			td, _ := TasksDone.Get()
			TasksDone.Set(td + 1)
		}

		tb.Add(Tasks[len(Tasks)-1].Box)
	}
	// tb = container.NewGridWithColumns(2)
	// for _, t := range Tasks { // + сортировку и вынести в отд. ф.
	// 	tb.Add(t.Box)
	// }

	pbarInf := widget.NewProgressBarInfinite()
	pbar := widget.NewProgressBarWithData(TasksDone)
	pbar.Max = float64(len(Tasks))
	pbar.Hide()

	addTask := widget.NewButton("Новая задача", func() {
		addTaskForm(tb, pbar)
	})

	cleanTask := widget.NewButton("Удалить отмеченные", func() {
		for i := 0; i < len(Tasks); {
			t := Tasks[i]
			if t.Check.Checked { // если пункт отмечен, то удалить
				Tasks = removeTask(Tasks, i) // удалить из среза
				tb.Remove(t.Box)             // удалить с формы
				// удалить из файла
			} else {
				i++
			}
		}
		pbar.Max = float64(len(Tasks))
		pbar.Refresh()
		TasksDone.Set(0)
	})

	notesEntry := widget.NewMultiLineEntry()
	notesEntry.Wrapping = fyne.TextWrapWord

	testButton := widget.NewButton("Записть файла", func() { // debug
		writeTasksIntoFile(Tasks)
	})

	buttonBox := container.NewBorder(nil, nil, cleanTask, addTask, testButton)
	tasksBox := container.NewVBox(buttonBox, tb)
	pb := container.NewVBox(pbarInf, pbar)

	tasksBox = container.NewBorder(tasksBox, pb, nil, nil, notesEntry)

	go func() {
		l := len(Tasks)

		sec := time.NewTicker(time.Second / 2)
		for range sec.C {
			if l != len(Tasks) {
				l = len(Tasks)

				if l == 0 {
					pbarInf.Show()
					pbar.Hide()
					tasksBox.Refresh()
				} else {
					pbarInf.Hide()
					pbar.Show()
					tasksBox.Refresh()
				}
				tasksBox.Refresh()
			}
		}
	}()

	return tasksBox
}

func removeTask(slice []taskType, i int) []taskType {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func addTaskForm(tb *fyne.Container, pbar *widget.ProgressBar) {
	//todo: или расположить на главной форме entry
	w := fyne.CurrentApp().NewWindow("Создать")
	w.Resize(fyne.NewSize(400, 130))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	nameEntry := widget.NewEntry()
	nameEntry.FocusGained()
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
		t.Init(nameEntry.Text, taskPriority(selectPriority.SelectedIndex()), false)
		Tasks = append(Tasks, t)
		tb.Add(t.Box)

		// file
		pbar.Max++
		pbar.Refresh()
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 35)), okButton)
	selectBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 35)), selectPriority)
	box := container.NewBorder(nameBox, nil, nil, buttonBox, selectBox)

	w.SetContent(box)
	w.Show()
}

func readTasksFromFile1() []taskType {
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

func readTasksFromFile() ([]taskType, error) {
	f, err := os.Open(taskJSON)
	defer f.Close() // до или после проверки ошибки ? todo:
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var saved []taskType
	if err = json.Unmarshal(data, &saved); err != nil {
		fmt.Println(err)
	}

	return saved, err
}

func writeTasksIntoFile(g []taskType) error {
	f, err := os.Open(taskJSON)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	jsData, err := json.MarshalIndent(g, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(taskJSON, jsData, 0777)

	return err
}

func readTaskNotes() (string, error) {
	in, err := os.ReadFile(TasksNoteFile)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(in)
	return string(in), err
}

func writeTaskNotes(s string) error {
	err := os.WriteFile(TasksNoteFile, []byte(s), 0777)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
