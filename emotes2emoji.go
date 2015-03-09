package main

import (
	"fmt"
	"os"

	"code.google.com/p/getopt"

	"github.com/ddavison/emotes2emoji/converter"
	"github.com/ddavison/emotes2emoji/hooker"
)

// Entry-point to the program.
func main() {
	installHook := getopt.BoolLong("install-hook", 0, "Install git commit-msg hook")
	listEmoji := getopt.BoolLong("list", 0, "List the available emotes")
	help := getopt.BoolLong("help", 0, "Show usage")
	getopt.Parse()

	if *help {
		getopt.Usage()
		os.Exit(0)
	}

	if *listEmoji {
		converter.ListEmotes()
		os.Exit(0)
	}

	if *installHook {
		hooker.SellSelf() // install the git hook
		os.Exit(0)
	}

	fmt.Println(converter.Convert(getopt.Args()[0]))
}
