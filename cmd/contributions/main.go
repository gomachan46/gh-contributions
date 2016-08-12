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
	for _, arg := range os.Args[1:] {
		c, err := contributions.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "contribution: %v, err: %v", c, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%d,%d\n", c.Username, c.Start, c.End, c.Total, c.CurrentStreak)
	}
}
