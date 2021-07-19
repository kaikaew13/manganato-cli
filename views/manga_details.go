package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const mangaDetailsName string = "MangaDetails"

// contains *gocui.View and data related
// to Manga eg. nato.Manga's struct
type MangaDetails struct {
	View        *gocui.View
	Name        string
	Manga       nato.Manga
	NameToIDMap map[string]string
	OriginX     int
	OriginY     int
}

// initiates MangaDetails and sets MangaDetails view by calling g.SetView
func GetMangaDetails(maxX, maxY int, g *gocui.Gui) (*MangaDetails, error) {
	md := MangaDetails{}
	x0, y0, x1, y1 := md.GetCoords(maxX, maxY)

	mdView, err := g.SetView(mangaDetailsName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	mdView.Title = mangaDetailsName
	mdView.SelFgColor = gocui.ColorBlack
	mdView.SelBgColor = gocui.ColorYellow
	mdView.BgColor = gocui.ColorBlack
	mdView.FgColor = gocui.ColorWhite
	mdView.Highlight = true
	mdView.Editable = true
	mdView.Wrap = true
	mdView.Editor = readOnlyEditor

	md.View = mdView
	md.Name = mangaDetailsName
	md.Manga = nato.Manga{}
	md.NameToIDMap = make(map[string]string)
	md.OriginX, md.OriginY = mdView.Origin()

	return &md, err
}

// returns a dimension relavtive to screen's width and height
func (md *MangaDetails) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return maxX / 2, 1, maxX - 1, ((maxY - SearchBarHeight - 1) / 2) - 1
}

// formats manga into more readable format
func (md *MangaDetails) FormatManga() string {
	s := "\n\n"

	s += fmt.Sprintf("\t\t\u001b[33mTITLE:\u001b[0m %s\n\n", md.Manga.Name)
	s += fmt.Sprintf("\t\t\u001b[33mALT_NAME:\u001b[0m %s\n\n", md.Manga.Alternatives)
	s += fmt.Sprintf("\t\t\u001b[33mSTATUS:\u001b[0m %s\n\n", md.Manga.Status)

	var genres string
	for _, v := range md.Manga.Genres {
		genres += v.GenreName + "\t"

		// maps genres to ids so user can search by genre
		md.NameToIDMap[v.GenreName] = v.ID
		// improves user experience by allowing to search with
		// all lower case or all upper case
		md.NameToIDMap[strings.ToLower(v.GenreName)] = v.ID
		md.NameToIDMap[strings.ToUpper(v.GenreName)] = v.ID
	}

	s += fmt.Sprintf("\t\t\u001b[33mGENRES:\u001b[0m %s\n\n", genres)
	s += fmt.Sprintf("\t\t\u001b[33mAUTHOR:\u001b[0m %s\n\n", md.Manga.Author.Name)
	s += fmt.Sprintf("\t\t\u001b[33mUPDATED:\u001b[0m %s\n\n", md.Manga.Updated)
	s += fmt.Sprintf("\t\t\u001b[33mVIEWS:\u001b[0m %s\n\n", md.Manga.Views)
	s += fmt.Sprintf("\t\t\u001b[33mRATING:\u001b[0m %s\n\n", md.Manga.Rating)
	s += fmt.Sprintf("\t\t\u001b[33mDESCRIPTION:\u001b[0m %s\n\n", md.Manga.Description)

	return s
}
