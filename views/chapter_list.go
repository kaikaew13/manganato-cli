package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const chapterListName = "ChapterList"

// contains *gocui.View and data related
// to Chapter eg. Chapters (slice of Chapter)
type ChapterList struct {
	View        *gocui.View
	Name        string
	MangaName   string
	MangaID     string
	Chapters    []nato.Chapter
	NameToIDMap map[string]string
	OriginX     int
	OriginY     int
}

// initiates ChapterList and sets ChapterList view by calling g.SetView
func GetChapterList(maxX, maxY int, g *gocui.Gui) (*ChapterList, error) {
	cl := ChapterList{}
	x0, y0, x1, y1 := cl.GetCoords(maxX, maxY)

	clView, err := g.SetView(chapterListName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	clView.Title = chapterListName
	clView.SelFgColor = gocui.ColorBlack
	clView.SelBgColor = gocui.ColorGreen
	clView.BgColor = gocui.ColorBlack
	clView.FgColor = gocui.ColorWhite
	clView.Highlight = true
	clView.Editable = true
	clView.Wrap = true
	clView.Editor = readOnlyEditor

	cl.View = clView
	cl.Name = chapterListName
	cl.Chapters = make([]nato.Chapter, 0)
	cl.NameToIDMap = make(map[string]string)
	cl.OriginX, cl.OriginY = clView.Origin()

	return &cl, err
}

// returns a dimension relavtive to screen's width and height
func (cl *ChapterList) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return maxX / 2, (maxY - SearchBarHeight - 1) / 2, maxX - 1, maxY - SearchBarHeight - 2
}

// formats chapters into a table like format
func (cl *ChapterList) FormatChapters() string {
	s := "\n\t\t\tCHAPTER NAME\t\t\t\t\tVIEWS\t\t\tUPLOADED\n\n"

	for _, chapter := range cl.Chapters {
		s += fmt.Sprintf(
			"\t\t\t%s %s\t\t\t\t%s\t\t\t%s\n\n",
			Selector, chapter.ChapterName, chapter.Views, chapter.Uploaded,
		)

		cl.NameToIDMap[chapter.ChapterName] = chapter.ID
	}

	return s
}
