package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

func TestValidateCommand(t *testing.T) {
	type test struct {
		description string
		command     string
		valid       bool
		cmd         string
		args        string
	}

	tests := []test{
		{
			description: "command for searching for mangas by name",
			command:     "search chainsaw man",
			valid:       true,
			cmd:         searchCommand,
			args:        "chainsaw man",
		},
		{
			description: "command for searching for mangas by author name",
			command:     "search-author tatsuki fujimoto",
			valid:       true,
			cmd:         searchByAuthorCommand,
			args:        "tatsuki fujimoto",
		},
		{
			description: "command for searching for mangas by genre",
			command:     "search-genre action",
			valid:       true,
			cmd:         searchByGenreCommand,
			args:        "action",
		},
		{
			description: "invalid command",
			command:     "seach chainsaw man",
			valid:       false,
			cmd:         "",
			args:        "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			valid, cmd, args := validateCommand(test.command)

			if valid != test.valid {
				t.Errorf("wanted valid to be %t, got %t", test.valid, valid)
			}

			if cmd != test.cmd {
				t.Errorf("wanted cmd to be %s, got %s", test.cmd, cmd)
			}

			if args != test.args {
				t.Errorf("wanted args to be %s, got %s", test.args, args)
			}
		})
	}
}

func TestGetInitialScreen(t *testing.T) {
	g, maxX, maxY := initHelper(t)
	getSearchListHelper(t, maxX, maxY, g)

	err := getInitialScreen()
	if err != nil {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	mgs, err := screen.searcher.SearchLatestUpdatedManga()
	if err != nil {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	if len(*mgs) != len(screen.sl.Mangas) {
		t.Errorf("wanted mangas with length %d, got %d", len(*mgs), len(screen.sl.Mangas))
	}

	g = nil
}

func TestGetMangaScreen(t *testing.T) {
	mgName, mgId := "Chainsaw Man", "dn980422"

	g, maxX, maxY := initHelper(t)
	getSearchListHelper(t, maxX, maxY, g)
	getMangaDetailsHelper(t, maxX, maxY, g)
	getChapterListHelper(t, maxX, maxY, g)

	screen.sl.NameToIDMap[mgName] = mgId

	err := getMangaScreen(fmt.Sprintf("%s %s", views.Selector, mgName))
	if err != nil {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	mg, err := screen.searcher.PickManga(mgId)
	if err != nil {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	if !reflect.DeepEqual(screen.md.Manga, *mg) {
		t.Errorf("manga fields are not equal:\n wanted %v, got %v", *mg, screen.md.Manga)
	}

	g = nil
}

func TestGetDirPath(t *testing.T) {
	dirpath, err := getDirPath("test", "1")
	if err != nil {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	_, err = os.Stat(dirpath)
	if err != nil && os.IsNotExist(err) {
		t.Error("directory does not exist")
	}
	os.RemoveAll("test")
}

func initHelper(t testing.TB) (g *gocui.Gui, maxX, maxY int) {
	t.Helper()

	screen.searcher = nato.NewSearcher()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		t.Errorf("not expected to have error: %s", err.Error())
	}
	maxX, maxY = g.Size()

	return
}

func getSearchListHelper(t testing.TB, maxX, maxY int, g *gocui.Gui) {
	t.Helper()

	sl, err := views.GetSearchList(maxX, maxY, g)
	if err != nil && err != gocui.ErrUnknownView {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	if sl == nil {
		t.Error("not expected sl to be nil")
	}
	screen.sl = sl
}

func getMangaDetailsHelper(t testing.TB, maxX, maxY int, g *gocui.Gui) {
	t.Helper()

	md, err := views.GetMangaDetails(maxX, maxY, g)
	if err != nil && err != gocui.ErrUnknownView {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	if md == nil {
		t.Error("not expected md to be nil")
	}
	screen.md = md
}

func getChapterListHelper(t testing.TB, maxX, maxY int, g *gocui.Gui) {
	t.Helper()

	cl, err := views.GetChapterList(maxX, maxY, g)
	if err != nil && err != gocui.ErrUnknownView {
		t.Errorf("not expected to have error: %s", err.Error())
	}

	if cl == nil {
		t.Error("not expected cl to be nil")
	}
	screen.cl = cl
}
