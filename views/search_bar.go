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

	for i, v := range *sb.Commands {
		if v == removeNewline(cmd) {
			(*sb.Commands) = append((*sb.Commands)[:i], (*sb.Commands)[i+1:]...)
			break
		}
	}

	*sb.Commands = append(*sb.Commands, removeNewline(cmd))
}

func (sb *SearchBar) GetPrevCommand(cmd string) string {
	cmds := *sb.Commands

	if len(cmds) == 0 {
		return ""
	}

	if cmd == "" {
		return cmds[len(cmds)-1]
	}

	for i, v := range cmds {
		if v == removeNewline(cmd) {
			if i-1 < 0 {
				return ""
			}
			return cmds[i-1]
		}
	}

	return cmds[len(cmds)-1]
}

func (sb *SearchBar) GetNextCommand(cmd string) string {
	cmds := *sb.Commands

	if len(cmds) == 0 || cmd == "" {
		return ""
	}

	for i, v := range cmds {
		if v == removeNewline(cmd) {
			if i+1 == len(cmds) {
				return ""
			}
			return cmds[i+1]
		}
	}

	return ""
}

func removeNewline(s string) string {
	return s[:len(s)-1]
}
