package main

import (
	"github.com/jroimartin/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func switchView(g *gocui.Gui, v *gocui.View) error {
	resetCursor(v)

	for i, name := range viewNames {
		if name == v.Name() {
			if i == len(viewNames)-1 {
				g.SetCurrentView(viewNames[0])
				break
			}

			g.SetCurrentView(viewNames[i+1])
			break
		}
	}

	return nil
}

func reverseSwitchView(g *gocui.Gui, v *gocui.View) error {
	resetCursor(v)

	for i, name := range viewNames {
		if name == v.Name() {
			if i == 0 {
				g.SetCurrentView(viewNames[len(viewNames)-1])
				break
			}

			g.SetCurrentView(viewNames[i-1])
			break
		}
	}

	return nil
}

func enterCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()

	x, y := v.Origin()
	if err := v.SetCursor(x, y); err != nil {
		return err
	}

	v.Clear()

	valid, cmd, args := validateCommand(s)
	if valid {
		screen.sb.SaveCommand(s)

		if err := runCommand(cmd, args); err != nil {
			return err
		}
	}

	return nil
}

func getPrevCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()
	s = screen.sb.GetPrevCommand(s)

	v.Clear()
	v.Write([]byte(s))

	return nil
}

func getNextCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()
	s = screen.sb.GetNextCommand(s)

	v.Clear()
	v.Write([]byte(s))

	return nil
}

func pickManga(g *gocui.Gui, v *gocui.View) error {
	s := trimLine(v)

	err := getMangaScreen(s)

	return err
}

func pickChapter(g *gocui.Gui, v *gocui.View) error {
	s := trimLine(v)

	err := downloadChapter(s)

	return err
}
