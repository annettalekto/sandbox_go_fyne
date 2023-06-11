package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TabForm() {
	w := fyne.CurrentApp().NewWindow("Input")
	w.Resize(fyne.NewSize(300, 300))
	w.CenterOnScreen()

	tabs := container.NewAppTabs(
		container.NewTabItem("вариант 1", tab1()),
		container.NewTabItem("вариант 2", tab2()),
	)
	tabs.SetTabLocation(container.TabLocationBottom)

	w.SetContent(tabs) // тут не использовать контейнер

	w.Show()
}

func tab1() *fyne.Container {
	return container.NewVBox(widget.NewLabel("Вкладка 1"))
}
func tab2() *fyne.Container {
	return container.NewVBox(widget.NewLabel("Вкладка 2"))
}
