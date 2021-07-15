package views

import (
	"strings"

	"github.com/jroimartin/gocui"
)

const Selector string = "[x]"

var readOnlyEditor gocui.Editor = gocui.EditorFunc(edit)

func edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyArrowDown, ch == 's':
		findNextLine(v)
	case key == gocui.KeyArrowUp, ch == 'w':
		findPrevLine(v)
	case key == gocui.KeyArrowLeft, ch == 'a':
		moveHorizontally(v, -1)
	case key == gocui.KeyArrowRight, ch == 'd':
		moveHorizontally(v, 1)
	}
}

func moveHorizontally(v *gocui.View, dir int) {
	if v.Name() != mangaDetailsName {
		return
	}

	v.MoveCursor(dir, 0, false)
}

func findNextLine(v *gocui.View) {
	_, y := v.Cursor()

	if v.Name() == mangaDetailsName {
		v.MoveCursor(0, 1, false)
		return
	}

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

func findPrevLine(v *gocui.View) {
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
