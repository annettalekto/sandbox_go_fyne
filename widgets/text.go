package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LabelWidgetsForm() {
	w := fyne.CurrentApp().NewWindow("Fyne Label")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	// варианты label
	// NewLabelWithData
	// NewLabelWithStyle
	l1 := widget.NewLabel("Test label: Italic style, Alignment = center")
	l1.TextStyle.Italic = true
	l1.Alignment = fyne.TextAlignCenter
	l2 := widget.NewLabel("Test label: Monospace style, Alignment = TextAlignLeading")
	l2.TextStyle.Monospace = true
	l2.Alignment = fyne.TextAlignLeading
	l3 := widget.NewLabel("Test label: Bold style, Alignment = TextAlignTrailing")
	l3.TextStyle.Bold = true
	l3.Alignment = fyne.TextAlignTrailing
	l4 := widget.NewLabel("Test label: Symbol style")
	l4.TextStyle.Symbol = true
	l5 := widget.NewLabel("Test label: TabWidth=1") // todo: tab?
	l5.TextStyle.TabWidth = 1
	l6 := widget.NewLabel("Test label: TabWidth=5")
	l6.TextStyle.TabWidth = 5

	// color.NRGBA{R: 214, G: 55, B: 55, A: 255}
	getStringColor := func(cl color.NRGBA) string {
		return fmt.Sprintf("Color text (rgba: %X %X %X %X)", cl.R, cl.G, cl.B, cl.A)
	}

	red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	t1 := canvas.NewText(getStringColor(red), red)
	t1.TextSize = 16
	t1.TextStyle.Bold = true

	orange := color.NRGBA{R: 255, G: 50, B: 20, A: 255}
	t2 := canvas.NewText(getStringColor(orange), orange)
	t2.TextSize = 18
	t2.TextStyle.Italic = true

	jellow := color.NRGBA{R: 255, G: 230, B: 5, A: 255}
	t3 := canvas.NewText(getStringColor(jellow), jellow)
	t3.TextSize = 20
	t3.TextStyle.Monospace = true

	green := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	t4 := canvas.NewText(getStringColor(green), green)
	t4.TextSize = 22
	t4.TextStyle.Symbol = true

	blue := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	t5 := canvas.NewText(getStringColor(blue), blue)
	t5.TextSize = 24
	t5.TextStyle.TabWidth = 4

	purple := color.NRGBA{R: 184, G: 15, B: 200, A: 255}
	t6 := canvas.NewText(getStringColor(purple), purple)
	t6.TextSize = 26
	t6.TextStyle.TabWidth = 8

	labelBox := container.NewVBox(l1, l2, l3, l4, l5, l6, t1, t2, t3, t4, t5, t6)

	w.SetContent(labelBox)
	w.Show()
}
