package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const mangaDetailsName string = "MangaDetails"

type MangaDetails struct {
	View        *gocui.View
	Name        string
	Manga       nato.Manga
	NameToIDMap map[string]string
	OriginX     int
	OriginY     int
}

func GetMangaDetails(maxX, maxY int, g *gocui.Gui) (*MangaDetails, error) {
	md := MangaDetails{}
	x0, y0, x1, y1 := md.GetCoords(maxX, maxY)

	mdView, err := g.SetView(mangaDetailsName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	mdView.Title = mangaDetailsName
	mdView.SelFgColor = gocui.ColorBlack
	mdView.SelBgColor = gocui.ColorGreen
	mdView.BgColor = gocui.ColorBlack
	mdView.FgColor = gocui.ColorWhite
	mdView.Highlight = true
	mdView.Editable = true
	mdView.Wrap = true
	mdView.Editor = readOnlyEditor

	md.View = mdView
	md.Name = mangaDetailsName
	md.Manga = nato.Manga{}
	md.NameToIDMap = make(map[string]string)
	md.OriginX, md.OriginY = mdView.Origin()

	return &md, err
}

func (md *MangaDetails) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return maxX / 2, 1, maxX - 1, ((maxY - SearchBarHeight - 1) / 2) - 1
}

func (md *MangaDetails) FormatManga() string {
	s := "\n\n"

	s += fmt.Sprintf("\t\tTITLE: %s\n\n", md.Manga.Name)
	s += fmt.Sprintf("\t\tALT_NAME: %s\n\n", md.Manga.Alternatives)
	s += fmt.Sprintf("\t\tSTATUS: %s\n\n", md.Manga.Status)

	var genres string
	for _, v := range md.Manga.Genres {
		genres += v.GenreName + "\t"
		md.NameToIDMap[v.GenreName] = v.ID
		md.NameToIDMap[strings.ToLower(v.GenreName)] = v.ID
		md.NameToIDMap[strings.ToUpper(v.GenreName)] = v.ID
	}

	s += fmt.Sprintf("\t\tGENRES: %s\n\n", genres)
	s += fmt.Sprintf("\t\tAUTHOR: %s\n\n", md.Manga.Author.Name)
	s += fmt.Sprintf("\t\tUPDATED: %s\n\n", md.Manga.Updated)
	s += fmt.Sprintf("\t\tVIEWS: %s\n\n", md.Manga.Views)
	s += fmt.Sprintf("\t\tRATING: %s\n\n", md.Manga.Rating)
	s += fmt.Sprintf("\t\tDESCRIPTION: %s\n\n", md.Manga.Description)

	return s
}
