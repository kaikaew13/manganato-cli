package main

import (
	"fmt"
	"strings"

	nato "github.com/kaikaew13/manganato-api"
)

func printList(output []string, err error) error {
	if err == nil {
		for _, l := range output {
			fmt.Println(strings.TrimSpace(l))
		}
	}
	return err
}

func outputMangas(mlist *[]nato.Manga, header string) []string {
	out := make([]string, 2)
	out[0] = header
	if mlist != nil {
		for _, m := range *mlist {
			out = append(out, fmt.Sprintf("[%v] %v (Author: %v, Chapters: %v)", m.ID, m.Name, m.Author.Name, len(m.Chapters)))
		}
	} else {
		out = append(out, "-- no manga found --")
	}
	return out
}

func searchManga(query string) (out []string, err error) {
	mangas, err := screen.searcher.SearchManga(query)
	if err == nil {
		out = outputMangas(mangas, fmt.Sprintf("Title search query: '%s'", query))
	}
	return
}

func listChapters(mangaid string) (out []string, err error) {
	manga, err := screen.searcher.PickManga(mangaid)
	if err == nil && manga != nil {
		out = make([]string, 2)
		out[0] = fmt.Sprintf("'%s' chapter list:", manga.Name)
		for i, ch := range manga.Chapters {
			out = append(out, fmt.Sprintf("[%v] %v (%d pages)", len(manga.Chapters)-i, ch.ChapterName, len(ch.Pages)))
		}
	}
	return
}

func downloadChapters(mangaid string, chapterspec string) error {
	return nil
}
