package main

import (
	"log"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

func main() {

	// gets the Searcher for fetching
	// data manganato
	screen.searcher = nato.NewSearcher()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// focused view will have different color due to
	// highlight is set to true (has green border
	// while others will be yellow)
	g.Highlight = true
	g.BgColor = gocui.ColorBlack
	// sets border color to be yellow
	g.FgColor = gocui.ColorYellow
	g.Cursor = true

	// layout function gets executed on every
	// iteration of gocui's GUI mainloop
	g.SetManagerFunc(layout)

	// terminal's width and height
	maxX, maxY := g.Size()

	// gets SearchBar view from GetSearchBar (defined in ./views directory)
	if screen.sb, err = views.GetSearchBar(maxX, maxY, g); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}

		// sets cursor to SearchBar when program first started
		if _, err = g.SetCurrentView(views.SearchBarName); err != nil {
			log.Panicln(err)
		}
	}

	// gets SearchList view
	if screen.sl, err = views.GetSearchList(maxX, maxY, g); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}

		// renders a list of latest updated mangas from manganato
		// when the program first started
		if err = getInitialScreen(); err != nil {
			log.Panicln(err)
		}
	}

	// gets MangaDetails view, which will be empty at first
	if screen.md, err = views.GetMangaDetails(maxX, maxY, g); err != nil && err != gocui.ErrUnknownView {
		log.Panicln(err)
	}

	// gets ChapterList view, also empty at first
	if screen.cl, err = views.GetChapterList(maxX, maxY, g); err != nil && err != gocui.ErrUnknownView {
		log.Panicln(err)
	}

	// all keybindings are defined in keybindings.go
	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	// starts the gocui's main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
