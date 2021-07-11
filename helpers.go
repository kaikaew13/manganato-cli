package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/kaikaew13/manganato-cli/views"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	var err error

	if screen.sb, err = views.GetSearchBar(maxX, maxY, g); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err = g.SetCurrentView(views.SearchBarName); err != nil {
			return err
		}
	}

	if screen.sl, err = views.GetSearchList(maxX, maxY, g); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err = getInitialScreen(); err != nil {
			return err
		}
	}

	if screen.md, err = views.GetMangaDetails(maxX, maxY, g); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	if screen.cl, err = views.GetChapterList(maxX, maxY, g); err != nil && err != gocui.ErrUnknownView {
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
