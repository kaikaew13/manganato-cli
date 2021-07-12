package main

import (
	"log"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

func main() {
	screen.searcher = nato.NewSearcher()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorYellow
	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	maxX, maxY := g.Size()

	if screen.sb, err = views.GetSearchBar(maxX, maxY, g); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
		if _, err = g.SetCurrentView(views.SearchBarName); err != nil {
			log.Panicln(err)
		}
	}

	if screen.sl, err = views.GetSearchList(maxX, maxY, g); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
		if err = getInitialScreen(); err != nil {
			log.Panicln(err)
		}
	}

	if screen.md, err = views.GetMangaDetails(maxX, maxY, g); err != nil && err != gocui.ErrUnknownView {
		log.Panicln(err)
	}

	if screen.cl, err = views.GetChapterList(maxX, maxY, g); err != nil && err != gocui.ErrUnknownView {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
