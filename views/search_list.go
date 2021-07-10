package views

import (
	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const SearchListName string = "SearchList"

type SearchList struct {
	View        *gocui.View
	Mangas      []nato.Manga
	NameToIDMap map[string]string
}

func GetSearchList(maxX, maxY int, g *gocui.Gui) (*SearchBar, error) {
	slView, err := g.SetView(SearchListName, 1, 1, maxX/2-1, maxY-SearchBarHeight-1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	slView.Title = SearchListName
	g.SelFgColor = gocui.ColorGreen
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	sb := SearchBar{
		View: slView,
	}
	return &sb, nil
}
