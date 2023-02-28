package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Test entry")
	w.Resize(fyne.NewSize(200, 200))

	in1 := widget.NewEntry() // при создании элемента Wrapping = TextTruncate, он самы адекватный, так что можно вообще не прописывать
	// in1.Wrapping = fyne.TextTruncate // появляется полоса прокрутки при заполнении больше ширины виджита
	// in1.Wrapping = fyne.TextWrapOff // окно проги разъезжается при заполнени виджета. тупо выглядит
	// in1.Wrapping = fyne.TextWrapBreak // для мультистрочного виджета (тут просто прокрутка)
	// in1.Wrapping = fyne.TextWrapWord // для мультистрочного виджета (тут просто прокрутка)

	in2 := widget.NewMultiLineEntry() // при создании элемента Wrapping = TextTruncate
	// in2.Wrapping = fyne.TextTruncate  // при заполнении виджита появляются вертикальная и горизонтальная полосы прокрутки
	// in2.Wrapping = fyne.TextWrapOff // окно проги разъезжается
	// in2.Wrapping = fyne.TextWrapBreak // перенос того что неуместилось на другую строку - прям по стредине слова (горизонтальной полосы прокрутки нет)
	in2.Wrapping = fyne.TextWrapWord // самый удобный для редактора: перенос по словам (горизонтальной полосы прокрутки нет)

	box := container.NewVBox(in1, in2)

	w.SetContent(box)
	w.ShowAndRun()
}
