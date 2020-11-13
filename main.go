package main

import (
	"fmt"

	"github.com/patryksandomierski/goditor/morestrings"
	"github.com/patryksandomierski/goditor/ui"
)

const appName = "Goditor"

func main() {
	fmt.Println(morestrings.ReverseChars("!dlrow ,olleH"))

	ui.DrawMainWindow(appName)
}
