# manganato-cli

Unofficial Manganato's manga downloader CUI and [CLI](https://github.com/kaikaew13/manganato-cli/blob/main/CLI.md) written in Go.

**note:** currently works on Mac and Linux, for Windows, please run via Docker and WSL (see [docker](#for-docker))

## preview

![mangagif](https://user-images.githubusercontent.com/77256757/126114177-2c9fbbff-6804-4d9c-a350-e5512eea2240.gif)

**for the preview of CLI mode see [CLI.md](https://github.com/kaikaew13/manganato-cli/blob/main/CLI.md)**

## dependencies

- [gocui](https://github.com/jroimartin/gocui) for CUI
- [manganato-api](https://github.com/kaikaew13/manganato-api) for web scraped API (use [gocolly](https://github.com/gocolly/colly) as a web scraper)

## install

```
git clone https://github.com/kaikaew13/manganato-cli.git
cd manganato-cli
go build
./manganato-cli
```

### for docker

```
git clone https://github.com/kaikaew13/manganato-cli.git
cd manganato-cli
make docker_build
make docker_run
```

## features

1. search mangas by name
2. search mangas by author name (only works with the author names that appear in search list at least once)
3. search mangas by genre (only works with the genre names that appear in the manga details at least once)
4. select a specific manga from the search list and:
   - display its details in manga details view
   - display its list of chapters in chapterlist view
5. select a chapter and download it to your own computer

**note:** the downloaded chapters can be found in Desktop/manganato-cli directory

## usage

list of commands:

- `search <manga name>` : search mangas by its name
- `search-author <author's name>` : search mangas by the author's name
- `search-genre <genre>` : search mangas by genre

keybindings:
| keys | description | views |
|------------------------------|--------------------------------------|---------------------------------------|
| <kbd>Ctrl</kbd>+<kbd>C</kbd> | exit the program | all |
| <kbd>Tab</kbd> | switch between views(clockwise) | all |
| <kbd>`</kbd> | switch between views(anti-clockwise) | all |
| <kbd>Enter</kbd> | search | SearchBar |
| <kbd>Enter</kbd> | select a manga/chapter | SearchList, ChapterList |
| <kbd>&uarr;</kbd> | get previous command | SearchBar |
| <kbd>&uarr;</kbd> | move the cursor up | SearchList, MangaDetails, ChapterList |
| <kbd>&darr;</kbd> | get following command | SearchBar |
| <kbd>&darr;</kbd> | move the cursor down | SearchList, MangaDetails, ChapterList |
| <kbd>&larr;</kbd> | move the cursor to the left | MangaDetails |
| <kbd>&rarr;</kbd> | move the cursor to the right | MangaDetails |

## views

<img width="1440" alt="Screen Shot 2564-07-19 at 13 43 47" src="https://user-images.githubusercontent.com/77256757/126114728-1aeb5fa8-33f6-4428-b756-417bca04cac4.png">

- SearchBar: allows user to search for manga
- SearchList: displays a list of mangas
- MangaDetails: displays details of the manga user picked, example: alternative names, genres, views, etc
- ChapterList: displays a list of chapters of the manga user picked

## credits

- [gocui](https://github.com/jroimartin/gocui) for CUI
- [gocolly](https://github.com/gocolly/colly) for web scraper
- [mellotanica](https://github.com/mellotanica) for CLI mode
