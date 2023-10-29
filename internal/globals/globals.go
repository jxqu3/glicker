package globals

var Started bool = false
var Clicking bool = false
var Cps float64 = 1
var WaitForKey bool = false
var ToggleKey uint16 = 66
var RandomVariation float64 = 0
var MouseButtons = map[string]bool{
	"left":       true,
	"right":      false,
	"center":     false,
	"wheelDown":  false,
	"wheelUp":    false,
	"wheelLeft":  false,
	"wheelRight": false,
}
