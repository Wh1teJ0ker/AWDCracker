package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("AWDCracker")

	w.SetContent(container.NewVBox(
		widget.NewLabel("Welcome to AWDCracker"),
		widget.NewButton("Start Scan", func() {
			// 调用扫描功能
		}),
	))

	w.ShowAndRun()
}
