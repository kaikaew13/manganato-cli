package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const MangaDetailsName string = "MangaDetails"

type MangaDetails struct {
	View        *gocui.View
	Manga       nato.Manga
	NameToIDMap map[string]string
}

func GetMangaDetails(maxX, maxY int, g *gocui.Gui) (*MangaDetails, error) {
	md := MangaDetails{}
	x0, y0, x1, y1 := md.GetCoords(maxX, maxY)

	mdView, err := g.SetView(MangaDetailsName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	mdView.Title = MangaDetailsName
	mdView.SelFgColor = gocui.ColorGreen
	mdView.BgColor = gocui.ColorBlack
	mdView.FgColor = gocui.ColorWhite
	mdView.Editable = true
	mdView.Editor = selectingEditor
	mdView.Wrap = true

	md.View = mdView
	md.Manga = nato.Manga{}
	md.NameToIDMap = make(map[string]string)

	return &md, err
}

func (md *MangaDetails) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return maxX / 2, 1, maxX - 1, ((maxY - SearchBarHeight - 1) / 2) - 1
}

func (md *MangaDetails) FormatManga() string {
	s := "\n\n"

	s += fmt.Sprintf("		TITLE: %s\n\n", md.Manga.Name)
	s += fmt.Sprintf("		ALT_NAME: %s\n\n", md.Manga.Alternatives)
	s += fmt.Sprintf("		STATUS: %s\n\n", md.Manga.Status)

	var genres string
	for _, v := range md.Manga.Genres {
		genres += v.GenreName + "\t"
	}

	s += fmt.Sprintf("		GENRES: %s\n\n", genres)
	s += fmt.Sprintf("		AUTHOR: %s\n\n", md.Manga.Author.Name)
	s += fmt.Sprintf("		UPDATED: %s\n\n", md.Manga.Updated)
	s += fmt.Sprintf("		VIEWS: %s\n\n", md.Manga.Views)
	s += fmt.Sprintf("		RATING: %s\n\n", md.Manga.Rating)
	s += fmt.Sprintf("		DESCRIPTION: %s\n\n", md.Manga.Description)

	return s
}
