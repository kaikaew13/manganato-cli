package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/kaikaew13/manganato-cli/views"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func switchView(g *gocui.Gui, v *gocui.View) error {
	for i, name := range viewNames {
		if name == v.Name() {
			g.SetCurrentView(viewNames[(i+1)%len(viewNames)])
			break
		}
	}
	return nil
}

func enterCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()

	mdView, err := g.View(views.MangaDetailsName)
	if err != nil {
		return err
	}
	fmt.Fprint(mdView, s)

	x, y := v.Origin()
	if err = v.SetCursor(x, y); err != nil {
		return err
	}

	v.Clear()

	return nil
}
