package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	url := "http://192.168.1.105:8080"
	//url := "http://com1software.com"
	app := "Goplex Client"
	password := "test"
	w := a.NewWindow("Goplex Client")
	fmt.Println("Starting " + app + url + password)
	res := ReadURL(url)
	fmt.Println("Result: " + res)
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

func ReadURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL:", err)

	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: Received non-OK HTTP status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	// Read the entire response body into a byte slice
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
		os.Exit(1)
	}
	// Convert byte slice to string
	bodyString := string(bodyBytes)
	return bodyString
}
