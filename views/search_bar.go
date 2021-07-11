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

func (sb *SearchBar) GetPrevCommand(cmd string) string {
	if len(*sb.Commands) == 0 {
		return ""
	}

	cmds := *sb.Commands

	if cmd == "" {
		return cmds[len(cmds)-1]
	}

	for i, v := range cmds {
		if v == cmd[:len(cmd)-1] {
			index := (i - 1) % len(cmds)
			if index < 0 {
				index += len(cmds)
			}
			return cmds[index]
		}
	}

	return cmds[len(cmds)-1]
}
