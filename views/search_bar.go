package views

import (
	"github.com/jroimartin/gocui"
)

const (
	SearchBarHeight int    = 2
	SearchBarName   string = "SearchBar"
)

type SearchBar struct {
	View     *gocui.View
	Commands *[]string
}

func GetSearchBar(maxX, maxY int, g *gocui.Gui) (*SearchBar, error) {
	sbView, err := g.SetView(SearchBarName, 1, maxY-SearchBarHeight-1, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	sbView.Title = SearchBarName
	sbView.SelFgColor = gocui.ColorGreen
	sbView.BgColor = gocui.ColorBlack
	sbView.FgColor = gocui.ColorWhite
	sbView.Editable = true

	cmds := make([]string, 0)

	sb := SearchBar{
		View:     sbView,
		Commands: &cmds,
	}
	return &sb, err
}

func (sb *SearchBar) SaveCommand(cmd string) {
	if cmd == "" {
		return
	}

	*sb.Commands = append(*sb.Commands, cmd[:len(cmd)-1])
}
