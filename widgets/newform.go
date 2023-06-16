package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newWindow() {
	w := fyne.CurrentApp().NewWindow("NewWindow")
	w.Resize(fyne.NewSize(300, 300))
	w.CenterOnScreen()

	lebel := widget.NewLabel("New window")
	box := container.NewBorder(nil, nil, nil, nil, lebel)

	w.SetContent(box) // тут не использовать контейнер

	w.Show()
}
