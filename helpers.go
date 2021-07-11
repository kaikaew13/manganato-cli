package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/kaikaew13/manganato-cli/views"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	var err error

	screen.sb, err = views.GetSearchBar(maxX, maxY, g)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetCurrentView(views.SearchBarName)
	}

	screen.sl, err = views.GetSearchList(maxX, maxY, g)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err = getInitialScreen(); err != nil {
			return err
		}
	}

	screen.md, err = views.GetMangaDetails(maxX, maxY, g)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	screen.cl, err = views.GetChapterList(maxX, maxY, g)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	// screen.sb = sb
	// screen.sl = sl
	// screen.md = md
	// screen.cl = cl

	return nil
}

func getInitialScreen() error {
	mgs, err := screen.searcher.SearchLatestUpdatedManga()
	if err != nil {
		return err
	}
	screen.sl.Mangas = *mgs
	fmt.Fprint(screen.sl.View, screen.sl.FormatMangas())
	return nil
}
