package main

import (
	"fmt"
	"image/color"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type MyData struct {
	City        string
	Temperature string
	Date        string
	Wind        string
	FeelsLike   string
}

var mydata MyData
var temp chan string

func main() {
	temp = make(chan string)
	defer close(temp)

	GetData(temp)
	CreateForm(temp)
}

func GetData(ch chan string) {
	var info []byte
	url := "https://pogoda1.ru/penza/"

	go func() {
		sec := time.NewTicker(1000 * time.Millisecond)
		for range sec.C {

			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			info, err = io.ReadAll(resp.Body)
			resp.Body.Close() // для предотвращения утечки ресурсов
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			// p := fmt.Println
			// pf := fmt.Printf
			// fileName := "D:\\temp\\f1.txt"
			// f, err := os.Create(fileName)
			// if err != nil {
			// 	p("С файлом беда!")
			// 	os.Exit(1)
			// }
			// n, err := f.Write(info)
			// // p(err)
			// if err == nil {
			// 	pf("Записано в файл %v символов", n)
			// } else {
			// 	pf("File writing error: %v", err)
			// }
			// f.Close()

			sl := strings.Split(string(info), "\n")
			for i, s := range sl {
				if strings.Contains(strings.ToLower(s), "погода в пензе") {
					mydata.City = "Пенза"
				} else if strings.Contains(s, "weather-date-select-day") { //<a href="/penza/25-05-2023/#main" class="weather-date-select-day active">
					mydata.Date = getData(s)
				} else if strings.Contains(s, "weather-now-temp") { //<div class="weather-now-temp">+25&deg;</div>
					d := getData(s)
					d = strings.Trim(d, "&deg;")
					mydata.Temperature = d
				} else if mydata.Wind == "" && strings.Contains(s, "<span class=\"wind-amount\">") { // <span class="wind-amount">Северо-восточный, 1 м/с</span>
					mydata.Wind = getData(s)
				} else if strings.Contains(s, "По ощущению") {
					if len(sl) > i+1 {
						d := sl[i+1]
						d = getData(d)
						d = strings.Trim(d, "&deg;")
						mydata.FeelsLike = d
					}
				}
			}
			ch <- mydata.Temperature
		}
	}()
}

func getData(in string) string {
	out := in
	q1 := strings.Index(out, "<")
	q2 := strings.Index(out, ">")
	if q1 == -1 || q2 == -1 {
		return ""
	}
	out = out[q2+1:]

	q1 = strings.Index(out, "<")
	q2 = strings.Index(out, ">")
	if q1 == -1 || q2 == -1 {
		return ""
	}
	out = out[:q1]

	return out
}

func CreateForm(ch chan string) {
	a := app.New()
	w := a.NewWindow("Погода в пензе на сегодня")
	w.Resize(fyne.NewSize(250, 120))
	w.CenterOnScreen()

	colorLightBlue := color.NRGBA{R: 180, G: 204, B: 214, A: 255}
	colorBlue := color.NRGBA{R: 42, G: 63, B: 163, A: 255}
	colorWhite := color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 255}
	colorGray := color.NRGBA{R: 30, G: 35, B: 51, A: 255}

	rect := canvas.NewRectangle(colorLightBlue)
	rect.FillColor = colorLightBlue
	rect.SetMinSize(fyne.NewSize(250, 120))
	rect.Refresh()

	labelCity := canvas.NewText(mydata.City+" "+mydata.Date, colorWhite)
	labelCity.TextSize = 18
	labelCity.TextStyle.Bold = true

	labelWhether := canvas.NewText(mydata.Temperature, colorBlue)
	labelWhether.TextSize = 32
	labelWhether.TextStyle.Monospace = true

	labelFeels := canvas.NewText(mydata.FeelsLike, colorGray)
	labelFeels.TextSize = 14
	labelFeels.TextStyle.Italic = true

	labelWind := canvas.NewText(mydata.Wind, colorGray)
	labelWind.TextSize = 14
	labelWind.TextStyle.Italic = true

	go func() {
		sec := time.NewTicker(1000 * time.Millisecond)
		for range sec.C {
			t, opend := <-temp
			if !opend {
				break
			}
			if len(labelCity.Text) < 2 {
				labelCity.Text = mydata.City + " " + mydata.Date
				labelCity.Refresh()
			}

			labelFeels.Text = "По ощущению: " + mydata.FeelsLike + "°"
			labelFeels.Refresh()

			labelWind.Text = mydata.Wind
			labelWind.Refresh()

			labelWhether.Text = t + "°"
			labelWhether.Refresh()
		}
	}()

	boxText := container.NewVBox(labelCity, labelWhether, labelFeels, labelWind)
	box := container.NewMax(rect, boxText)
	w.SetContent(box)
	w.ShowAndRun()
}

// <a href="https://www.flaticon.com/ru/free-icons/" title="солнце иконки">Солнце иконки от iconixar - Flaticon</a>
