package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()                           // one application should have and app.New()
	mainWindow := a.NewWindow("Hello World") // at least has one window

	// make UI
	output, entry, btn := myApp.makeUI()

	// put some content inside main_window
	// container vbox
	mainWindow.SetContent(container.NewVBox(output, entry, btn))
	mainWindow.Resize(fyne.Size{Width: 500, Height: 500})
	mainWindow.ShowAndRun() // run an event loop (listen for events)
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello, World!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})

	btn.Importance = widget.HighImportance

	app.output = output
	return output, entry, btn
}
