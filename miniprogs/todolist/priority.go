package main

import "image/color"

type taskPriority int

const (
	Housework taskPriority = iota
	ComputerStuff
	AnotherOne
	Impotant
	VeryImpotant
)

var priorityMap = map[taskPriority]string{
	Housework:     "дом. дела",
	ComputerStuff: "комп. дела",
	AnotherOne:    "другое",
	Impotant:      "срочно",
	VeryImpotant:  "очень срочно",
}

var (
	purple = color.NRGBA{R: 184, G: 15, B: 200, A: 255} // 4: очень срочно
	red    = color.NRGBA{R: 255, G: 0, B: 0, A: 255}    // 3: срочно!
	jellow = color.NRGBA{R: 255, G: 230, B: 5, A: 255}  // 2: другое
	blue   = color.NRGBA{R: 0, G: 0, B: 255, A: 255}    // 1: дела за компом (обучение, работа)
	green  = color.NRGBA{R: 0, G: 255, B: 0, A: 255}    // 0: домашние дела
	//orange = color.NRGBA{R: 255, G: 50, B: 20, A: 255}
)

var colorMap = map[taskPriority]color.NRGBA{
	Housework:     green,
	ComputerStuff: blue,
	AnotherOne:    jellow,
	Impotant:      red,
	VeryImpotant:  purple,
}

func getPrioritySlice() []string {
	var priority []string
	for _, s := range priorityMap {
		// widget select какой то баг: обрезаются слова в версии v2.4.0, но в v2.3.4 этого нет
		priority = append(priority, s+"     ")
	}
	return priority
}

func getColorOfPriority(p taskPriority) color.NRGBA {
	return colorMap[p]
}
