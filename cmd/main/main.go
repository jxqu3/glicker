package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/checkm4ted/Glicker/internal/utils"
)

func main() {

	a := app.New()

	w := *utils.Setup(&a)

	utils.Gk = utils.Glicker{
		App:    &a,
		Window: &w,
	}

	utils.Layout(&w)

	(w).ShowAndRun()
}
