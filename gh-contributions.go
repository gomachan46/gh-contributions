package gh_contributions

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func calcCurrentStreak(count int, before int) int {
	if count == 0 {
		return 0
	}
	return before + 1
}

func calcLongestStreak(current int, before int) int {
	if current > before {
		return current
	}
	return before
}

func get(username string) (*Contribution, error) {
	rects, err := scrape(username)
	if err != nil {
		return nil, err
	}

	contribution := &Contribution{username: username}
	contribution.from = rects[0].date
	contribution.to = rects[len(rects)-1].date

	for _, r := range rects {
		c, err := strconv.Atoi(r.count)
		if err != nil {
			return contribution, err
		}
		contribution.total += c
		contribution.currentStreak = calcCurrentStreak(c, contribution.currentStreak)
		contribution.longestStreak = calcLongestStreak(contribution.currentStreak, contribution.longestStreak)
	}
	return contribution, nil
}

func Get(usernames []string) ([]*Contribution, []string) {
	var wg sync.WaitGroup
	var successes []*Contribution
	var failures []string
	for _, username := range usernames {
		wg.Add(1)
		go func(username string) {
			defer wg.Done()
			c, err := get(username)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fail get contribution username: %s, contribution: %v, err: %s\n", username, c, err.Error())
				failures = append(failures, username)
			} else {
				successes = append(successes, c)
			}
		}(username)
	}
	wg.Wait()
	return successes, failures
}
