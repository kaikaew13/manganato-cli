package views

import (
	nato "github.com/kaikaew13/manganato-api"
)

type SearchList struct {
	BaseType
	Mangas      []nato.Manga
	NameToIDMap map[string]string
}
