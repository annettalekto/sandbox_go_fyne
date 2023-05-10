package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ButtonForm() {
	w := fyne.CurrentApp().NewWindow("Button")
	w.Resize(fyne.NewSize(400, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	box := container.NewVBox(widget.NewButton("Кнопка", func() {}))

	w.SetContent(box)
	w.Show()
}
