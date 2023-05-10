package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
		EntryWidgets  = "Ввод"
		DialogWidgets = "Диалоги"
		MenuWidgets   = "Меню"
		BindWidgets   = "Связные"
		ProgressBar   = "ProgressBar"
		ButtonWidgets = "Кнопки" // todo:
		SelectWidgets = "Выбор"  // todo:
	)

	l1 := widget.NewLabel("widget.NewList: список виджетов")
	dataForList := []string{LabelWidgets, MenuWidgets, DialogWidgets, EntryWidgets,
		BindWidgets, ProgressBar, SelectWidgets, ButtonWidgets}

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

	list.OnSelected = func(id widget.ListItemID) {
		wt := dataForList[id]
		l1.SetText(fmt.Sprintf("Выбран элемент «%s»", wt))

		// текст, разный шрифт и цвет
		if wt == LabelWidgets {
			LabelWidgetsForm()
		}

		// ввод или вывод данных в поля
		if wt == EntryWidgets {
			EntryForm()
		}

		if wt == ButtonWidgets {
			ButtonForm()
		}

		if wt == MenuWidgets {
			MenuForm()
		}

		// диалоги, вывод ошибок
		if wt == DialogWidgets {
			DialogForm()
		}

		// связные элементы
		if wt == BindWidgets {
			BindForm()
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
