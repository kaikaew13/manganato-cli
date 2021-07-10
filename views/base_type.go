package views

import "github.com/jroimartin/gocui"

type BaseType struct {
	View *gocui.View
	Name string
	X0   int
	X1   int
	Y0   int
	Y1   int
}
