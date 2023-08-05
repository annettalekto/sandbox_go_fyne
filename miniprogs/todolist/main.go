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

	// var tabs *container.AppTabs
	tabs := container.NewAppTabs(
		container.NewTabItem("My goals", goalForm()),
		container.NewTabItem("My tasks", taskForm()),
	)
	tabs.SetTabLocation(container.TabLocationBottom)

	// go func() {
	// 	sec := time.NewTicker(5 * time.Second)
	// 	for range sec.C {
	// 		s := fmt.Sprintf("len %v:\t", len(Goals))
	// 		for _, g := range Goals {
	// 			s += fmt.Sprintf("%v, ", g.Name)
	// 		}
	// 		fmt.Println(s)

	// 		// ss := fmt.Sprintf("box %v:\t", len(GoalsBox.Objects))
	// 		// for i, _ := range GoalsBox.Objects {
	// 		// 	ss += fmt.Sprintf("%v, ", i)
	// 		// }
	// 		// fmt.Println(s)
	// 		w.Canvas().Refresh(tabs.Selected().Content)
	// 		tabs.Refresh()
	// 	}
	// }()

	w.SetContent(tabs)
	w.ShowAndRun()
}

/*
todo:
В форме изменить нет проверок на превышение макс
	!тест удалить - работает на половину
	тест завершить
	тест границ на формах создать и изменить

проверять наличие всех файлов (+ можно создавать пустые) где то в начале?
log fatal заменить? если файл пустой - ошибка
os.ReadFile перед - проверка на наличие файла
progress bar: value > max на 1?
оповещать если более 10 целей? см будет ли ош
задачи: убрать кнопку новая задача. что с контейнером??
добавить черный цвет к задачам.
Сортировать по приоритету.
При сохранении в файл ставить дату(?)
добавить напоминалку (сообщение по дате)
будильник? О точно, может сюда помидор?
счетчик раб часовы
Фон - цветом зоны и заголовок выделить
*/
