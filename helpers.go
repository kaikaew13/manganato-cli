package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/kaikaew13/manganato-cli/views"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	x0, y0, x1, y1 := screen.sb.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.SearchBarName, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.sl.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.SearchListName, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.md.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.MangaDetailsName, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.cl.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.ChapterListName, x0, y0, x1, y1); err != nil {
		return err
	}

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
