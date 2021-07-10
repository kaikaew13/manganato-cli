package views

import (
	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

type MangaDetails struct {
	View        *gocui.View
	Manga       nato.Manga
	NameToIDMap map[string]string
}
