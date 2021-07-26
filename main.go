package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

var wg sync.WaitGroup

func runCui() {
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
	// sets border color to be yellow when unfocused
	g.FgColor = gocui.ColorWhite
	g.SelFgColor = gocui.ColorGreen
	g.SelBgColor = gocui.ColorBlack
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
		if _, err = g.SetCurrentView(screen.sb.Name); err != nil {
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

func main() {
	mangaQuery := flag.String(searchCommand, "", "search manga based on title pattern")

	mangaId := flag.String("manga-id", "", "download/show resources of selected manga id")

	downloadSelection := flag.String("download", "", "download manga chapters ('*' to download all, chapters numbers to download specific chapters, comma-separated lists or dash-ranges to perform batch download)")

	flag.Parse()

	// gets the Searcher for fetching
	// data manganato
	screen.searcher = nato.NewSearcher()

	var err error = nil
	if *mangaQuery != "" {
		err = printList(searchManga(*mangaQuery))
		fmt.Printf("\nTo list manga chapters run '%v -manga-id MANGAID', where MANGAID is the value between square braces in the list above\n", os.Args[0])
	} else if *mangaId != "" {
		if *downloadSelection != "" {
			err = downloadChapters(*mangaId, *downloadSelection)
		} else {
			err = printList(listChapters(*mangaId))
			fmt.Printf("\nTo download chapters run '%v -manga-id %v -download SELECTION', where SELECTION is a list or single value from those between square braces in the list above\n", os.Args[0], *mangaId)
		}
	} else if *downloadSelection != "" {
		flag.PrintDefaults()
		err = errors.New("no manga-id specified")
	} else {
		runCui()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
