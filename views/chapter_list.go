package views

import (
	nato "github.com/kaikaew13/manganato-api"
)

type ChapterList struct {
	BaseType
	Chapters    []nato.Chapter
	NameToIDMap map[string]string
}
