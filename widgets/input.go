package main

import (
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

	w.SetContent(container.NewVBox(selectEntry))

	w.Show()
}
