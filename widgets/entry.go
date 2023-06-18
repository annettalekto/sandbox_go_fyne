package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EntryForm() {
	w := fyne.CurrentApp().NewWindow("Entry")
	w.Resize(fyne.NewSize(400, 200))
	w.CenterOnScreen()

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
	in2.TextStyle.Italic = true

	in2.OnChanged = func(s string) {
		fmt.Println("OnChanged")
	}
	in2.TextStyle.Bold = true
	in2.OnSubmitted = func(s string) {
		fmt.Println("OnSubmitted")
	}

	// расширение возможностей базового типа
	// ввод только цифор
	in3 := newNumericalEntry()
	in3.Entry.TextStyle.Monospace = true

	box := container.NewVBox(
		widget.NewLabel("Однострочный:"), in1,
		widget.NewLabel("Многострочный: "), in2,
		widget.NewLabel("Только цифры: "), in3)

	w.SetContent(box)
	w.Show()
}
