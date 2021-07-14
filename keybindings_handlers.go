package main

import (
	"github.com/jroimartin/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func switchView(g *gocui.Gui, v *gocui.View) error {
	resetCursor(v)

	vNames := getViewNames(g)
	for i, name := range vNames {
		if name == v.Name() {
			if i == len(vNames)-1 {
				g.SetCurrentView(vNames[0])
				break
			}

			g.SetCurrentView(vNames[i+1])
			break
		}
	}

	return nil
}

func reverseSwitchView(g *gocui.Gui, v *gocui.View) error {
	resetCursor(v)

	vNames := getViewNames(g)
	for i, name := range vNames {
		if name == v.Name() {
			if i == 0 {
				g.SetCurrentView(vNames[len(vNames)-1])
				break
			}

			g.SetCurrentView(vNames[i-1])
			break
		}
	}

	return nil
}

func enterCommand(g *gocui.Gui, v *gocui.View) error {
	if err := openModal(g); err != nil {
		return err
	}

	go processCommand(g, v)

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
	s := trimViewLine(v)

	err := getMangaScreen(s)

	return err
}

func pickChapter(g *gocui.Gui, v *gocui.View) error {
	if err := openModal(g); err != nil {
		return err
	}

	go func() {
		g.Update(func(g *gocui.Gui) error {
			s := trimViewLine(v)

			if err := downloadChapter(s); err != nil {
				return nil
			}

			err := closeModal(g)
			return err
		})
	}()

	return nil
}
