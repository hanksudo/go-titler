# gotitler

Auto copy title and link as markdown to clipboard

([Orginial](https://github.com/hanksudo/titler) version is made by nodejs)

## Feature

- Auto copy url's title and url to [Markdown link](https://www.markdownguide.org/basic-syntax/#links) format

## Installation

```bash
go get github.com/hanksudo/gotitler
```

## Usage

```bash
gotitler https://github.com
gotitler github.com
```

The content below will auto copied to your clipboard:

```bash
- [GitHub: Where the world builds software · GitHub](https://github.com/)
```

Rendered in Markdown

- [GitHub · Build software better, together.](https://github.com)
