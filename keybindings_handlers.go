package main

import "github.com/jroimartin/gocui"

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
