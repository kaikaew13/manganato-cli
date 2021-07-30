package main

import (
	"time"

	"github.com/jroimartin/gocui"
)

// successfully exits the program
func quit(g *gocui.Gui, v *gocui.View) error {

	// gocui's special error to indicates
	// program exits successfully
	return gocui.ErrQuit
}

// switches between views in clockwise direction
func switchView(g *gocui.Gui, v *gocui.View) error {
	resetCursor(v)

	vNames := getViewNames(g)
	for i, name := range vNames {
		if name == v.Name() {
			if i == len(vNames)-1 {
				g.SetCurrentView(vNames[0])
				break
			}

			g.SetCurrentView(vNames[i+1])
			break
		}
	}

	return nil
}

// switches between views in anti-clockwise direction
func reverseSwitchView(g *gocui.Gui, v *gocui.View) error {
	resetCursor(v)

	vNames := getViewNames(g)
	for i, name := range vNames {
		if name == v.Name() {
			if i == 0 {
				g.SetCurrentView(vNames[len(vNames)-1])
				break
			}

			g.SetCurrentView(vNames[i-1])
			break
		}
	}

	return nil
}

// starts processing command typed in SearchBar view
func enterCommand(g *gocui.Gui, v *gocui.View) error {
	if err := openModal(g); err != nil {
		return err
	}

	// must use go routine when processing command
	// so it does not block openModal function
	// (or else user wont see the loading modal)
	go processCommand(g, v)

	return nil
}

// returns previously entered command to SearchBar view
func getPrevCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()
	s = screen.sb.GetPrevCommand(s)
	if len(s) == 0 {
		resetCursor(v)
	}

	v.Clear()
	v.Write([]byte(s))

	return nil
}

// exact opposite of getPrevCommand
func getNextCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()
	s = screen.sb.GetNextCommand(s)
	if len(s) == 0 {
		resetCursor(v)
	}

	v.Clear()
	v.Write([]byte(s))

	return nil
}

// picks a manga and loads its info onto
// MangaDetails and ChapterList views
func pickManga(g *gocui.Gui, v *gocui.View) error {
	s := trimViewLine(v)

	err := getMangaScreen(s)

	return err
}

// picks a chapter and starts downloading its pages
func pickChapter(g *gocui.Gui, v *gocui.View) error {
	if err := openModal(g); err != nil {
		return err
	}

	done := make(chan bool)
	timer := time.NewTimer(time.Second * time.Duration(downloadTimeoutSecond))

	// must run downloading process in
	// go routine or else the it will
	// block the openModal so loading modal
	// will not be shown to the user
	go func() {
		s := trimViewLine(v)
		prepDownloadChapter(s)
		done <- true
	}()

	// in case downloading takes longer than
	// downloadTimeoutSecond, close the modal
	// and continue to download in background
	go func() {
		select {
		case <-timer.C:
			setClosingMessage(g, "continuing to download\nin background...")
			return
		case <-done:
			g.Update(func(g *gocui.Gui) error {
				err := closeModal(g)
				return err
			})
		}
	}()

	return nil
}
