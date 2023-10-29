package utils

import (
	"fmt"
	"math"
	"math/rand"
	sc "strconv"
	"time"

	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var Started bool = false
var Clicking bool = false
var Cps float64 = 1
var randomVariation float64 = 10
var toggleKey uint16 = 66
var button string = "left"
var waitForKey bool = false
var slvalue = 0

func keycodeToName(k uint16) string {
	for i, v := range hook.Keycode {
		if v == k {
			return i
		}
	}
	return "Unknown"
}

func logslider(position float64) float64 {
	return math.Pow(10, position/10)
}

func Layout(w *f.Window) {

	var setKeyBtn *widget.Button
	setKeyBtn = widget.NewButton("Toggle Key: "+keycodeToName(toggleKey), func() {
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
					//wait until clicking to save CPU.
					for !Clicking {
						time.Sleep(100 * time.Millisecond)
					}
					if Clicking {
						robotgo.MouseDown(button)
						time.Sleep(1 * time.Millisecond)
						robotgo.MouseUp(button)
						// Sleep for CPS + random variation
						if Cps < 0.1 {
							time.Sleep(50 * time.Millisecond)
							continue
						}
						time.Sleep(time.Duration(1000/Cps+(rand.Float64()*randomVariation)) * time.Millisecond)
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
					setKeyBtn.SetText("Toggle Key: " + keycodeToName(toggleKey))
					startBtn.SetText(fmt.Sprintf("Stop (Toggle with %s)", keycodeToName(toggleKey)))
				}
			}
		}
	}()

	cpslabel := widget.NewLabel("CPS: " + sc.FormatFloat(float64(Cps), 'f', 1, 32))

	// CPS Slider
	slider := widget.NewSlider(0, 30)
	slider.Step = 0.1
	slider.SetValue(float64(slvalue))
	slider.OnChanged = func(value float64) {
		Cps = logslider(value)
		cpslabel.SetText("CPS: " + sc.FormatFloat(Cps, 'f', 1, 32))
	}

	rndlabel := widget.NewLabel("Random Variation: " + sc.FormatFloat(float64(randomVariation), 'f', 1, 32) + "ms")

	// Random Variation Slider
	rSlider := widget.NewSlider(0, 1000)

	rSlider.OnChanged = func(value float64) {
		randomVariation = value
		rndlabel.SetText("Random Variation: " + sc.FormatFloat(value, 'f', 1, 32) + "ms")
	}

	rSlider.Step = 0.1

	dropdown := widget.NewSelect([]string{"left", "right", "center", "wheelDown", "wheelUp", "wheelLeft", "wheelRight"}, func(s string) {
		button = s
	})

	dropdown.SetSelected("left")

	(*w).SetContent(
		container.New(layout.NewVBoxLayout(),
			container.New(layout.NewVBoxLayout(), widget.NewLabelWithStyle("Glicker", f.TextAlignCenter, f.TextStyle{Bold: true}),
				cpslabel,
				slider,
				rndlabel,
				rSlider,
				widget.NewLabel("Mouse Button:"),
				dropdown,
			),
			setKeyBtn,
			startBtn,
		))

}
