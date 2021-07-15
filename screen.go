package main

import (
	nato "github.com/kaikaew13/manganato-api"
	"github.com/kaikaew13/manganato-cli/views"
)

var screen Screen

type Screen struct {

	// gocui's View with some extra fields
	sb *views.SearchBar
	sl *views.SearchList
	md *views.MangaDetails
	cl *views.ChapterList

	// manganato-api's Searcher struct
	searcher nato.Searcher
}
