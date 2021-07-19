package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

const (
	manganatoURL  string = "https://readmanganato.com/"
	modalViewName string = "Modal"

	// list of commands
	searchCommand         string = "search"
	searchByAuthorCommand        = searchCommand + "-author"
	searchByGenreCommand         = searchCommand + "-genre"
)

// is called on every iteration of g.Mainloop.
// adjusts views dimension dynamically in case
// user resize the screen
func layout(g *gocui.Gui) error {

	// maxX and maxY will change depending on the screen size
	maxX, maxY := g.Size()

	// each View struct (defined in ./views)
	// has a method to calculate its new dimension
	x0, y0, x1, y1 := screen.sb.GetCoords(maxX, maxY)
	if _, err := g.SetView(screen.sb.Name, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.sl.GetCoords(maxX, maxY)
	if _, err := g.SetView(screen.sl.Name, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.md.GetCoords(maxX, maxY)
	if _, err := g.SetView(screen.md.Name, x0, y0, x1, y1); err != nil {
		return err
	}

	x0, y0, x1, y1 = screen.cl.GetCoords(maxX, maxY)
	if _, err := g.SetView(screen.cl.Name, x0, y0, x1, y1); err != nil {
		return err
	}

	return nil
}

// returns a string slice with names of every views
func getViewNames(g *gocui.Gui) []string {
	viewNames := []string{}

	for _, v := range g.Views() {
		viewNames = append(viewNames, v.Name())
	}

	return viewNames
}

// uses manganato-cli's Searcher to fetch
// latest updated mangas from manganato
// then displays to SearchList view
func getInitialScreen() error {
	mgs, err := screen.searcher.SearchLatestUpdatedManga()
	if err != nil {
		return err
	}
	screen.sl.Mangas = *mgs

	screen.sl.View.Write([]byte(screen.sl.FormatMangas()))

	return nil
}

// uses manganato-cli's Searcher to fetch
// a specific manga's info then displays
// to MangaDetails and ChapterList views
func getMangaScreen(s string) error {
	// checks if the line selected is a valid manga name
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

// checks whether command entered in SearchBar is valid.
// if valid, the command and arguments are returned
func validateCommand(s string) (valid bool, cmd, args string) {

	// has to be checked before searchCommand because has 'search'
	// as a prefix
	switch {
	case strings.HasPrefix(s, searchByAuthorCommand):

		// avoid user to just type command without args
		if len(s) <= len(searchByAuthorCommand)+1 {
			return
		}

		valid = true
		cmd = searchByAuthorCommand
		// extracts what's after the command
		args = removePref(s, len(searchByAuthorCommand))
	case strings.HasPrefix(s, searchByGenreCommand):
		if len(s) <= len(searchByGenreCommand)+1 {
			return
		}

		valid = true
		cmd = searchByGenreCommand
		args = removePref(s, len(searchByGenreCommand))
	case strings.HasPrefix(s, searchCommand):
		if len(s) <= len(searchCommand)+1 {
			return
		}

		valid = true
		cmd = searchCommand
		args = removePref(s, len(searchCommand))
	}

	return
}

// starts fetching data if the command is valid
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

	// must clear all buffer in the SearchList
	// before writing new data in
	screen.sl.View.Clear()
	if err == nato.ErrPageNotFound {
		screen.sl.View.Write([]byte(nato.ErrPageNotFound.Error()))
		return nil
	}

	// writes new data into Searchlist view
	screen.sl.Mangas = *mgs
	s := screen.sl.FormatMangas()
	screen.sl.View.Write([]byte(s))

	return nil
}

// validate the command and executes them
//
// this function gets called as a go routine
func processCommand(g *gocui.Gui, v *gocui.View) {
	var val bool

	wg.Add(1)

	g.Update(func(g *gocui.Gui) error {
		defer wg.Done()
		s := v.Buffer()
		s = strings.TrimSpace(s)

		// returns the cursor to its origin
		// and clear the SearchBar after user
		// pressed enter
		x, y := v.Origin()
		if err := v.SetCursor(x, y); err != nil {
			return err
		}
		v.Clear()

		valid, cmd, args := validateCommand(s)
		if valid {
			val = valid
			// saves a valid command to SearchBar's Commands slice
			screen.sb.SaveCommand(s)
			if err := runCommand(cmd, args); err != nil {
				return err
			}
		} else {
			return nil
		}

		// close the loading modal after
		// finished executing the command
		err := closeModal(g)
		return err
	})

	// wait for val to be defined in g.Update
	wg.Wait()

	if !val {
		// in case of invalid command display this message
		// as a closing message
		setClosingMessage(g, "unknown command...")
	}
}

// trim all prefixed and suffixed space
// from the selected line of the view
func trimViewLine(v *gocui.View) string {
	_, y := v.Cursor()
	s, _ := v.Line(y)

	return strings.TrimSpace(s)
}

// prepares to download pages from the selected chapter
func prepDownloadChapter(s string) error {
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

		downloadChapter(*pgs, chapterName)
	}

	return nil
}

func downloadChapter(pgs []nato.Page, chapterName string) error {

	// user has user.HomeDir which will be used
	// to specify the download path
	user, err := user.Current()
	if err != nil {
		return err
	}

	dirpath, err := getDirPath(user.HomeDir, chapterName)
	if err != nil {
		return err
	}

	wg.Add(len(pgs))

	// downloads each page concurrently
	for _, pg := range pgs {

		go func(id, url string) {
			defer wg.Done()

			// each page downloaded will have a name of <id>.jpg
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

// specify download path to Desktop/manganato-cli
// if the directory does not exist then create a new one
// with that name
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

// makes a request to manganato to download
// the page and saves it to the specified filepath
func downloadPage(fp, url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	// adds header so the request wont be blocked
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

// resets cursor to its origin (for SearchList, MangaDetails
// and ChapterList) after switching views
func resetCursor(v *gocui.View) {
	switch v.Name() {
	case screen.sl.Name:
		v.SetCursor(screen.sl.OriginX, screen.sl.OriginY)
		v.SetOrigin(screen.sl.OriginX, screen.sl.OriginY)
	case screen.md.Name:
		v.SetCursor(screen.md.OriginX, screen.md.OriginY)
		v.SetOrigin(screen.md.OriginX, screen.md.OriginY)
	case screen.cl.Name:
		v.SetCursor(screen.cl.OriginX, screen.cl.OriginY)
		v.SetOrigin(screen.cl.OriginX, screen.cl.OriginY)
	case screen.sb.Name:
		v.SetCursor(screen.sb.OriginX, screen.sb.OriginY)
		v.SetOrigin(screen.sb.OriginX, screen.sb.OriginY)
	}
}

// opens modal when fetching data or
// downloading files
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
	// must clear buffers in the modal
	// before deleting the view
	lv.Clear()

	g.DeleteView(lv.Name())
	// sets current view back to SearchBar
	g.SetCurrentView(screen.sb.Name)
	g.Cursor = true
	return nil
}

// sets a modal message to msg then close
// the modal after one second
func setClosingMessage(g *gocui.Gui, msg string) {
	wg.Add(1)

	g.Update(func(g *gocui.Gui) error {
		defer wg.Done()
		lv, err := g.View(modalViewName)
		if err != nil {
			return err
		}
		lv.Clear()
		// writes a new message to the modal
		// after clearing the old one
		lv.Write([]byte(msg))
		return nil
	})

	// waits for message updating process
	wg.Wait()

	// then closes the modal after one second
	g.Update(func(g *gocui.Gui) error {
		time.Sleep(time.Second)
		err := closeModal(g)
		return err
	})
}
