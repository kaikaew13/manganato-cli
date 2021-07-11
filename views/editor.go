package views

import (
	"strings"

	"github.com/jroimartin/gocui"
)

var selectingEditor gocui.Editor = gocui.EditorFunc(edit)

func edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyArrowDown:
		findNextLine(v)
	case key == gocui.KeyArrowUp:
		findPrevLine(v)
		// case key == gocui.KeyEnter:
		// 	pickLine(v)
	}
}

func findNextLine(v *gocui.View) {
	_, y := v.Cursor()
	tmpy := y + 1

	for {
		s, err := v.Line(tmpy)
		if err != nil {
			break
		}

		s = strings.TrimSpace(s)
		if len(s) > 2 && s[:3] == "[x]" {
			v.MoveCursor(0, tmpy-y, false)
			return
		}
		tmpy++
	}
}

func findPrevLine(v *gocui.View) {
	_, y := v.Cursor()
	tmpy := y - 1

	for {
		s, err := v.Line(tmpy)
		if err != nil {
			break
		}

		s = strings.TrimSpace(s)
		if len(s) > 2 && s[:3] == "[x]" {
			v.MoveCursor(0, tmpy-y, false)
			return
		}
		tmpy--
	}
}

// might be better if control key enter from keybindings() in root dir

// func pickLine(v *gocui.View) {
// 	_, y := v.Cursor()

// 	s, _ := v.Line(y)
// 	s = strings.TrimSpace(s)

// 	if len(s) > 2 && s[:3] == "[x]" {
// 		name := s[4:]

// 		log.Panicln(name)
// 	}
// }
