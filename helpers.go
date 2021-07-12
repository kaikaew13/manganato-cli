package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

const (
	searchCommand string = "search"
	// searchByAuthorCommand  = searchCommand + "-author"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	x0, y0, x1, y1 := screen.sb.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.SearchBarName, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.sl.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.SearchListName, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.md.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.MangaDetailsName, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.cl.GetCoords(maxX, maxY)
	if _, err := g.SetView(views.ChapterListName, x0, y0, x1, y1); err != nil {
		return err
	}

	return nil
}

func getInitialScreen() error {
	mgs, err := screen.searcher.SearchLatestUpdatedManga()
	if err != nil {
		return err
	}
	screen.sl.Mangas = *mgs

	screen.sl.View.Write([]byte(screen.sl.FormatMangas()))

	return nil
}

func getMangaScreen(s string) error {
	if strings.HasPrefix(s, views.Selector) {
		mgName, mgId := getMangaNameAndID(s, len(views.Selector))

		mg, err := screen.searcher.PickManga(mgId)
		if err != nil {
			return err
		}

		screen.md.Manga = *mg
		s = screen.md.FormatManga()
		screen.md.View.Clear()
		screen.md.View.Write([]byte(s))

		screen.cl.MangaName = mgName
		screen.cl.MangaID = mgId
		screen.cl.Chapters = mg.Chapters
		s = screen.cl.FormatChapters()
		screen.cl.View.Clear()
		screen.cl.View.Write([]byte(s))
	}

	return nil
}

func getMangaNameAndID(s string, prefLen int) (mgName, mgId string) {
	mgName = removePref(s, prefLen)
	mgId = screen.sl.NameToIDMap[mgName]
	return
}

func removePref(s string, prefLen int) string {
	return s[prefLen+1:]
}

func validateCommand(s string) (valid bool, cmd, args string) {
	if strings.HasPrefix(s, searchCommand) {
		valid = true
		cmd = searchCommand
		args = s[len(searchCommand)+1:]

		return
	}

	return
}

func runCommand(cmd, args string) error {
	switch cmd {
	case searchCommand:
		mgs, err := screen.searcher.SearchManga(args)
		if err != nil && err != nato.ErrPageNotFound {
			return err
		}

		screen.sl.View.Clear()
		if err == nato.ErrPageNotFound {
			screen.sl.View.Write([]byte(nato.ErrPageNotFound.Error()))
			return nil
		}

		screen.sl.Mangas = *mgs
		s := screen.sl.FormatMangas()
		screen.sl.View.Write([]byte(s))
	}

	return nil
}

func trimLine(v *gocui.View) string {
	_, y := v.Cursor()

	s, _ := v.Line(y)

	return strings.TrimSpace(s)
}

func downloadChapter(s string) error {
	if strings.HasPrefix(s, views.Selector) {
		chapterName := removePref(s, len(views.Selector))
		chapterName = strings.Split(chapterName, "\t")[0]

		pgs, err := screen.searcher.ReadMangaChapter(screen.cl.MangaID, screen.cl.NameToIDMap[chapterName])
		if err != nil {
			return err
		}

		setupDownloadPath(*pgs)
	}

	return nil
}

func setupDownloadPath(pgs []nato.Page) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	dirpath, err := getDirPath(user.HomeDir)
	if err != nil {
		return err
	}

	for _, pg := range pgs {
		fp := filepath.Join(dirpath, fmt.Sprintf("%s.jpg", pg.ID))

		err = downloadPage(fp, pg.ImageURL)
		if err != nil {
			return err
		}
	}

	return nil
}

func getDirPath(homedir string) (dirpath string, err error) {
	dirpath = filepath.Join(homedir, "Desktop", "manganato-cli")
	err = os.Mkdir(dirpath, 0755)
	if err != nil {
		return "", err
	}

	return dirpath, nil
}

func downloadPage(fp, url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("referer", "https://readmanganato.com/")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	return err
}
