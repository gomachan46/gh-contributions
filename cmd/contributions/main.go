package main

import (
	"fmt"

	"os"
	contributions "github.com/gomachan46/gh-contributions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "do nothing\n")
		os.Exit(1)
	}

	contributions, err := contributions.Get(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, "start,end,contributions,current_streak\n")
	fmt.Fprintf(os.Stdout, contributions)
}
