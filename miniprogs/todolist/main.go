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
если в json файлах нет ничего, падаем?
log fatal заменить? если файл пустой - ошибка
os.ReadFile перед - проверка на наличие файла
кнопка Завершить доделать
progress bar: value > max на 1?
оповещать если более 10 целей? см будет ли ош
задачи: убрать кнопку новая задача. что с контейнером??
добавить черный цвет к задачам.
Сортировать по приоритету.
При сохранении в файл ставить дату(?)
добавить напоминалку (сообщение по дате)
будильник?
Фон - цветом зоны и заголовок выделить
*/
