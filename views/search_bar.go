package views

import "github.com/jroimartin/gocui"

const (
	SearchBarHeight int    = 3
	SearchBarName   string = "SearchBar"
)

type SearchBar struct {
	View    *gocui.View
	Command string
}

func GetSearchBar(maxX, maxY int, g *gocui.Gui) (*SearchBar, error) {
	sbView, err := g.SetView(SearchBarName, 1, maxY-SearchBarHeight, maxX-1, maxY)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	sbView.Title = SearchBarName
	g.SelFgColor = gocui.ColorGreen
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	sb := SearchBar{
		View: sbView,
	}
	return &sb, nil
}
