package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EntryForm() {
	w := fyne.CurrentApp().NewWindow("Entry")
	w.Resize(fyne.NewSize(400, 350))
	w.CenterOnScreen()

	e := widget.NewEntry() // при создании элемента Wrapping = TextTruncate, он самы адекватный, так что можно вообще не прописывать
	// in1.Wrapping = fyne.TextTruncate // появляется полоса прокрутки при заполнении больше ширины виджита
	// in1.Wrapping = fyne.TextWrapOff // окно проги разъезжается при заполнени виджета. тупо выглядит
	// in1.Wrapping = fyne.TextWrapBreak // для мультистрочного виджета (тут просто прокрутка)
	// in1.Wrapping = fyne.TextWrapWord // для мультистрочного виджета (тут просто прокрутка)

	eMylty := widget.NewMultiLineEntry() // при создании элемента Wrapping = TextTruncate
	eMylty.SetMinRowsVisible(5)          // при создании, чтобы было 5 строк сразу
	// in2.Wrapping = fyne.TextTruncate  // при заполнении виджита появляются вертикальная и горизонтальная полосы прокрутки
	// in2.Wrapping = fyne.TextWrapOff // окно проги разъезжается
	// in2.Wrapping = fyne.TextWrapBreak // перенос того что неуместилось на другую строку - прям по стредине слова (горизонтальной полосы прокрутки нет)
	eMylty.Wrapping = fyne.TextWrapWord // самый удобный для редактора: перенос по словам (горизонтальной полосы прокрутки нет)
	eMylty.TextStyle.Italic = true
	eMylty.TextStyle.Bold = true

	/*	OnChanged вызывается на каждое изменение (каждый введенный символ)
		s - то, что введено в поле сейчас
	*/
	eMylty.OnChanged = func(s string) {
		fmt.Println("OnChanged: " + s)
	}
	/*	OnSubmitted вызывается когда все ввели и нажали ввод (enter) для однострочного Entry
		Для многострочного вызывается при нажатии shift + enter, тк enter тут новая строка
		s - все строки, что ввели в Еntry
	*/
	eMylty.OnSubmitted = func(s string) {
		fmt.Println("OnSubmitted: " + s) // shift + enter
	}

	// расширение возможностей базового типа
	// ввод только цифор
	eNumber := newNumericalEntry()
	eNumber.Entry.TextStyle.Monospace = true

	box := container.NewVBox(
		widget.NewLabel("Однострочный:"), e,
		widget.NewLabel("Многострочный: "), eMylty,
		widget.NewLabel("Только цифры: "), eNumber,
	)

	// Можно установить или снять фокус с нужного элемента
	// только это не помогает, т.к. начать писать можно только посте того, как мышкой ткнешь на элемент...
	// eMylty.FocusLost()
	eMylty.FocusGained()

	w.SetContent(box)
	w.Show()
}
