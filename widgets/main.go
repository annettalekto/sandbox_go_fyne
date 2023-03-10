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
		LabelWidgets  = "Label метки"
		ButtonWidgets = "Кнопки"
		SelectWidgets = "Выбор"
	)

	// лист
	dataForList := []string{LabelWidgets, SelectWidgets, ButtonWidgets}
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

func LabelWidgetsForm() {
	w := fyne.CurrentApp().NewWindow("Fyne Label")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))

	// варианты label
	l1 := widget.NewLabel("Test label: Italic style")
	l1.TextStyle.Italic = true
	l2 := widget.NewLabel("Test label: Monospace style")
	l2.TextStyle.Monospace = true
	l3 := widget.NewLabel("Test label: Bold style")
	l3.TextStyle.Bold = true
	l4 := widget.NewLabel("Test label: Symbol style")
	l4.TextStyle.Symbol = true
	l5 := widget.NewLabel("Test label: TabWidth=1") // todo: tab?
	l5.TextStyle.TabWidth = 1
	l6 := widget.NewLabel("Test label: TabWidth=5")
	l6.TextStyle.TabWidth = 5

	labelBox := container.NewVBox(l1, l2, l3, l4, l5, l6)
	w.SetContent(labelBox)
	w.Show()
}
