package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func ProgressBarForm() {
	w := fyne.CurrentApp().NewWindow("ProgressBar")
	w.Resize(fyne.NewSize(400, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	// можно просто создать и изменять только Value:
	pb1 := widget.NewProgressBar()
	// pb1.Min = 0 // минимальное значение Value, от которого прогресс бар начинает заполнение
	pb1.Max = 10  // максисальное значение Value, при котром прогресс бар заполниться на 100% (по умолчанию 1.0)
	pb1.Value = 5 // пользовательское число

	// может быть ошибка:
	pb11 := widget.NewProgressBar()
	pb11.Max = 0   // если Max = 0, что по идее не верно, но может быть, если динамически формировать Max
	pb11.Value = 0 // тогда не верно выводиться процент

	// для исправления можно сделать как pb12
	// подробнее в miniprogs/todolist
	pb12 := widget.NewProgressBar()
	pb12.Max = 0.1   // если задач на сегодня нет, то вместо нуля можно установить небольшое значение
	pb12.Min = 1     // это значение не позволит прогрессбар заполняться, когда задач будет меньше 0 (и не будет отображаться как 1.1 а будет 1)
	pb12.Value = 0.1 // тут 0 или 0.1 покажет пустой или заполненный прогресс бар, например заполненный для того чтобы показать "задач нет", те все выполненно

	// связанный с переменной
	done := binding.NewFloat()
	done.Set(0.1)
	pb2 := widget.NewProgressBarWithData(done)
	pb2.Max = 0.1
	pb2.Value = 0 // вот это вообще не надо использовать, изменять только переменную done. тут видно, что ноль не влияет

	// бесконечный - не знаю зачем
	pb21 := widget.NewProgressBarInfinite()
	// pb21.Stop()
	// pb21.Stop()
	// ok := pb21.Running()

	b := container.NewVBox(pb1, pb11, pb12, pb2, pb21)

	w.SetContent(b)
	w.Show()
}
