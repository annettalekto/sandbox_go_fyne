package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InputForm() {
	w := fyne.CurrentApp().NewWindow("Input")
	w.Resize(fyne.NewSize(300, 300))
	w.CenterOnScreen()

	// выпадающее меню + entry возможность вписать
	selectEntry := widget.NewSelectEntry([]string{
		"Вариант 1",
		"Вариант 2",
		"Вариант 3",
	})
	selectEntry.PlaceHolder = "выбрать"

	// выбор из выпадающего меню
	slct := widget.NewSelect([]string{"Выбор 1", "Выбор 2", "Выбор 3", "Выбор 4"}, func(s string) { fmt.Println("selected", s) })

	// квадратики
	checkGroup := widget.NewCheckGroup([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, func(s []string) { fmt.Println("выбран квадратик", s) })
	checkGroup.Horizontal = true

	// кружочки
	radio := widget.NewRadioGroup([]string{"Вариант 1", "Вариант 2", "Вариант 3", "Вариант 4"}, func(s string) { fmt.Println("выбран кружок", s) })
	radio.Horizontal = true

	// ползунок
	lab := widget.NewLabel("0")
	sl := widget.NewSlider(0, 1000)
	sl.OnChanged = func(f float64) {
		lab.SetText(fmt.Sprintf("%.0f", f))
	}

	w.SetContent(container.NewVBox(selectEntry, slct, checkGroup, radio, lab, sl))

	w.Show()
}
