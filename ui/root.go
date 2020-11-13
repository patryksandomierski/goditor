package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// DrawMainWindow show and runs new main window
func DrawMainWindow(appName string) {
	app := app.New()
	window := app.NewWindow(appName)
	window.Resize(fyne.NewSize(640, 480))

	window.SetMainMenu(
		fyne.NewMainMenu(
			fyne.NewMenu("File",
				fyne.NewMenuItem("Menu item 1", func() {}),
				fyne.NewMenuItemSeparator(),
				fyne.NewMenuItem("Menu item 2", func() {}),
			),
			fyne.NewMenu("Search",
				fyne.NewMenuItem("Find", func() {
					position := fyne.NewPos(window.Canvas().Size().Width/2,
						window.Canvas().Size().Height/2)
					popup := widget.NewPopUpAtPosition(getSearchDialogContent(),
						window.Canvas(),
						position)
					popup.Show()
				}),
			),
		),
	)

	

	text1 := canvas.NewText("top", Red)
	text1.Alignment = fyne.TextAlignCenter
	text2 := canvas.NewText("bottom", Red)
	text2.Alignment = fyne.TextAlignCenter
	text3 := canvas.NewText("left", Red)
	text4 := canvas.NewText("right", Red)
	text5 := canvas.NewRectangle(color.Black)


	

	layout := layout.NewBorderLayout(text1, text2, text3, text4)

	cont := fyne.NewContainerWithLayout(layout, text1, text2, text3, text4, text5)

	window.SetContent(cont)

	window.ShowAndRun()
}

func getSearchDialogContent() *widget.Form {
	queryEntry := widget.NewEntry()
	// queryEntry.Resize(fyne.NewSize(500, 500))
	queryEntry.KeyDown(&fyne.KeyEvent{
		Name: fyne.KeyDown,
	})
	queryEntry.SetPlaceHolder("type here")
	queryFormItem := &widget.FormItem{
		Text:   "Search for:",
		Widget: queryEntry,
	}
	form := widget.NewForm(queryFormItem)
	form.SubmitText = "search"
	form.CancelText = "cancel"
	form.OnSubmit = func() { fmt.Println("onSubmit action") }
	form.OnCancel = func() { fmt.Println("onCancel") }
	return form
}
