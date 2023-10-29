package layout

import (
	"fmt"
	"math/rand"
	sc "strconv"
	"time"

	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	g "github.com/checkm4ted/Glicker/internal/globals"
	"github.com/checkm4ted/Glicker/internal/utils"
	"github.com/go-vgo/robotgo"
)

var randomVariation float64 = 0
var button string = "left"
var slvalue = 0

func Layout(w *f.Window) {

	var setKeyBtn *widget.Button
	setKeyBtn = widget.NewButton("Toggle Key: "+utils.KeycodeToName(g.ToggleKey), func() {
		g.WaitForKey = true
		setKeyBtn.SetText("Press A key")
	})

	var startBtn *widget.Button
	startBtn = widget.NewButton("Start", func() {
		g.Started = !g.Started
		if g.Started {
			startBtn.SetText(fmt.Sprintf("Stop (Toggle with %s)", utils.KeycodeToName(g.ToggleKey)))
			g.Clicking = false
			go func() {
				for g.Started {
					//wait until clicking to save CPU.
					for !g.Clicking {
						time.Sleep(100 * time.Millisecond)
					}
					if g.Clicking {
						robotgo.MouseDown(button)
						time.Sleep(1 * time.Millisecond)
						robotgo.MouseUp(button)
						// Sleep for CPS + random variation
						if g.Cps < 0.1 {
							time.Sleep(50 * time.Millisecond)
							continue
						}
						time.Sleep(time.Duration(1000/g.Cps+(rand.Float64()*randomVariation)) * time.Millisecond)
					}
				}
			}()
		} else {

			startBtn.SetText("Start")
		}
	})

	cpslabel := widget.NewLabel("CPS: " + sc.FormatFloat(float64(g.Cps), 'f', 1, 32))

	// CPS Slider
	slider := widget.NewSlider(0, 30)
	slider.Step = 0.1
	slider.SetValue(float64(slvalue))
	slider.OnChanged = func(value float64) {
		g.Cps = utils.LogSlider(value)
		cpslabel.SetText("CPS: " + sc.FormatFloat(g.Cps, 'f', 1, 32))
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

	utils.StartKeyboard(&setKeyBtn, &startBtn)

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
