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
		v.MoveCursor(0, -1, false)
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
		if len(s) > 0 && s[0] == '-' {
			v.MoveCursor(0, tmpy-y, false)
			return
		}
		tmpy++
	}

}
