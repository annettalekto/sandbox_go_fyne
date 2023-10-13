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
	Max, Value        float64
	ProgressBar       *widget.ProgressBar `json:"-"`
	Box               *fyne.Container     `json:"-"`
}

// не для каждого типа, а для вкладки:
// Notes             string              `json:",omitempty"`

var Goals []goalType
var GoalsBox *fyne.Container
var FileName string = "data.json"

// Init for goalType's progressBar
func (g *goalType) Init(name, description string, max, value float64) {
	g.Name = name
	g.Description = description
	g.Max = max

	text := canvas.NewText("     "+g.Name, color.Black) // без пробелов выходит за прогресс бар слева
	text.TextStyle.Italic = true
	textBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(0, 30)), text)

	g.ProgressBar = widget.NewProgressBar()
	g.ProgressBar.Max = max
	g.ProgressBar.Value = value
	g.ProgressBar.Min = 0
	g.ProgressBar.SetValue(0)

	plusButton := widget.NewButton("  +  ", func() {
		g.ProgressBar.Value++
		g.Value = g.ProgressBar.Value // т.к. *widget.ProgressBar не сохраняется в файл
		g.ProgressBar.Refresh()
	})
	changeButton := widget.NewButton("  ...  ", func() {
		g.ChangeGoalForm()
	})
	buttonBox := container.NewHBox(plusButton, changeButton)
	g.Box = container.NewBorder(nil, nil, textBox, buttonBox, g.ProgressBar)
}

// IncrementProgress прибавить прогресс
func (g *goalType) IncrementProgress() {
	g.Value = g.ProgressBar.Value
	g.ProgressBar.Value++
	g.ProgressBar.Refresh()
}

// ChangeGoalForm форма для изменения парамметров цели
func (g *goalType) ChangeGoalForm() {

	w := fyne.CurrentApp().NewWindow("Изменить")
	w.Resize(fyne.NewSize(400, 190))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	name := canvas.NewText(g.Name, color.Black)
	name.TextSize = 16
	name.TextStyle.Monospace = true
	nameBox := container.NewHBox(layout.NewSpacer(), name, layout.NewSpacer())
	descriptionEntry := widget.NewEntry()
	descriptionEntry.OnChanged = func(s string) {
		g.Description = s
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
		}
	}
	valueEntry.SetPlaceHolder(fmt.Sprintf("%v", g.ProgressBar.Value))
	boxValue := container.NewBorder(nil, nil, widget.NewLabel("Сделано: "),
		widget.NewLabel(fmt.Sprintf("(из %v)", g.ProgressBar.Max)), valueEntry)

	doneButton := widget.NewButton("Завершить", func() {
		if g.Max != g.ProgressBar.Value {

			msg := fmt.Sprintf("Завершение цели \"%s\"", g.Name)
			d := dialog.NewConfirm(msg, "Прогресс не 100% Завершить?", func(ok bool) {
				w.Close()
				if ok {
					Goals = removeGoals(Goals, g.Name)
					GoalsBox.Remove(g.Box)
					// сохранить в отдельный файл завершенные проекты
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
			w.Close()
			if ok {
				Goals = removeGoals(Goals, g.Name)
				GoalsBox.Remove(g.Box)
				// удалить из файла
			}
		}, w)
		d.SetDismissText("Отмена")
		d.SetConfirmText("Да")
		d.Show()
	})
	okButton := widget.NewButton("Ok", func() {
		w.Close()
	})
	buttonBox := container.NewHBox(deleteButton, doneButton, layout.NewSpacer(), okButton)

	goalsBox := container.NewVBox(nameBox, descriptionEntry, boxValue, widget.NewLabel(""), buttonBox)
	w.SetContent(goalsBox)
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
func (g *goalType) IncValue() {
	g.Value++
}
func goalForm() *fyne.Container {

	Goals = append(Goals, readGoalsFromFile()...)
	// GoalsBox = createGoalsBox(Goals)
	GoalsBox = container.NewVBox()
	Goals[0].IncValue()
	Goals[0].IncValue()
	for i := 0; i < len(Goals); i++ {
		GoalsBox.Add(Goals[i].Box)
	}
	addGoalButton := widget.NewButton("Новая цель", func() {
		newGoalForm(GoalsBox)
	})
	testButton := widget.NewButton("Записть файла", func() {
		writeGoalsIntoFile(Goals)
	})
	button := container.NewBorder(nil, nil, testButton, addGoalButton)

	notesEntry := widget.NewMultiLineEntry()
	notesEntry.Wrapping = fyne.TextWrapWord
	notesEntry.OnChanged = func(s string) {
		// todo: ...
	}

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
		w.Close()
	})
	buttonBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 30)), buttonOk) // size
	buttonBox = container.NewBorder(nil, nil, nil, buttonBox, nil)                       // left
	box := container.NewVBox(grid, buttonBox, errorLabel)

	w.SetContent(box)
	w.Show()
}

func createGoalsBox(goals []goalType) *fyne.Container {

	box := container.NewVBox()
	// for _, g := range goals {
	// 	box.Add(g.Box)
	// }
	for i := 0; i < len(goals); i++ {
		box.Add(goals[i].Box)
	}

	return box
}

func readGoalsFromFile() []goalType {
	var goals []goalType
	var goal1, goal2, goal3 goalType
	goal1.Init("Читать ITM:", "", 300, 5)
	goal2.Init("Читать ENG:", "", 1300, 5)
	goal3.Init("Перебрать тетради:", "", 15, 5)
	goals = append(goals, goal1, goal2, goal3)
	return goals
}
func readGoalsFromFile1() ([]goalType, error) {
	filename, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()

	data, err := io.ReadAll(filename)
	if err != nil {
		log.Fatal(err)
	}

	var goals []goalType
	jsonErr := json.Unmarshal(data, &goals)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return goals, err
}

func writeGoalsIntoFile(g []goalType) error {
	filename, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()

	jsData, err := json.MarshalIndent(g, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(FileName, jsData, 0777)

	return err
}
