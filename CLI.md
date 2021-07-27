# CLI mode

credits to [mellotanica](https://github.com/mellotanica)

## preview

![final](https://user-images.githubusercontent.com/77256757/127091031-a52e206c-c6ba-4f00-97b9-f0a30d6a1e93.gif)

## usage

list of flags:

- `--help` : provides a list of available flags with description
- `--search=<manga name>` : search manga based on title pattern
- `--manga-id=<manga id>` : download/show resources of selected manga id
- `--download=<manga chapter>` : download manga chapters ('-' to download all, chapters numbers to download specific chapters, comma-separated lists or dash-ranges to perform batch download)
- `--output=<path>` : downloaded images will be put in this path
- `--all-together` : download all chapters in parallel (may lead to errors for too much requests)
- `--ignore-errors` : ignore download errors and keep going
