package utils

import (
	f "fyne.io/fyne/v2"
)

func Setup(a *f.App) *f.Window {
	w := (*a).NewWindow("Glicker")
	w.SetFixedSize(true)
	w.Resize(f.Size{Width: 400, Height: 600})
	w.CenterOnScreen()
	return &w
}
