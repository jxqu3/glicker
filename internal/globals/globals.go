package globals

type MouseButton struct {
	Name  string
	State bool
}

var Started bool = false
var Clicking bool = false
var Cps float64 = 1
var WaitForKey bool = false
var ToggleKey uint16 = 66
var RandomVariation float64 = 0
var MouseButtons = []MouseButton{
	{"left", true},
	{"right", false},
	{"center", false},
	{"wheelDown", false},
	{"wheelUp", false},
	{"wheelLeft", false},
	{"wheelRight", false},
}
