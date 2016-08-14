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

	args := os.Args[1:]
	successes, failures := contributions.Get(args)
	fmt.Fprint(os.Stdout, "username,from,to,total,currentStreak,longestStreak\n")
	for _, c := range successes {
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%d,%d,%d\n", c.Username(), c.From(), c.To(), c.Total(), c.CurrentStreak(), c.LongestStreak())
	}
	if len(failures) != 0 {
		fmt.Fprintf(os.Stderr, "fail get contributions. usernames: %v\n", failures)
	}
}
