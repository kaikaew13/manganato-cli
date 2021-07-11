package main

import (
	"strings"

	"github.com/jroimartin/gocui"
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

	screen.sb.SaveCommand(s)

	x, y := v.Origin()
	if err := v.SetCursor(x, y); err != nil {
		return err
	}

	v.Clear()

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
	_, y := v.Cursor()

	s, _ := v.Line(y)
	s = strings.TrimSpace(s)

	if len(s) > 2 && s[:3] == "[x]" {
		mgName := s[4:]
		mgId := screen.sl.NameToIDMap[mgName]

		mg, err := screen.searcher.PickManga(mgId)
		if err != nil {
			return err
		}

		screen.md.Manga = *mg
		s = screen.md.FormatManga()
		screen.md.View.Clear()
		screen.md.View.Write([]byte(s))

		screen.cl.Chapters = mg.Chapters
		s = screen.cl.FormatChapters()
		screen.cl.View.Clear()
		screen.cl.View.Write([]byte(s))
	}

	return nil
}
