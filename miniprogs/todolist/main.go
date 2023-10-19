package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO List")
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.SetMaster()

	var tabs *container.AppTabs
	tabs = container.NewAppTabs(
		container.NewTabItem("My goals", goalForm()),
		container.NewTabItem("My tasks", taskForm(tabs)),
	)
	tabs.SetTabLocation(container.TabLocationBottom)

	w.SetContent(tabs)
	w.ShowAndRun()
}

/*
todo:
progress bar: value > max на 1?
работа с файлом
задачи: убрать кнопку новая задача. что с контейнером??
для append функция должна возвращать срез и перезаписывать переменную
добавить черный цвет к задачам.
Сортировать по приоритету.
При сохранении в файл ставить дату(?)
добавить напоминалку (сообщение по дате)
будильник?
Фон - цветом зоны и заголовок выделить
*/
