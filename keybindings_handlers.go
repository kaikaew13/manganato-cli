package main

import "github.com/jroimartin/gocui"

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
