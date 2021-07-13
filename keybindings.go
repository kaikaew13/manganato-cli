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

	if err := g.SetKeybinding("", '`', gocui.ModNone, reverseSwitchView); err != nil {
		return err
	}

	if err := g.SetKeybinding(views.SearchBarName, gocui.KeyEnter, gocui.ModNone, enterCommand); err != nil {
		return err
	}

	if err := g.SetKeybinding(views.SearchBarName, gocui.KeyArrowUp, gocui.ModNone, getPrevCommand); err != nil {
		return err
	}

	if err := g.SetKeybinding(views.SearchBarName, gocui.KeyArrowDown, gocui.ModNone, getNextCommand); err != nil {
		return err
	}

	if err := g.SetKeybinding(views.SearchListName, gocui.KeyEnter, gocui.ModNone, pickManga); err != nil {
		return err
	}

	if err := g.SetKeybinding(views.ChapterListName, gocui.KeyEnter, gocui.ModNone, pickChapter); err != nil {
		return err
	}

	return nil
}
