package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"

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
	Max, Value        float64             // todo: может быть дробным?
	TextPB            *canvas.Text        `json:"-"`
	ProgressBar       *widget.ProgressBar `json:"-"`
	Box               *fyne.Container     `json:"-"`
	// Notes             string              `json:",omitempty"`
}

var Goals = make([]goalType, 0, 10)
var GoalsBox = container.NewVBox()

var goalJSON string = "data\\goal.json"
var GoalsNoteFile string = "data\\goal_notes.txt"

// Init for goalType's progressBar
func (g *goalType) Init(name, description string, max, value float64) {
	g.Name = name
	g.Description = description
	g.Max = max
	g.Value = value

	g.TextPB = canvas.NewText("Goal", color.Black)

	g.TextPB.Text = fillOutPB(g.Name, g.Value, g.Max)
	g.TextPB.TextStyle.Italic = true
	textGoalsBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(0, 30)), g.TextPB)

	g.ProgressBar = widget.NewProgressBar()
	g.ProgressBar.Max = max
	g.ProgressBar.Value = value
	g.ProgressBar.SetValue(value)

	plusButton := widget.NewButton("  +  ", func() {
		g.ProgressBar.Value++
		g.Value = g.ProgressBar.Value // т.к. *widget.ProgressBar не сохраняется в файл
		g.ProgressBar.Refresh()
		g.TextPB.Text = fillOutPB(g.Name, g.Value, g.Max)
		g.TextPB.Refresh()
		writeGoalsIntoFile(Goals)
	})
	changeButton := widget.NewButton("  ...  ", func() {
		g.ChangeGoalForm()
	})
	buttonBox := container.NewHBox(plusButton, changeButton)
	g.Box = container.NewBorder(nil, nil, textGoalsBox, buttonBox, g.ProgressBar)

}

func fillOutPB(name string, val, max float64) string {
	return fmt.Sprintf("     %s (%.0f из %.0f)", name, val, max) // без пробелов выходит за прогресс бар слева
}

// // IncrementProgress прибавить прогресс
// func (g *goalType) IncrementProgress() {
// 	g.Value = g.ProgressBar.Value
// 	g.ProgressBar.Value++
// 	g.ProgressBar.Refresh()
// 	writeGoalsIntoFile(Goals)
// }

// ChangeGoalForm форма для изменения параметров цели
func (g *goalType) ChangeGoalForm() {

	w := fyne.CurrentApp().NewWindow("Изменить")
	w.Resize(fyne.NewSize(400, 230))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	name := canvas.NewText(g.Name, color.Black)
	name.TextSize = 16
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
			g.Value = g.ProgressBar.Value
			g.ProgressBar.Refresh()
			g.TextPB.Text = fillOutPB(g.Name, g.Value, g.Max)
			g.TextPB.Refresh()
			writeGoalsIntoFile(Goals)
		}
	}
	valueEntry.Entry.FocusGained() //курсор есть, но не вводит цифры, пока не ткнешь мышкой
	valueEntry.Entry.OnSubmitted = func(s string) {
		selectAll()
		fmt.Println("onsubmitted")
	}
	valueEntry.SetText(fmt.Sprintf("%v", g.ProgressBar.Value))

	boxValue := container.NewBorder(nil, nil, widget.NewLabel("Сделано: "),
		widget.NewLabel(fmt.Sprintf("(из %v)", g.ProgressBar.Max)), valueEntry)

	doneButton := widget.NewButton("Завершить", func() {
		// сделать не активной пока не будет 100%?
		// добавить файл завершенных проектов?
		if g.Max != g.ProgressBar.Value {

			msg := fmt.Sprintf("Завершение цели \"%s\"", g.Name)
			d := dialog.NewConfirm(msg, "Прогресс не 100% Завершить?", func(ok bool) {
				w.Close()
				if ok {
					Goals = removeGoals(Goals, g.Name)
					GoalsBox.Remove(g.Box)
					writeGoalsIntoFile(Goals)
				}
			}, w)
			d.SetDismissText("Отмена")
			d.SetConfirmText("Да")
			d.Show()
		}
	})
	deleteButton := widget.NewButton("Удалить", func() {
		msg := fmt.Sprintf("Удаление цели \"%s\"", g.Name)
		d := dialog.NewConfirm(msg, "Точно удалить?", func(ok bool) {
			if ok {
				Goals = removeGoals(Goals, g.Name)
				GoalsBox.Remove(g.Box)
				writeGoalsIntoFile(Goals)
			}
			w.Close()
		}, w)
		d.SetDismissText("Отмена")
		d.SetConfirmText("Да")
		d.Show()
	})
	okButton := widget.NewButton("Ok", func() {
		writeGoalsIntoFile(Goals)
		w.Close()
	})
	buttonBox := container.NewHBox(deleteButton, doneButton, layout.NewSpacer(), okButton)

	box := container.NewVBox(nameBox, descriptionEntry, boxValue, widget.NewLabel(""), buttonBox)
	w.SetContent(box)
	w.Show()
}

func removeGoals(slice []goalType, name string) []goalType {
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

// ----------------------------------------------------------------------------
//
//	goal form
//
// ----------------------------------------------------------------------------

func goalForm() *fyne.Container {

	savedGoals, err := readGoalsFromFile()
	if err != nil {
		fmt.Printf("ошибка получения данных json: %v", err)
	}

	for _, saved := range savedGoals {
		Goals = append(Goals, goalType{Name: ""})
		Goals[len(Goals)-1].Init(saved.Name, saved.Description, saved.Max, saved.Value)

		GoalsBox.Add(Goals[len(Goals)-1].Box)
	}

	addGoalButton := widget.NewButton("Новая цель", func() {
		newGoalForm(GoalsBox)
	})

	notesEntry := widget.NewMultiLineEntry()
	notesEntry.Wrapping = fyne.TextWrapWord
	s, _ := readGoalNotes()
	notesEntry.Text = s
	notesEntry.OnChanged = func(s string) {
		writeGoalNotes(s)
	}

	// testButton := widget.NewButton("Запись файла", func() { // debug
	// 	writeGoalsIntoFile(Goals)
	// })
	button := container.NewBorder(nil, nil, nil, addGoalButton)

	box := container.NewVBox(GoalsBox, button)

	// go func() {
	// 	l := len(Goals)
	// 	sec := time.NewTicker(time.Second / 2)
	// 	for range sec.C {
	// 		if l != len(Goals) {
	// 			l = len(Goals)

	// 			box.Refresh()
	// 		}
	// 	}
	// }()

	return container.NewBorder(box, nil, nil, nil, notesEntry)
}

func newGoalForm(goalsBox *fyne.Container) {
	w := fyne.CurrentApp().NewWindow("Создать")
	w.Resize(fyne.NewSize(500, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	var err error
	var name, description string // todo: как передать данные
	var max int
	errorLabel := widget.NewLabel("...") // вывод ошибок

	nameStr := "Название"
	nameEntry := widget.NewEntry()
	noteStr := "Примечание"
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
			err = fmt.Errorf(fmt.Sprintf("Поле ввода \"%s\" не может быть пустым", nameStr))
			errorLabel.Text = err.Error()
			errorLabel.Refresh()
			return
		}
		description = descriptionEntry.Text
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
		g.Init(name, description, float64(max), 0)
		Goals = append(Goals, g)
		goalsBox.Add(g.Box)
		writeGoalsIntoFile(Goals)
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 30)), buttonOk) // size
	buttonBox = container.NewBorder(nil, nil, nil, buttonBox, nil)                       // left
	box := container.NewVBox(grid, buttonBox, errorLabel)

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

func readGoalNotes() (string, error) {
	in, err := os.ReadFile(GoalsNoteFile)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(in)
	return string(in), err
}

func writeGoalNotes(s string) error {
	err := os.WriteFile(GoalsNoteFile, []byte(s), 0777)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
