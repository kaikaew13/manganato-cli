package views

import "github.com/jroimartin/gocui"

const (
	SearchBarHeight int    = 2
	SearchBarName   string = "SearchBar"
)

type SearchBar struct {
	View    *gocui.View
	Command string
}

func GetSearchBar(maxX, maxY int, g *gocui.Gui) (*SearchBar, error) {
	sbView, err := g.SetView(SearchBarName, 1, maxY-SearchBarHeight-1, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		g.SetCurrentView(SearchBarName)
	}

	sbView.Title = SearchBarName
	sbView.SelFgColor = gocui.ColorGreen
	sbView.BgColor = gocui.ColorBlack
	sbView.FgColor = gocui.ColorWhite
	sbView.Editable = true

	sb := SearchBar{
		View: sbView,
	}
	return &sb, nil
}
