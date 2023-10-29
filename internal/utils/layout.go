package utils

import (
	"fmt"
	"math/rand"
	sc "strconv"
	"time"

	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"

	_ "github.com/defia/trf"
)

var Started bool = false
var Clicking bool = false
var Cps float32 = 10
var randomVariation float32 = 10
var toggleKey uint16 = 66
var button string = "left"
var waitForKey bool = false

func keycodeToName(k uint16) string {
	for i, v := range hook.Keycode {
		if v == k {
			return i
		}
	}
	return "Unknown"
}

func Layout(w *f.Window) {

	var setKeyBtn *widget.Button
	setKeyBtn = widget.NewButton("Key: "+keycodeToName(toggleKey), func() {
		waitForKey = true
		setKeyBtn.SetText("Press A key")
	})

	var startBtn *widget.Button
	startBtn = widget.NewButton("Start", func() {
		Started = !Started
		if Started {
			startBtn.SetText(fmt.Sprintf("Stop (Toggle with %s)", keycodeToName(toggleKey)))
			Clicking = false
			go func() {
				for Started {
					for !Clicking {
						time.Sleep(50 * time.Millisecond)
					}
					if Clicking {
						robotgo.MouseDown(button)
						time.Sleep(1 * time.Millisecond)
						robotgo.MouseUp(button)
						// Sleep for CPS + random variation
						time.Sleep(time.Duration(1000/Cps+(rand.Float32()*randomVariation)) * time.Millisecond)
					}
				}
			}()
		} else {
			startBtn.SetText("Start")
		}
	})

	go func() {
		eventHook := hook.Start()
		for e := range eventHook {
			if e.Kind == hook.KeyUp {
				if e.Keycode == toggleKey {
					Clicking = !Clicking
				}
				if waitForKey {
					toggleKey = e.Keycode
					waitForKey = false
					setKeyBtn.SetText("Key: " + keycodeToName(toggleKey))
					startBtn.SetText(fmt.Sprintf("Stop (Toggle with %s)", keycodeToName(toggleKey)))
				}
			}
		}
	}()

	cpslabel := widget.NewLabel("CPS: " + sc.FormatFloat(float64(Cps), 'f', 1, 32))

	// CPS Slider
	slider := widget.NewSlider(0.1, 1000)
	slider.SetValue(float64(Cps))
	slider.OnChanged = func(value float64) {
		Cps = float32(slider.Value)
		cpslabel.SetText("CPS: " + sc.FormatFloat(value, 'f', 1, 32))
	}

	rndlabel := widget.NewLabel("Random Variation: " + sc.FormatFloat(float64(randomVariation), 'f', 1, 32) + "ms")

	// Random Variation Slider
	rSlider := widget.NewSlider(0, 1000)

	rSlider.OnChanged = func(value float64) {
		randomVariation = float32(slider.Value)
		rndlabel.SetText("Random Variation: " + sc.FormatFloat(value, 'f', 1, 32) + "ms")
	}

	(*w).SetContent(
		container.New(layout.NewVBoxLayout(),
			widget.NewLabelWithStyle("Glicker", f.TextAlignCenter, f.TextStyle{Bold: true}),
			cpslabel,
			slider,
			rndlabel,
			rSlider,
			widget.NewSelect([]string{"left", "right", "center", "wheelDown", "wheelUp", "wheelLeft", "wheelRight"}, func(s string) {
				button = s
			}),
			startBtn,
			setKeyBtn,
		))

}
