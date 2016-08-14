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

	fmt.Fprint(os.Stdout, "username,from,to,total,currentStreak\n")
	args := os.Args[1:]
	ch := contributions.Get(args)
	for range args {
		c := <-ch
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%d,%d\n", c.Username, c.From, c.To, c.Total, c.CurrentStreak)
	}
}
