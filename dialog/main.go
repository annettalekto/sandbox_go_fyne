package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("fyne test")
	w.Resize(fyne.NewSize(300, 300))

	btn := widget.NewButton("Закрыть без сохранения", func() {
		d := dialog.NewConfirm("Вопрос", "Точно не сохранять?", func(ok bool) {
			if ok {
				w.Close()
			}
		}, w)
		d.SetDismissText("Hет") // переписываем стандарнтое NO and Yes
		d.SetConfirmText("Да")
		d.Show()
	})

	box := container.NewHBox(layout.NewSpacer(), btn)
	w.SetContent(container.NewBorder(nil, box, nil, nil))

	w.ShowAndRun()
}
