package utils

import (
	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Started bool = false

func Layout(w *f.Window) {
	var startBtn *widget.Button
	startBtn = widget.NewButton("Start", func() {
		Started = !Started
		if Started {
			startBtn.SetText("Stop")
		} else {
			startBtn.SetText("Start")
		}
	})

	(*w).SetContent(
		container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Glicker"),
			startBtn,
		))
}
