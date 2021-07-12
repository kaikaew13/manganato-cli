package main

import (
	"errors"
	"strings"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

const (
	searchCommand string = "search"
	// searchByAuthorCommand  = searchCommand + "-author"
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

	screen.sl.View.Write([]byte(screen.sl.FormatMangas()))

	return nil
}

func getMangaScreen(s string) error {
	if len(s) >= len(views.Selector) && strings.HasPrefix(s, views.Selector) {
		mgName, mgId := getMangaNameAndID(s, len(views.Selector))

		mg, err := screen.searcher.PickManga(mgId)
		if err != nil {
			return err
		}

		screen.md.Manga = *mg
		s = screen.md.FormatManga()
		screen.md.View.Clear()
		screen.md.View.Write([]byte(s))

		screen.cl.MangaName = mgName
		screen.cl.MangaID = mgId
		screen.cl.Chapters = mg.Chapters
		s = screen.cl.FormatChapters()
		screen.cl.View.Clear()
		screen.cl.View.Write([]byte(s))

		return nil
	}

	return errors.New("not a selectable line")
}

func getMangaNameAndID(s string, prefLen int) (mgName, mgId string) {
	mgName = s[prefLen+1:]
	mgId = screen.sl.NameToIDMap[mgName]
	return
}

func validateCommand(s string) (valid bool, cmd, args string) {
	if strings.HasPrefix(s, searchCommand) {
		valid = true
		cmd = searchCommand
		args = s[len(searchCommand)+1:]

		return
	}

	return
}

func runCommand(cmd, args string) error {
	switch cmd {
	case searchCommand:
		mgs, err := screen.searcher.SearchManga(args)
		if err != nil && err != nato.ErrPageNotFound {
			return err
		}

		screen.sl.View.Clear()
		if err == nato.ErrPageNotFound {
			screen.sl.View.Write([]byte(nato.ErrPageNotFound.Error()))
			return nil
		}

		screen.sl.Mangas = *mgs
		s := screen.sl.FormatMangas()
		screen.sl.View.Write([]byte(s))
	}

	return nil
}
