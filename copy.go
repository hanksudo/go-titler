package main

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func Copy(title, url string) {
	clipboard.WriteAll(fmt.Sprintf("- [%s](%s)", title, url))
}
