package utils

import (
	"image/color"
	"net/http"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func MakeSuccessMessage(a fyne.App, message string) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	messageWindow := a.NewWindow("Success")
	messageWindow.Resize(fyne.NewSize(200, 100))
	messageWindow.CenterOnScreen()

	messageToShow := canvas.NewText(message, green)
	messageToShow.Alignment = fyne.TextAlignCenter

	messageWindow.SetContent(messageToShow)
	messageWindow.Show()
}

func MakeErrorMessage(a fyne.App, message string) {
	red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	messageWindow := a.NewWindow("Error")
	messageWindow.Resize(fyne.NewSize(200, 100))
	messageWindow.CenterOnScreen()

	messageToShow := canvas.NewText(message, red)
	messageToShow.Alignment = fyne.TextAlignCenter

	messageWindow.SetContent(messageToShow)
	messageWindow.Show()
}

func TestWebhook(a fyne.App, webhook string) bool {
	if webhook == "" {
		return false
	}

	_, err := url.ParseRequestURI(webhook)
	if err != nil {
		MakeErrorMessage(a, "Invalid URL")
		return false
	}

	resp, err := http.Get(webhook)
	if err != nil {
		MakeErrorMessage(a, "Invalid URL")
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}
