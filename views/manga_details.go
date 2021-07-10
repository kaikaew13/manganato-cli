package views

import (
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
	mdView, err := g.SetView(MangaDetailsName, maxX/2, 1, maxX-1, ((maxY-SearchBarHeight-1)/2)-1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	mdView.Title = MangaDetailsName
	mdView.SelFgColor = gocui.ColorGreen
	mdView.BgColor = gocui.ColorBlack
	mdView.FgColor = gocui.ColorWhite

	md := MangaDetails{
		View: mdView,
	}
	return &md, err
}
