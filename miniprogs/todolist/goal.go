package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// goalType data
type goalType struct {
	Name, Description string
	Max, Value        int16
	TextOnProgressBar *canvas.Text        `json:"-"`
	ProgressBar       *widget.ProgressBar `json:"-"`
	Box               *fyne.Container     `json:"-"`
	Start             time.Time
	// Notes             string              `json:",omitempty"`
	Deleted bool
}

// type goalBoxType struct {
// 	box *fyne.Container
// 	i   int
// }

var Goals = make([]goalType, 0, 10) // todo: зачем тут емкость?
// var GoalsBox = container.NewVBox()
var MainForm = container.NewVBox()

// var GoalBoxes = make([]goalBoxType, 0, 10)

var goalJSON string = "data\\goal.json"
var GoalsNoteFile string = "data\\goal_notes.txt"
var HistoryFile string = "data\\history.txt"

// Init инициализирует все элементы переменной типа goalType, прописывает имена и тексты
func (g *goalType) Init(goalsBox *fyne.Container, name, description string, max, value int16) {
	g.Name = name
	g.Description = description
	g.Max = max
	g.Value = value
	g.Start = time.Now()
	// writeHistoryFile("Create", g.Name, g.Description, g.Start, g.Value, g.Max) // сохраняется по нескольку раз тк вызывается и при создании вновь и при загрузки

	g.TextOnProgressBar = canvas.NewText("Goal", color.Black)
	g.TextOnProgressBar.Text = fillOutProgressBar(g.Name, g.Value, g.Max)
	g.TextOnProgressBar.TextStyle.Italic = true
	textGoalsBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(0, 30)), g.TextOnProgressBar)

	g.ProgressBar = widget.NewProgressBar()
	g.ProgressBar.Max = float64(max)
	g.ProgressBar.Value = float64(value)
	g.ProgressBar.SetValue(float64(value))

	plusButton := widget.NewButton("  +  ", func() {
		g.ProgressBar.Value++
		g.ProgressBar.Refresh()
		g.Value = int16(g.ProgressBar.Value)
		g.TextOnProgressBar.Text = fillOutProgressBar(g.Name, g.Value, g.Max)
		g.TextOnProgressBar.Refresh()
		writeGoalsIntoFile(Goals)
	})
	changeButton := widget.NewButton("  ...  ", func() {
		g.ChangeGoalForm(goalsBox)
	})
	buttonBox := container.NewHBox(plusButton, changeButton)
	g.Box = container.NewBorder(nil, nil, textGoalsBox, buttonBox, g.ProgressBar)
}

func fillOutProgressBar(name string, val, max int16) string {
	return fmt.Sprintf("     %s (%v из %v)", name, val, max) // без пробелов выходит за прогресс бар слева
}

func Refresh() {
	for i := range Goals {
		Goals[i].TextOnProgressBar.Refresh()
	}
}

// ChangeGoalForm форма для изменения параметров цели, удаления и завершения
func (g *goalType) ChangeGoalForm(goalsBox *fyne.Container) {

	w := fyne.CurrentApp().NewWindow("Изменить")
	w.Resize(fyne.NewSize(400, 270))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	name := canvas.NewText(g.Name, color.Black)
	name.TextSize = 18
	name.TextStyle.Monospace = true
	nameBox := container.NewHBox(layout.NewSpacer(), name, layout.NewSpacer())
	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.Wrapping = fyne.TextWrapWord
	descriptionEntry.OnChanged = func(s string) {
		g.Description = s
		writeGoalsIntoFile(Goals)
	}
	if g.Description == "" {
		descriptionEntry.SetPlaceHolder("Введите примечание...")
	} else {
		descriptionEntry.SetText(g.Description)
	}

	valueEntry := newNumericalEntry()
	valueEntry.OnChanged = func(s string) {
		v, err := strconv.Atoi(s)
		if err == nil {
			g.ProgressBar.Value = float64(v)
			g.Value = int16(g.ProgressBar.Value)
			g.ProgressBar.Refresh()
			g.TextOnProgressBar.Text = fillOutProgressBar(g.Name, g.Value, g.Max)
			g.TextOnProgressBar.Refresh()
			writeGoalsIntoFile(Goals)
		}
	}
	valueEntry.Entry.FocusGained() //курсор есть, но не вводит цифры, пока не ткнешь мышкой
	valueEntry.SetText(fmt.Sprintf("%v", g.ProgressBar.Value))

	boxValue := container.NewBorder(nil, nil, widget.NewLabel("Сделано: "),
		widget.NewLabel(fmt.Sprintf("(из %v)", g.ProgressBar.Max)), valueEntry)

	doneButton := widget.NewButton("Завершить", func() {
		/*if g.Max != int16(g.ProgressBar.Value) {
			msg := fmt.Sprintf("Завершение цели \"%s\"", g.Name)
			d := dialog.NewConfirm(msg, "Прогресс не 100%. Завершить?", func(ok bool) {
				if ok {
					w.Close()
					Goals = removeGoal(Goals, g.Name)
					// GoalsBox.Remove(g.Box)
					goalsBox = makeBox(Goals)
					goalsBox.Refresh()
					Refresh()
					writeGoalsIntoFile(Goals)
					writeHistoryFile("Done", g.Name, g.Description, g.Start, g.Value, g.Max)
				}
			}, w)
			d.SetDismissText("Отмена") // todo: при отмене не откатывается набранное
			d.SetConfirmText("Да")
			d.Show()
		}*/
	})
	deleteButton := widget.NewButton("Удалить", func() {
		msg := fmt.Sprintf("Удаление цели \"%s\"", g.Name)
		d := dialog.NewConfirm(msg, "Точно удалить?", func(ok bool) {
			if ok {
				g.Deleted = true
				fmt.Println(g.Name, g.Deleted)
				w.Close()

				/*
					Напрямую как то можно обратиться??
					todo: как то так
					добавить индекс в слайс
					там прописывать удаление прям в глобальной (может и прибавление)
					в одной форме проверять в цикле метки, доб и удалять. Там где бокс создан, без ссылок
				*/
				for i, gg := range Goals {
					if g.Name == gg.Name {
						Goals[i].Deleted = true
					}
				}
				// Goals = removeGoal(Goals, g.Name) // работает
				// goalsBox.Remove(g.Box)            // не верно, удаляется последний?
				// GoalsBox = makeBox(Goals)

				// goalsBox.Refresh()
				// MainForm.Refresh()
				// w.Canvas().Refresh(goalsBox)
				// w.Content().Refresh()
				// Refresh()
				writeGoalsIntoFile(Goals) // работает
				// writeHistoryFile("Delete", g.Name, g.Description, g.Start, g.Value, g.Max) // не работает, пишет про последний объект
			}
		}, w)
		d.SetDismissText("Отмена") // todo: хм...
		d.SetConfirmText("Да")
		d.Show()
	})
	okButton := widget.NewButton("Ok", func() {
		writeGoalsIntoFile(Goals)
		w.Close()
	})

	buttonBox := container.NewHBox(deleteButton, doneButton, layout.NewSpacer(), okButton)
	dateLabel := widget.NewLabel(fmt.Sprintf("%v", g.Start.Format("02.01.2006")))
	dateLabel.Alignment = fyne.TextAlignTrailing

	box := container.NewVBox(nameBox, descriptionEntry, boxValue, dateLabel, widget.NewLabel(""), buttonBox)
	w.SetContent(box)
	w.Show()
}

func removeGoal(slice []goalType, name string) []goalType {
	if len(slice) == 0 {
		return nil
	}
	pos := 0
	for i, g := range slice {
		if g.Name == name {
			pos = i
			break
		}
	}
	copy(slice[pos:], slice[pos+1:])
	return slice[:len(slice)-1]
}

// func removeBox(rem fyne.CanvasObject) {
// 	rem.
// }

// ----------------------------------------------------------------------------
//
//	goal form
//
// ----------------------------------------------------------------------------

// func makeBox(goals []goalType) *fyne.Container {
// 	box := container.NewVBox()
// 	for _, goal := range goals {
// 		// box.Add(goals[len(goals)-1].Box)
// 		box.Add(goal.Box)
// 	}
// 	return box
// }

func goalForm() *fyne.Container {
	var GoalsBox = container.NewVBox()
	savedGoals, err := readGoalsFromFile()
	if err != nil {
		fmt.Printf("ошибка получения данных json: %v", err)
	}

	for _, saved := range savedGoals {
		Goals = append(Goals, goalType{Name: ""})
		Goals[len(Goals)-1].Init(GoalsBox, saved.Name, saved.Description, saved.Max, saved.Value)

		GoalsBox.Add(Goals[len(Goals)-1].Box) //makeBox работает
	}

	addGoalButton := widget.NewButton("Новая цель", func() {
		newGoalForm(GoalsBox)
	})

	delButton := widget.NewButton("del", func() {
		for _, g := range Goals {
			if g.Deleted {
				Goals = removeGoal(Goals, g.Name) // работает
				GoalsBox.Remove(g.Box)            // не верно, удаляется последний?
				writeGoalsIntoFile(Goals)         // работает
				writeHistoryFile("Delete", g.Name, g.Description, g.Start, g.Value, g.Max)
			}
		}
		if len(Goals) == 0 {
			GoalsBox.RemoveAll()
		}
		GoalsBox.Refresh()
	})

	buttonOnLeftBox := container.NewBorder(nil, nil, nil, addGoalButton)
	MainForm = container.NewVBox(GoalsBox, buttonOnLeftBox, delButton)
	return MainForm
}

// newGoalForm форма для создания новой цели, открывается по нажатию кнопки «Новая цель» на главном экране
func newGoalForm(goalsBox *fyne.Container) {
	w := fyne.CurrentApp().NewWindow("Создать цель")
	w.Resize(fyne.NewSize(500, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	var err error
	var name, description string
	var max int
	errorLabel := widget.NewLabel("...") // для вывода ошибок

	nameStr := "Название цели"
	nameEntry := widget.NewEntry() //todo: можно так же проверять по нажатию enter
	noteStr := "Примечание к цели"
	descriptionEntry := widget.NewEntry()
	maxValueStr := "Максимальноe число задач"
	maxValueEntry := newNumericalEntry()

	grid := container.NewGridWithColumns(2,
		widget.NewLabel(nameStr+": "), nameEntry,
		widget.NewLabel(noteStr+": "), descriptionEntry,
		widget.NewLabel(maxValueStr+": "), maxValueEntry,
	)

	buttonOk := widget.NewButton("OK", func() {
		name = nameEntry.Text
		if name == "" {
			err = fmt.Errorf("поле ввода \"%s\" не может быть пустым", nameStr)
			errorLabel.Text = err.Error() // todo: текс должен умещаться!
			errorLabel.Refresh()
			return
		}
		description = descriptionEntry.Text
		maxStr := maxValueEntry.Text
		if maxStr == "" {
			err = fmt.Errorf("поле ввода \"%s\" не может быть пустым", maxValueStr)
			errorLabel.Text = "Ошибка: " + err.Error()
			errorLabel.Refresh()
			return
		}
		max, err = strconv.Atoi(maxStr)
		if err != nil {
			err = fmt.Errorf("ошибка в поле ввода \"%s\"", maxValueStr)
			errorLabel.Text = "Ошибка: " + err.Error()
			errorLabel.Refresh()
			return
		}
		if max <= 0 {
			err = fmt.Errorf("\"%s\" должно быть меньше нуля", maxValueStr)
			errorLabel.Text = "Ошибка: " + err.Error()
			errorLabel.Refresh()
			return
		}
		if max > math.MaxInt16 { // 32767
			err = fmt.Errorf("\"%s\" больше %v", maxValueStr, math.MaxInt16)
			errorLabel.Text = "Ошибка: " + err.Error()
			errorLabel.Refresh()
			return
		}
		errorLabel.Text = "ок"

		// var g goalType; Тут не работает так, если отдельно создать переменную типа goalType
		// g.Init(name, description, float64(max), 0); ТО все вроде бы появляется где нужно
		// Goals = append(Goals, g); НО не работает плюс к goal.value (в файл пишется 0 всегда)
		Goals = append(Goals, goalType{Name: ""})
		i := len(Goals) - 1 // текущий элемент (после добавления)
		Goals[i].Init(goalsBox, name, description, int16(max), 0)
		goalsBox.Add(Goals[i].Box) // todo: a зачем через ссылку можно на прямую... makeBox
		writeGoalsIntoFile(Goals)

		// todo: тут немного криво, то имена, то через слайс, но запихиванть в Init, будут дубликаты
		writeHistoryFile("Create", Goals[i].Name, Goals[i].Description, Goals[i].Start, 0, Goals[i].Max)
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 30)), buttonOk) // size
	buttonBox = container.NewBorder(nil, nil, nil, buttonBox, nil)                       // left
	box := container.NewVBox(grid, errorLabel, buttonBox)

	w.SetContent(box)
	w.Show()
}

func readGoalsFromFile() ([]goalType, error) {
	file, err := os.Open(goalJSON)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var saved []goalType
	if err = json.Unmarshal(data, &saved); err != nil {
		log.Fatal(err) // сбой демаршалинга
	}

	return saved, err
}

func writeGoalsIntoFile(g []goalType) error {
	file, err := os.Open(goalJSON)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsData, err := json.MarshalIndent(g, "", "	")
	if err != nil {
		log.Fatal(err) // сбой маршалинга
	}
	err = os.WriteFile(goalJSON, jsData, 0777)

	return err
}

// Записи: создание, удаление, завершение
func writeHistoryFile(prefix, name, description string, t time.Time, val, max int16) error {
	text := fmt.Sprintf("%v %v goal: %v — %v (max: %v, done: %v)\n", t.Format("02.01.2006 15:04:05"), prefix, name, description, max, val)
	f, err := os.OpenFile(HistoryFile, os.O_RDWR|os.O_APPEND, os.ModeType)
	if err != nil {
		fmt.Printf("Ошибка записи файла HistoryFile: %v", err)
	}
	defer f.Close()
	_, err = f.WriteString(text)
	if err != nil {
		fmt.Printf("Ошибка записи строки в HistoryFile: %v", err)
	}

	return err
}
