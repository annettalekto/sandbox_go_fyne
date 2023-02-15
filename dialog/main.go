package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/*
	Далоговые окна FYNE
	https://developer.fyne.io/explore/dialogs
*/

func main() {
	a := app.New()
	w := a.NewWindow("Fyne test...")
	w.Resize(fyne.NewSize(300, 300))

	out := widget.NewLabel("")

	// Form - FormItem дилоговое окно с виджетом // todo проверка?
	var items []*widget.FormItem
	input := widget.NewEntry()
	items = append(items, newFormItem("Введите имя:", "(тут вы пишете ваше имя))", input))

	btnForm := widget.NewButton("Познакомиться", func() {
		out.Text = ""
		input.SetText("")
		out.Refresh()
		d := dialog.NewForm("Dialog FormItem", "Хорошо", "Ну уж нет", items, func(ok bool) {
			if ok {
				out.SetText("Привет, " + input.Text + "!")
			} else {
				out.SetText("Ну и ладно...")
			}
		}, w)
		d.Show()
	})

	// Information - обычное информационное окно
	btnInf := widget.NewButton("Прочитать сообщение", func() {
		dialog.ShowInformation("Dialog Information", "Это пример диалогового окна Information", w)
	})

	// Custom - сообщение с объектом canvas (круг, линия, цветной текст)
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	// circle := canvas.NewCircle(color.White)
	// circle.StrokeWidth = 10
	// circle.StrokeColor = red
	// line := canvas.NewLine(red)
	msg := canvas.NewText("Неверное значение!", red)
	btn := widget.NewButton("Сообщение об ошибке", func() {
		d := dialog.NewCustom("Dialog Custom", "Ок", msg, w)
		d.Show()
	})

	// Confirm - диалоговое окно формата Да/нет
	btnConfirm := widget.NewButton("Закрыть без сохранения", func() {
		d := dialog.NewConfirm("Dialog Confirm", "Точно не сохранять?", func(ok bool) {
			if ok {
				w.Close()
			}
		}, w)
		d.SetDismissText("Отмена") // переписываем стандарнтое NO and Yes
		d.SetConfirmText("Да")
		d.Show()
	})

	bottom := container.NewHBox(layout.NewSpacer(), btnConfirm)
	box := container.NewVBox(out, btnInf, btn)
	w.SetContent(container.NewBorder(btnForm, bottom, nil, nil, box))

	w.ShowAndRun()
}

func newFormItem(text, hint string, w fyne.CanvasObject) *widget.FormItem {
	fi := widget.NewFormItem(text, w)
	fi.HintText = hint
	return fi
}
