package main

import (
	"log"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

var viewNames = []string{
	views.SearchBarName,
	views.SearchListName,
	views.MangaDetailsName,
	views.ChapterListName,
}

func main() {
	nato.NewSearcher()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite
	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
