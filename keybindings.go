package main

import (
	"github.com/jroimartin/gocui"
	"github.com/kaikaew13/manganato-cli/views"
)

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switchView); err != nil {
		return err
	}

	if err := g.SetKeybinding(views.SearchBarName, gocui.KeyEnter, gocui.ModNone, enterCommand); err != nil {
		return err
	}

	return nil
}
