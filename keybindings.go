package main

import (
	"github.com/jroimartin/gocui"
)

func keybindings(g *gocui.Gui) error {

	// g.SetKeybinding with the first argument as ""
	// means that keybind will apply to every views

	// quits the program when ctrl + c and does not panic
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// switch between views in clockwise direction by tab
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switchView); err != nil {
		return err
	}

	// use back tick (`) to switch view anti-clockwisely
	if err := g.SetKeybinding("", '`', gocui.ModNone, reverseSwitchView); err != nil {
		return err
	}

	if err := g.SetKeybinding(screen.sb.Name, gocui.KeyEnter, gocui.ModNone, enterCommand); err != nil {
		return err
	}

	if err := g.SetKeybinding(screen.sb.Name, gocui.KeyArrowUp, gocui.ModNone, getPrevCommand); err != nil {
		return err
	}

	if err := g.SetKeybinding(screen.sb.Name, gocui.KeyArrowDown, gocui.ModNone, getNextCommand); err != nil {
		return err
	}

	if err := g.SetKeybinding(screen.sl.Name, gocui.KeyEnter, gocui.ModNone, pickManga); err != nil {
		return err
	}

	if err := g.SetKeybinding(screen.cl.Name, gocui.KeyEnter, gocui.ModNone, pickChapter); err != nil {
		return err
	}

	return nil
}
