package grasshopper

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func scrape(username string) ([]*rect, error) {
	url := fmt.Sprintf("https://github.com/users/%s/contributions", username)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	var rects []*rect
	doc.Find("rect").Each(func(_ int, s *goquery.Selection) {
		d, _ := s.Attr("data-date")
		c, _ := s.Attr("data-count")
		f, _ := s.Attr("fill")
		rects = append(rects, &rect{date: d, count: c, fill: f})
	})

	return rects, nil
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
		contribution.updateStreak(c)
	}
	return contribution, nil
}

// Get returns Contribution each username
// If Contribution are not created, username is stored in failures
// If Contribution are created, Contribution is stored in successes
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
