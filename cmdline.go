package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	nato "github.com/kaikaew13/manganato-api"
)

type Chapter struct {
	chapter *nato.Chapter
	index   int
}

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
			manga, err := screen.searcher.PickManga(m.ID)
			if err == nil {
				out = append(out, fmt.Sprintf("[%v] %v (Author: %v, Chapters: %v)", manga.ID, manga.Name, manga.Author.Name, len(manga.Chapters)))
			} else {
				out = append(out, fmt.Sprintf("[%v] %v (Author: %v)", m.ID, m.Name, m.Author.Name))
			}
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

func indexToChapter(chapter int, manga nato.Manga) (*Chapter, error) {
	if chapter > len(manga.Chapters) || chapter < 1 {
		return nil, errors.New(fmt.Sprintf("chapter '%v' out of range (1 - %v)", chapter, len(manga.Chapters)))
	}
	return &Chapter{
		&manga.Chapters[len(manga.Chapters)-chapter],
		chapter,
	}, nil
}

func getChapterRange(minch, maxch int, manga nato.Manga) ([]*Chapter, error) {
	chapters := make([]*Chapter, maxch-minch)
	var err error
	for i := 0; i < len(chapters); i++ {
		chapters[i], err = indexToChapter(minch+i, manga)
		if err != nil {
			return nil, err
		}
	}
	return chapters, nil
}

func chapterspecToChapters(manga nato.Manga, chapterspec string) ([]*Chapter, error) {
	if strings.Contains(chapterspec, "-") {
		var minch int
		var maxch int
		var err error
		limits := strings.Split(chapterspec, "-")
		if limits[0] == "" {
			minch = 1
		} else {
			minch, err = strconv.Atoi(limits[0])
			if err != nil {
				return nil, err
			}
		}
		if limits[1] == "" {
			maxch = len(manga.Chapters)
		} else {
			maxch, err = strconv.Atoi(limits[1])
			if err != nil {
				return nil, err
			}
		}

		if minch == maxch {
			return chapterspecToChapters(manga, fmt.Sprintf("%d", minch))
		}

		if minch > maxch {
			tch := minch
			minch = maxch
			maxch = tch
		}
		return getChapterRange(minch, maxch, manga)
	}

	if strings.Contains(chapterspec, ",") {
		chptrs := strings.Split(chapterspec, ",")
		chapters := make([]*Chapter, len(chptrs))
		for i, c := range chptrs {
			ch, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			chapters[i], err = indexToChapter(ch, manga)
			if err != nil {
				return nil, err
			}
		}
		return chapters, nil
	}

	ch, err := strconv.Atoi(chapterspec)
	if err != nil {
		return nil, err
	}
	chapter, err := indexToChapter(ch, manga)
	if err != nil {
		return nil, err
	}
	return []*Chapter{chapter}, nil
}

func downloadChapters(mangaid string, chapterspec string, destination string, oneatatime, ignoreerrors bool) error {
	var err error = nil
	manga, err := screen.searcher.PickManga(mangaid)
	if err != nil {
		return err
	}

	chapters, err := chapterspecToChapters(*manga, chapterspec)
	if err != nil {
		return err
	}
	basedir := path.Join(destination, manga.Name)

	for _, chapter := range chapters {
		pgs, err := screen.searcher.ReadMangaChapter(manga.ID, chapter.chapter.ID)
		if err != nil {
			if ignoreerrors {
				fmt.Fprintf(os.Stderr, "Download error, chapter %v not found: %v\n)", chapter.index, err)
			} else {
				return err
			}
		}

		chdir := path.Join(basedir, fmt.Sprintf("%d", chapter.index))
		os.MkdirAll(chdir, 0755)

		fmt.Printf("downloading chapter: '%v' to %v\n", chapter.chapter.ChapterName, chdir)
		if oneatatime {
			err = downloadPages(*pgs, chdir)
		} else {
			err = downloadPagesNowait(*pgs, chdir)
		}
		if err != nil {
			if ignoreerrors {
				fmt.Fprintf(os.Stderr, "Download error (%v): %v\n)", chapter.index, err)
			} else {
				return err
			}
		}
	}

	if !oneatatime {
		wg.Wait()
	}

	return err
}
