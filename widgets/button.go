package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ButtonForm() {
	w := fyne.CurrentApp().NewWindow("Button")
	w.Resize(fyne.NewSize(400, 350))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	b1 := widget.NewButton("Кнопка", func() {})
	b2 := widget.NewButtonWithIcon("Ok", theme.ConfirmIcon(), func() { fmt.Println("tapped") })

	b3 := widget.NewButton(" ", nil)
	b3 = &widget.Button{
		Text:       "Warning",
		Importance: widget.WarningImportance,
		OnTapped:   func() { fmt.Println("Warning button") },
	}

	var b4 *widget.Button
	b4 = &widget.Button{
		Text:       "Danger button",
		Importance: widget.DangerImportance,
		OnTapped:   func() { fmt.Println("tapped danger button") },
	}

	box := container.NewVBox(b1, b2, b3, b4, &widget.Button{
		Alignment:     widget.ButtonAlignTrailing,
		IconPlacement: widget.ButtonIconTrailingText,
		Text:          "выравнивание по краю",
		Icon:          theme.ConfirmIcon(),
		OnTapped:      func() { fmt.Println("кнопка нажата") },
	},
		layout.NewSpacer(),
	)

	w.SetContent(box)
	w.Show()
}
