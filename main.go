package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/kaikaew13/manganato-cli/views"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite
	g.Cursor = true

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	maxX, maxY := g.Size()

	views.GetSearchBar(maxX, maxY, g)
	views.GetSearchList(maxX, maxY, g)
	views.GetMangaDetails(maxX, maxY, g)
	views.GetChapterList(maxX, maxY, g)

	g.SetCurrentView(views.SearchBarName)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
