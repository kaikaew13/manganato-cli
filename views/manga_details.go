package views

import nato "github.com/kaikaew13/manganato-api"

type MangaDetails struct {
	BaseType
	Manga       nato.Manga
	NameToIDMap map[string]string
}
