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

var overallpt int32 = 0

func update_overall(added_pt int32, overall_label *widget.Label) {
	overallpt = overallpt + added_pt
	overall_label.SetText(fmt.Sprintf("Playtime: %d hours %d minutes %d seconds", overallpt/3600, (overallpt%3600)/60, overallpt%60))
}

func main() {
	app := app.New()
	window := app.NewWindow("Minecraft Playtime")
	window.Resize(fyne.NewSize(400, 400))

	//console tab
	consoleLogs := widget.NewMultiLineEntry()
	consoleTab := container.NewScroll(consoleLogs)

	//playtime tab
	versions := tools.Get_versions()

	playtimeContent := container.NewVBox()

	overallPlaytime := widget.NewLabel("Overall: 0 hours 0 minutes 0 seconds")
	playtimeContent.Add(overallPlaytime)

	for _, version := range versions {
		img := canvas.NewImageFromFile(version.Picture)
		img.SetMinSize(fyne.NewSize(60, 60))
		img.FillMode = canvas.ImageFillContain
		titleLable := widget.NewLabel(version.Title)
		playtimeText := widget.NewLabel("Playtime: Calculating")

		go func(path string, label *widget.Label, console *widget.Entry) {
			playtime := tools.Get_time_for_directory(path, consoleLogs)
			update_overall(playtime, overallPlaytime)
			label.SetText(fmt.Sprintf("Playtime: %d hours %d minutes %d seconds", playtime/3600, (playtime%3600)/60, playtime%60))
		}(version.Path, playtimeText, consoleLogs)

		entry := container.NewHBox(img, container.NewVBox(titleLable, playtimeText))
		playtimeContent.Add(entry)
	}

	playtimeScroll := container.NewScroll(playtimeContent)

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Title")

	playtimeEntry := widget.NewEntry()
	playtimeEntry.SetPlaceHolder("Path")

	addButton := widget.NewButton("Add", func() {
		if titleEntry.Text == "" || playtimeEntry.Text == "" {
			consoleLogs.SetText(consoleLogs.Text + "Arguments missing!\n")
			return
		}

		newVersion := tools.Version{
			Path:    playtimeEntry.Text,
			Picture: "./assets/unknown_version.png",
			Title:   titleEntry.Text,
		}
		versions = append(versions, newVersion)

		img := canvas.NewImageFromFile(newVersion.Picture)
		img.SetMinSize(fyne.NewSize(60, 60))
		img.FillMode = canvas.ImageFillContain
		titleLabel := widget.NewLabel(newVersion.Title)
		generatedText := widget.NewLabel("Playtime: Calculating")

		go func(path string, label *widget.Label, console *widget.Entry) {
			playtime := tools.Get_time_for_directory(path, consoleLogs)
			update_overall(playtime, overallPlaytime)
			label.SetText(fmt.Sprintf("Playtime: %d hours %d minutes %d seconds", playtime/3600, (playtime%3600)/60, playtime%60))
		}(playtimeEntry.Text, generatedText, consoleLogs)

		playtimeEntry.SetText("")
		titleEntry.SetText("")

		entry := container.NewHBox(img, container.NewVBox(titleLabel, generatedText))
		playtimeContent.Add(entry)

		playtimeContent.Refresh()
	})
	addContainer := container.New(layout.NewBorderLayout(nil, nil, nil, addButton), addButton, container.NewVBox(titleEntry, playtimeEntry))

	playtimeTab := container.NewBorder(nil, addContainer, nil, nil, playtimeScroll)

	tabs := container.NewAppTabs(
		container.NewTabItem("Playtime", playtimeTab),
		container.NewTabItem("Console", consoleTab),
	)

	window.SetContent(tabs)
	window.ShowAndRun()
}
