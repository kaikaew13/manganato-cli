package views

import "github.com/jroimartin/gocui"

var selectingEditor gocui.Editor = gocui.EditorFunc(edit)

func edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
	}
}
