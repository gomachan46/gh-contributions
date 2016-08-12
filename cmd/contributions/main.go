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

	fmt.Fprint(os.Stdout, "username,start,end,total,currentStreak\n")
	ch := contributions.Get(os.Args[1:])
	for range os.Args[1:] {
			c := <-ch
			fmt.Fprintf(os.Stdout, "%s,%s,%s,%d,%d\n", c.Username, c.Start, c.End, c.Total, c.CurrentStreak)
	}
}
