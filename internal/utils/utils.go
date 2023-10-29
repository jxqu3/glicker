package utils

import (
	"fmt"
	"math"

	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	g "github.com/checkm4ted/Glicker/internal/globals"
	hook "github.com/robotn/gohook"
)

func Setup(a *f.App) *f.Window {
	w := (*a).NewWindow("Glicker")
	w.SetFixedSize(true)
	w.Resize(f.Size{Width: 400, Height: 600})
	w.CenterOnScreen()
	return &w
}

func KeycodeToName(k uint16) string {
	for i, v := range hook.Keycode {
		if v == k {
			return i
		}
	}
	return "Unknown"
}

func LogSlider(position float64) float64 {
	return math.Pow(10, position/10)
}

func StartKeyboard(setKeyBtn **widget.Button, startBtn **widget.Button) {
	go func() {
		eventHook := hook.Start()
		for e := range eventHook {
			if e.Kind == hook.KeyUp {
				if e.Keycode == g.ToggleKey {
					g.Clicking = !g.Clicking
				}
				if g.WaitForKey {
					g.ToggleKey = e.Keycode
					g.WaitForKey = false
					(*setKeyBtn).SetText("Toggle Key: " + KeycodeToName(g.ToggleKey))
					(*startBtn).SetText(fmt.Sprintf("Stop (Toggle with %s)", KeycodeToName(g.ToggleKey)))
				}
			}
		}
	}()
}
