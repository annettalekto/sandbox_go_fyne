package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne test")
	w.Resize(fyne.NewSize(200, 200))

	// чекбокс + bool
	sw := binding.NewBool()                                                 // создана переменная типа binding bool
	check := widget.NewCheckWithData("чек-бокс управляется переменной", sw) // привзана к объекту

	in1 := widget.NewEntry()
	in1.OnChanged = func(s string) {
		if strings.Contains(s, "true") {
			sw.Set(true) // теперь меняем только переменную и чекбокс меняется сам
		} else if strings.Contains(s, "false") {
			sw.Set(false)
		}
	}
	box1 := container.NewVBox(widget.NewLabel("Введите true или false:"), in1, check)

	// label + string
	var style fyne.TextStyle
	style.Italic = true
	style.Bold = true
	label := widget.NewLabel("Статус") // объект
	label.TextStyle = style
	str := binding.NewString() // строковая переменная
	label.Bind(str)            // связать

	in2 := widget.NewEntry()
	in2.OnChanged = func(s string) {
		str.Set(s) // изменять только переменную
	}
	box2 := container.NewVBox(widget.NewLabel("Введите строку для метки:"), in2, label)

	box := container.NewVSplit(box1, box2)
	w.SetContent(box)
	w.ShowAndRun()
}
