package main

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	//url := "https://forecast.weather.gov/MapClick.php?lat=41.25&lon=-81.44&unit=0&lg=english&FcstType=dwml"
	w := a.NewWindow("Goplex Client")
	//	fd := ""
	//	content := widget.NewLabel(fd)
	//	scrollableContent := container.NewVScroll(content)
	//	scrollableContent.SetMinSize(fyne.NewSize(400, 400))
	//	dialog.ShowCustom("Weather Forecast", "Close", scrollableContent, w)

	exitButton := widget.NewButton("Exit", func() {
		os.Exit(0)
	})

	w.SetContent(container.NewVBox(
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
