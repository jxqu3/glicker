package layout

import (
	"fmt"
	sc "strconv"

	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	g "github.com/checkm4ted/Glicker/internal/globals"
	"github.com/checkm4ted/Glicker/internal/utils"
)

var mouseButtons = []string{
	"left", "right", "center", "wheelDown", "wheelUp", "wheelLeft", "wheelRight",
}

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
			utils.StartClicking()
		} else {
			startBtn.SetText("Start")
		}
	})

	cpslabel := widget.NewLabel("CPS: " + sc.FormatFloat(float64(g.Cps), 'f', 1, 32))

	// CPS Slider
	slider := widget.NewSlider(0, 30)
	slider.Step = 0.1
	slider.SetValue(0)
	slider.OnChanged = func(value float64) {
		g.Cps = utils.LogSlider(value)
		cpslabel.SetText("CPS: " + sc.FormatFloat(g.Cps, 'f', 1, 32))
	}

	rndlabel := widget.NewLabel("Random Variation: " + sc.FormatFloat(float64(g.RandomVariation), 'f', 1, 32) + "ms")

	// Random Variation Slider
	rSlider := widget.NewSlider(0, 1000)

	rSlider.OnChanged = func(value float64) {
		g.RandomVariation = value
		rndlabel.SetText("Random Variation: " + sc.FormatFloat(value, 'f', 1, 32) + "ms")
	}

	rSlider.Step = 0.1

	dropdown := widget.NewSelect(mouseButtons, func(s string) {
		g.MouseButton = s
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
