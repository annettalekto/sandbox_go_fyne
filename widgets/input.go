package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InputForm() {
	w := fyne.CurrentApp().NewWindow("Dialogs")
	w.Resize(fyne.NewSize(300, 300))
	w.CenterOnScreen()

	selectEntry := widget.NewSelectEntry([]string{
		"Вариант 1",
		"Вариант 2",
		"Вариант 3",
	})
	selectEntry.PlaceHolder = "выбрать"

	checkGroup := widget.NewCheckGroup([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, func(s []string) { fmt.Println("выбор", s) })
	checkGroup.Horizontal = true

	w.SetContent(container.NewVBox(selectEntry, checkGroup))

	w.Show()
}
