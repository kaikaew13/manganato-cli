package views

import (
	"strings"

	"github.com/jroimartin/gocui"
)

// used to determine which line is
// selectable. Selector will be the
// prefix of every selectable line
const Selector string = "[x]"

// editor for SearchList, MangaDetails and ChapterList views,
// can only use arrow keys to scroll through the views data
// and select a line by pressing enter
var readOnlyEditor gocui.Editor = gocui.EditorFunc(edit)

// defines the functionality of each key pressed
func edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyArrowDown, ch == 's':
		moveDown(v)
	case key == gocui.KeyArrowUp, ch == 'w':
		moveUp(v)
	case key == gocui.KeyArrowLeft, ch == 'a':
		moveHorizontally(v, -1)
	case key == gocui.KeyArrowRight, ch == 'd':
		moveHorizontally(v, 1)
	}
}

// only works in MangaDetails view
func moveHorizontally(v *gocui.View, dir int) {
	if v.Name() != mangaDetailsName {
		return
	}

	v.MoveCursor(dir, 0, false)
}

// for MangaDetails view, just move down a line.
// as for SearchList and ChapterList view,
// finds the next line with a Selector as a prefix
func moveDown(v *gocui.View) {
	_, y := v.Cursor()

	// case 1: MangaDetails view
	if v.Name() == mangaDetailsName {
		v.MoveCursor(0, 1, false)
		return
	}

	// case 2: SearchList and ChapterList view
	tmpy := y + 1

	for {
		s, err := v.Line(tmpy)
		if err != nil {
			break
		}

		s = strings.TrimSpace(s)
		if len(s) > 2 && s[:3] == Selector {
			v.MoveCursor(0, tmpy-y, false)
			return
		}
		tmpy++
	}
}

// same with moveDown but in opposite direction
func moveUp(v *gocui.View) {
	_, y := v.Cursor()

	if v.Name() == mangaDetailsName {
		v.MoveCursor(0, -1, false)
		return
	}

	tmpy := y - 1

	for {
		s, err := v.Line(tmpy)
		if err != nil {
			break
		}

		s = strings.TrimSpace(s)
		if len(s) > 2 && s[:3] == Selector {
			v.MoveCursor(0, tmpy-y, false)
			return
		}
		tmpy--
	}
}
