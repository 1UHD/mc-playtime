package main

import (
	"fmt"
	"mc-playtime/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Minecraft Playtime")
	window.Resize(fyne.NewSize(400, 400))
	window.SetFixedSize(true)

	//playtime tab
	versions, err := tools.Get_versions(tools.Get_path("/Desktop/cloud_desktop/projects/mc-playtime/") + "versions.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	playtimeContent := container.NewVBox()
	for _, version := range versions {
		img := canvas.NewImageFromFile(version.Picture)
		img.SetMinSize(fyne.NewSize(60, 60))
		img.FillMode = canvas.ImageFillContain
		titleLable := widget.NewLabel(version.Title)
		playtimeText := widget.NewLabel("Playtime: Calculating")
		entry := container.NewHBox(img, container.NewVBox(titleLable, playtimeText))
		playtimeContent.Add(entry)
	}

	playtimeScroll := container.NewScroll(playtimeContent)

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Title")

	playtimeEntry := widget.NewEntry()
	playtimeEntry.SetPlaceHolder("Path")

	addButton := widget.NewButton("Add", func() {
		newVersion := tools.Version{
			Path:    playtimeEntry.Text,
			Picture: "placeholder.png",
			Title:   titleEntry.Text,
		}
		versions = append(versions, newVersion)

		img := canvas.NewImageFromFile(newVersion.Picture)
		img.SetMinSize(fyne.NewSize(60, 60))
		titleLabel := widget.NewLabel(newVersion.Title)
		generatedText := widget.NewLabel("Playtime: Calculating")
		entry := container.NewHBox(img, container.NewVBox(titleLabel, generatedText))
		playtimeContent.Add(entry)

		playtimeContent.Refresh()
	})
	addContainer := container.New(layout.NewBorderLayout(nil, nil, nil, addButton), addButton, container.NewVBox(titleEntry, playtimeEntry))

	playtimeTab := container.NewBorder(nil, addContainer, nil, nil, playtimeScroll)

	//console tab
	consoleLogs := widget.NewMultiLineEntry()
	consoleTab := container.NewScroll(consoleLogs)

	tabs := container.NewAppTabs(
		container.NewTabItem("Playtime", playtimeTab),
		container.NewTabItem("Console", consoleTab),
	)

	window.SetContent(tabs)
	window.ShowAndRun()
}
