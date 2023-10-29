package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	g "github.com/checkm4ted/Glicker/internal/globals"
	"github.com/go-vgo/robotgo"
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

func StartClicking() {
	go func() {
		for g.Started {
			//wait until clicking to save CPU.
			for !g.Clicking {
				if !g.Started {
					break
				}
				time.Sleep(100 * time.Millisecond)
			}
			if g.Clicking {
				for i, v := range g.MouseButtons {
					if v {
						robotgo.MouseDown(i)
						time.Sleep(time.Duration(rand.Float32()*0.8+0.2) * time.Millisecond)
						robotgo.MouseUp(i)
					}
				}

				// Sleep for CPS + random variation
				if g.RandomVariation > 0 {
					time.Sleep(time.Duration(1000/g.Cps+(rand.Float64()*g.RandomVariation)) * time.Millisecond)
				} else {
					time.Sleep(time.Duration(1000/g.Cps) * time.Millisecond)
				}
			}
		}
	}()
}
