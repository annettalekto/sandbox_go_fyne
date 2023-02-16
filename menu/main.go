package main

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne test")
	w.Resize(fyne.NewSize(400, 300))

	menu := fyne.NewMainMenu(
		fyne.NewMenu("Файл",
			fyne.NewMenuItem("Тема", func() { changeTheme(a) }),
		),

		fyne.NewMenu("Справка",
			fyne.NewMenuItem("Посмотреть справку", func() { aboutHelp() }),
			// fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("О программе", func() { abautProgramm() }),
		),
	)
	w.SetMainMenu(menu)

	go func() { // простите
		sec := time.NewTicker(1 * time.Second)
		for range sec.C {
			for _, item := range menu.Items[0].Items {
				if strings.Contains(item.Label, "Quit") {
					item.Label = "Выход"
					menu.Refresh()
				}
			}
		}
	}()

	w.SetContent(widget.NewLabel("Пример меню в приложении Fyne"))
	w.ShowAndRun()
}

var currentTheme bool // читать/сохранять из файла

func changeTheme(a fyne.App) {
	currentTheme = !currentTheme

	if currentTheme {
		a.Settings().SetTheme(theme.DarkTheme())
	} else {
		a.Settings().SetTheme(theme.LightTheme())
	}
}

func abautProgramm() {
}

func aboutHelp() {
}
