package main

import (
	"runtime"

	"fyne.io/fyne/v2/app"
	"github.com/checkm4ted/Glicker/internal/layout"
	"github.com/checkm4ted/Glicker/internal/utils"
	"golang.org/x/sys/windows"
)

func main() {
	// Call timeBeginPeriod to fix timer ;)
	if runtime.GOOS == "windows" {
		windows.TimeBeginPeriod(1)
	}

	a := app.New()

	w := *utils.Setup(&a)

	layout.Layout(&w)

	(w).ShowAndRun()

}
