package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print(`Usage: gotitler <url>

Example:
  $ gotitler https://github.com
  $ gotitler cnn.com`)
		os.Exit(0)
	}
	url := os.Args[1]

	res := Fetch(url)
	title := Parse(res)
	Copy(title, res.Request.URL.String())
}
