# manganato-cli

Unofficial Manganato's manga downloader CUI written in Go.

## dependencies

- [gocui](https://github.com/jroimartin/gocui) for CUI
- [manganato-api](https://github.com/kaikaew13/manganato-api) for web scraped API (use [gocolly](https://github.com/gocolly/colly) as a web scraper)

## install

```
git clone https://github.com/kaikaew13/manganato-cli.git
cd manganato-cli
go build
```

## features

1. search mangas by name
2. search mangas by author name (only works with the author names that appear in search list at least once)
3. search mangas by genre (only works with the genre names that appear in the manga details at least once)
4. select a specific manga from the search list and:
   - display its details in manga details view
   - display its list of chapters in chapterlist view
5. select a chapter and download it to your own computer

## usage

to start the app:

```
./manganato-cli
```

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
