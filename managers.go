package main

import (
	"github.com/jroimartin/gocui"
	"github.com/kaikaew13/manganato-cli/views"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	_, err := views.GetSearchBar(maxX, maxY, g)
	if err != nil {
		return err
	}

	_, err = views.GetSearchList(maxX, maxY, g)
	if err != nil {
		return err
	}

	_, err = views.GetMangaDetails(maxX, maxY, g)
	if err != nil {
		return err
	}

	_, err = views.GetChapterList(maxX, maxY, g)
	if err != nil {
		return err
	}

	return nil
}
