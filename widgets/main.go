package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne test")
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.SetMaster()

	var (
		LabelWidgets  = "Текст"
		ButtonWidgets = "Кнопки"
		SelectWidgets = "Выбор"
		ProgressBar   = "Progress Bar"
	)

	// лист
	dataForList := []string{LabelWidgets, SelectWidgets, ButtonWidgets, ProgressBar}
	list := widget.NewList(
		func() int {
			return len(dataForList)
		},
		func() fyne.CanvasObject {
			var style fyne.TextStyle
			style.Monospace = true
			temp := widget.NewLabelWithStyle("temp", fyne.TextAlignLeading, style)
			return temp
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			if i < len(dataForList) {
				o.(*widget.Label).SetText(dataForList[i])
			}
		})

	l1 := widget.NewLabel(" ")
	list.OnSelected = func(id widget.ListItemID) {
		wt := dataForList[id]
		l1.SetText(fmt.Sprintf("Выбран элемент «%s»", wt))
		if wt == LabelWidgets {
			LabelWidgetsForm()
		}
		if wt == ProgressBar {
			ProgressBarForm()
		}
	}
	l2 := widget.NewLabel(" ")
	list.OnUnselected = func(id widget.ListItemID) {
		wt := dataForList[id]
		l2.SetText(fmt.Sprintf("Снято выделение с элемента «%s»", wt))
	}

	box := container.NewVBox(l1, l2)
	w.SetContent(container.NewHSplit(list, box))
	w.ShowAndRun()
}

func ProgressBarForm() {

	w := fyne.CurrentApp().NewWindow("Progress Bar") // CurrentApp!
	w.Resize(fyne.NewSize(400, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	// чтобы создать нужно:
	pb1 := widget.NewProgressBar()
	pb1.Max = 10  // всего этапов
	pb1.Value = 5 // сколько завершено этапов

	// может быть ошибка:
	pb11 := widget.NewProgressBar()
	pb11.Max = 0   // если Max = 0, что по идее не верно, но может быть если динамически формировать этапы
	pb11.Value = 0 // может не верно выводиться процент

	// для исправления можно сделать как pb12
	// подробнее в miniprogs/todolist
	pb12 := widget.NewProgressBar()
	pb12.Max = 0.1   // если задач на сегодня нет, то вместо нуля можно установить небольшое значени
	pb12.Min = 1     // это значение не позволит прогрессбар заполняться когда задач будет больше 0 (не будет отображаться как 1.1 а буде 1), т.е.
	pb12.Value = 0.1 // тут 0 или 0.1 покажет пустой или заполненный прогресс бар, например заполненный для того чтобы показать "задач нет", те все выполненно

	// связанный с переменной
	done := binding.NewFloat()
	done.Set(0.1)
	pb2 := widget.NewProgressBarWithData(done)
	pb2.Max = 0.1
	pb2.Value = 0 // вот это вообще не надо использовать, изменять только переменную done

	// бесконечный
	// pb21 := widget.NewProgressBar()
	// pb21.Max = 5 // всего этапов

	b := container.NewVBox(pb1, pb11, pb12, pb2)

	w.SetContent(b)
	w.Show() // ShowAndRun -- panic!
}
