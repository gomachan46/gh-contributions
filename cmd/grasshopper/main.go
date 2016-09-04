package main

import (
	"fmt"

	"os"

	"github.com/gomachan46/grasshopper"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "do nothing\n")
		os.Exit(1)
	}

	args := os.Args[1:]
	successes, failures := grasshopper.Get(args)
	fmt.Fprint(os.Stdout, "username,from,to,total,currentStreak,longestStreak\n")
	for _, c := range successes {
		for _, r := range c.Rects() {
			fmt.Fprintf(os.Stdout, "%s\n", r.Fill())
		}
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%d,%d,%d\n", c.Username(), c.From(), c.To(), c.Total(), c.CurrentStreak(), c.LongestStreak())
	}
	if len(failures) != 0 {
		fmt.Fprintf(os.Stderr, "fail get contributions. usernames: %v\n", failures)
	}
}
