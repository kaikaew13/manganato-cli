package views

import (
	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

type ChapterList struct {
	View        *gocui.View
	Chapters    []nato.Chapter
	NameToIDMap map[string]string
}
