package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

const (
	manganatoURL          string = "https://readmanganato.com/"
	modalViewName         string = "Modal"
	searchCommand         string = "search"
	searchByAuthorCommand        = searchCommand + "-author"
	searchByGenreCommand         = searchCommand + "-genre"
)

var viewNames = []string{
	views.SearchBarName,
	views.SearchListName,
	views.MangaDetailsName,
	views.ChapterListName,
}

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

	// has to be checked before searchCommand because has 'search' as
	// a prefix
	switch {
	case strings.HasPrefix(s, searchByAuthorCommand):
		valid = true
		cmd = searchByAuthorCommand
		args = s[len(searchByAuthorCommand)+1 : len(s)-1]
	case strings.HasPrefix(s, searchByGenreCommand):
		valid = true
		cmd = searchByGenreCommand
		args = s[len(searchByGenreCommand)+1 : len(s)-1]
	case strings.HasPrefix(s, searchCommand):
		valid = true
		cmd = searchCommand
		args = s[len(searchCommand)+1 : len(s)-1]
	}

	return
}

func runCommand(cmd, args string) error {
	var err error
	var mgs *[]nato.Manga

	switch cmd {
	case searchCommand:
		mgs, err = screen.searcher.SearchManga(args)
	case searchByAuthorCommand:
		mgs, err = screen.searcher.PickAuthor(screen.sl.NameToIDMap[args])
	case searchByGenreCommand:
		mgs, err = screen.searcher.PickGenre(screen.md.NameToIDMap[args])
	}

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

	return nil
}

func processCommand(g *gocui.Gui, v *gocui.View) {
	var val bool
	var wg sync.WaitGroup

	wg.Add(1)

	g.Update(func(g *gocui.Gui) error {
		defer wg.Done()
		s := v.Buffer()

		x, y := v.Origin()
		if err := v.SetCursor(x, y); err != nil {
			return err
		}
		v.Clear()

		valid, cmd, args := validateCommand(s)
		if valid {
			val = valid
			screen.sb.SaveCommand(s)
			if err := runCommand(cmd, args); err != nil {
				return err
			}
		} else {
			return nil
		}

		err := closeModal(g)
		return err
	})

	wg.Wait()

	if !val {
		setModalMessage(g, "unknown command...")
	}
}

func trimViewLine(v *gocui.View) string {
	_, y := v.Cursor()

	s, _ := v.Line(y)

	return strings.TrimSpace(s)
}

func downloadChapter(s string) error {
	if strings.HasPrefix(s, views.Selector) {
		chapterName := removePref(s, len(views.Selector))
		chapterName = strings.Split(chapterName, "\t")[0]

		pgs, err := screen.searcher.ReadMangaChapter(
			screen.cl.MangaID,
			screen.cl.NameToIDMap[chapterName],
		)
		if err != nil {
			return err
		}

		setupDownloadPath(*pgs, chapterName)
	}

	return nil
}

func setupDownloadPath(pgs []nato.Page, chapterName string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	dirpath, err := getDirPath(user.HomeDir, chapterName)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(pgs))

	for _, pg := range pgs {

		go func(id, url string) {
			defer wg.Done()

			fp := filepath.Join(dirpath, fmt.Sprintf("%s.jpg", id))

			err = downloadPage(fp, url)
		}(pg.ID, pg.ImageURL)

		if err != nil {
			return err
		}
	}

	wg.Wait()

	return nil
}

func getDirPath(homedir, chapterName string) (dirpath string, err error) {
	dirpath = filepath.Join(
		homedir, "Desktop", "manganato-cli", screen.cl.MangaName,
		screen.cl.NameToIDMap[chapterName],
	)
	err = os.MkdirAll(dirpath, 0755)
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
	req.Header.Add("referer", manganatoURL)

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

func resetCursor(v *gocui.View) {
	switch v.Name() {
	case views.SearchListName:
		v.SetCursor(screen.sl.OriginX, screen.sl.OriginY)
		v.SetOrigin(screen.sl.OriginX, screen.sl.OriginY)
	case views.MangaDetailsName:
		v.SetCursor(screen.md.OriginX, screen.md.OriginY)
		v.SetOrigin(screen.md.OriginX, screen.md.OriginY)
	case views.ChapterListName:
		v.SetCursor(screen.cl.OriginX, screen.cl.OriginY)
		v.SetOrigin(screen.cl.OriginX, screen.cl.OriginY)
	}
}

func openModal(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if lv, err := g.SetView(modalViewName, maxX/2-10, maxY/2-2, maxX/2+10, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		lv.BgColor = gocui.ColorBlack
		lv.FgColor = gocui.ColorWhite
		lv.Write([]byte("\n\n\t\t\t\t\tLoading..."))

		g.Cursor = false
		g.SetViewOnTop(lv.Name())
		g.SetCurrentView(lv.Name())
	}

	return nil
}

func closeModal(g *gocui.Gui) error {
	lv, err := g.View(modalViewName)
	if err != nil {
		return err
	}
	lv.Clear()

	g.DeleteView(lv.Name())
	g.SetCurrentView(views.SearchBarName)
	g.Cursor = true
	return nil
}

// sets a modal message to msg then close the modal
// after one second
func setModalMessage(g *gocui.Gui, msg string) {
	var wg sync.WaitGroup
	wg.Add(1)

	g.Update(func(g *gocui.Gui) error {
		defer wg.Done()
		lv, err := g.View(modalViewName)
		if err != nil {
			return err
		}
		lv.Clear()
		lv.Write([]byte(msg))
		return nil
	})

	wg.Wait()

	g.Update(func(g *gocui.Gui) error {
		time.Sleep(time.Second)
		err := closeModal(g)
		return err
	})
}
