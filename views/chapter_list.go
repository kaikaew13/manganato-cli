package views

import (
	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const ChapterListName = "ChapterList"

type ChapterList struct {
	View        *gocui.View
	Chapters    []nato.Chapter
	NameToIDMap map[string]string
}

func GetChapterList(maxX, maxY int, g *gocui.Gui) (*ChapterList, error) {
	clView, err := g.SetView(ChapterListName, maxX/2, (maxY-SearchBarHeight-1)/2, maxX-1, maxY-SearchBarHeight-2)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	clView.Title = ChapterListName
	clView.SelFgColor = gocui.ColorGreen
	clView.BgColor = gocui.ColorBlack
	clView.FgColor = gocui.ColorWhite

	cl := ChapterList{
		View:        clView,
		Chapters:    make([]nato.Chapter, 0),
		NameToIDMap: make(map[string]string),
	}
	return &cl, err
}
