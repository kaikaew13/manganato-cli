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
	slView, err := g.SetView(SearchListName, 1, 1, maxX/2-1, maxY-SearchBarHeight-2)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	slView.Title = SearchListName
	slView.SelFgColor = gocui.ColorGreen
	slView.BgColor = gocui.ColorBlack
	slView.FgColor = gocui.ColorWhite
	slView.Editable = true

	sl := SearchList{
		View:        slView,
		NameToIDMap: make(map[string]string),
	}
	return &sl, err
}

func (sl *SearchList) FormatMangas() string {
	s := ""

	for _, mg := range sl.Mangas {
		s += fmt.Sprintf("	- %s\n		Author: %s\n		Views: %s\n\n", mg.Name, mg.Author.Name, mg.Views)
		sl.NameToIDMap[mg.Name] = mg.ID
	}

	// for i := 0; i < 10; i++ {
	// 	s += fmt.Sprintf("	- %s\n		Author: %s\n\n", "nice", "not nice")
	// }

	return s
}
