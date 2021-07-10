package main

import (
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

var screen Screen

type Screen struct {
	sb       *views.SearchBar
	sl       *views.SearchList
	md       *views.MangaDetails
	cl       *views.ChapterList
	searcher nato.Searcher
}
