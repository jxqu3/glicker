package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/checkm4ted/Glicker/internal/utils"
	"golang.org/x/sys/windows"
)

func main() {
	// Call timeBeginPeriod to fix timer ;)
	windows.TimeBeginPeriod(1)

	a := app.New()

	w := *utils.Setup(&a)

	utils.Gk = utils.Glicker{
		App:    &a,
		Window: &w,
	}

	utils.Layout(&w)

	(w).ShowAndRun()
}
