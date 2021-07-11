package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const SearchListName string = "SearchList"

type SearchList struct {
	View        *gocui.View
	Mangas      []nato.Manga
	NameToIDMap map[string]string
}

func GetSearchList(maxX, maxY int, g *gocui.Gui) (*SearchList, error) {
	sl := SearchList{}
	x0, y0, x1, y1 := sl.GetCoords(maxX, maxY)

	slView, err := g.SetView(SearchListName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	slView.Title = SearchListName
	slView.SelFgColor = gocui.ColorBlack
	slView.SelBgColor = gocui.ColorGreen
	slView.Highlight = true
	slView.BgColor = gocui.ColorBlack
	slView.FgColor = gocui.ColorWhite
	slView.Editor = selectingEditor
	slView.Editable = true
	slView.Wrap = true

	sl.View = slView
	sl.Mangas = make([]nato.Manga, 0)
	sl.NameToIDMap = make(map[string]string)

	return &sl, err
}

func (sl *SearchList) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return 1, 1, maxX/2 - 1, maxY - SearchBarHeight - 2
}

func (sl *SearchList) FormatMangas() string {
	s := "\n\n"

	for _, mg := range sl.Mangas {
		if mg.Views != "" {
			s += fmt.Sprintf("	- %s\n			Author: %s\n			Views: %s\n\n", mg.Name, mg.Author.Name, mg.Views)
		} else {
			s += fmt.Sprintf("	- %s\n			Author: %s\n\n", mg.Name, mg.Author.Name)
		}
		sl.NameToIDMap[mg.Name] = mg.ID
	}

	return s
}
