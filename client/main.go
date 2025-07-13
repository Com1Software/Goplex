package main

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	weather "github.com/Com1Software/Go-Weather"
)

func main() {
	a := app.New()
	url := "https://forecast.weather.gov/MapClick.php?lat=41.25&lon=-81.44&unit=0&lg=english&FcstType=dwml"
	w := a.NewWindow("Weather App")
	time := widget.NewLabel("Current Time :" + weather.GetWeather(url, 0))
	loc := widget.NewLabel("Location :" + weather.GetWeather(url, 1))
	temp := widget.NewLabel("Current Temperature :" + weather.GetWeather(url, 3))
	cond := widget.NewLabel("Current Conditions :" + weather.GetWeather(url, 5))
	sw := widget.NewLabel("Sustained Wind :" + weather.GetWeather(url, 7))
	wg := widget.NewLabel("Wind Gusts :" + weather.GetWeather(url, 8))
	bar := widget.NewLabel("Barometer :" + weather.GetWeather(url, 9))
	hum := widget.NewLabel("Humidity :" + weather.GetWeather(url, 10))
	dp := widget.NewLabel("Dew Point :" + weather.GetWeather(url, 11))
	vis := widget.NewLabel("Visibility :" + weather.GetWeather(url, 13))
	wc := widget.NewLabel("Wind Chill :" + weather.GetWeather(url, 14))
	hw := widget.NewLabel("Hazard Warning :" + weather.GetWeather(url, 2))

	forcastButton := widget.NewButton("Forecast", func() {
		fd := weather.GetWeather(url, 15) + " " + weather.GetWeather(url, 16) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 18)) + "\n\n"
		fd = fd + weather.GetWeather(url, 19) + " " + weather.GetWeather(url, 20) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 22)) + "\n\n"
		fd = fd + weather.GetWeather(url, 23) + " " + weather.GetWeather(url, 24) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 26)) + "\n\n"
		fd = fd + weather.GetWeather(url, 27) + " " + weather.GetWeather(url, 28) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 30)) + "\n\n"
		fd = fd + weather.GetWeather(url, 31) + " " + weather.GetWeather(url, 32) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 34)) + "\n\n"
		fd = fd + weather.GetWeather(url, 35) + " " + weather.GetWeather(url, 36) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 38)) + "\n\n"
		fd = fd + weather.GetWeather(url, 39) + " " + weather.GetWeather(url, 40) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 42)) + "\n\n"
		fd = fd + weather.GetWeather(url, 43) + " " + weather.GetWeather(url, 44) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 46)) + "\n\n"
		fd = fd + weather.GetWeather(url, 47) + " " + weather.GetWeather(url, 48) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 50)) + "\n\n"
		fd = fd + weather.GetWeather(url, 51) + " " + weather.GetWeather(url, 52) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 54)) + "\n\n"
		fd = fd + weather.GetWeather(url, 55) + " " + weather.GetWeather(url, 56) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 58)) + "\n\n"
		fd = fd + weather.GetWeather(url, 59) + " " + weather.GetWeather(url, 60) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 62)) + "\n\n"
		fd = fd + weather.GetWeather(url, 63) + " " + weather.GetWeather(url, 64) + "\n "
		fd = fd + wordWrap(weather.GetWeather(url, 66)) + "\n"

		content := widget.NewLabel(fd)
		scrollableContent := container.NewVScroll(content)
		scrollableContent.SetMinSize(fyne.NewSize(400, 400)) // Adjust the size as needed
		dialog.ShowCustom("Weather Forecast", "Close", scrollableContent, w)

	})

	exitButton := widget.NewButton("Exit", func() {
		os.Exit(0)
	})

	w.SetContent(container.NewVBox(
		time,
		loc,
		temp,
		cond,
		sw,
		wg,
		bar,
		hum,
		dp,
		vis,
		wc,
		hw,
		forcastButton,
		exitButton,
	))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

func wordWrap(s string) string {
	max := 40
	xdata := ""
	cl := 0
	words := strings.Split(s, " ")

	for _, word := range words {
		if cl+len(word) > max {
			xdata = strings.TrimSpace(xdata) + "\n"
			cl = 0
		}
		xdata += word + " "
		cl += len(word) + 1
	}
	fmt.Println(xdata)
	return strings.TrimSpace(xdata)
}
