package gh_contributions

import (
	"fmt"
	"strconv"

	"os"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Contribution struct {
	Username      string
	From          string
	To            string
	Total         int
	CurrentStreak int
}

func getFrom(doc *goquery.Document) string {
	f, _ := doc.Find("rect").First().Attr("data-date")
	return f
}

func getTo(doc *goquery.Document) string {
	t, _ := doc.Find("rect").Last().Attr("data-date")
	return t
}

func getCounts(doc *goquery.Document) []string {
	return doc.Find("rect").Map(func(_ int, s *goquery.Selection) string {
		c, _ := s.Attr("data-count")
		return c
	})
}

func calcCurrentStreak(count int, before int) int {
	if count == 0 {
		return 0
	}
	return before + 1
}

func get(username string) (*Contribution, error) {
	contribution := &Contribution{Username: username}
	url := fmt.Sprintf("https://github.com/users/%s/contributions", contribution.Username)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return contribution, err
	}

	contribution.From = getFrom(doc)
	contribution.To = getTo(doc)
	counts := getCounts(doc)

	for _, cnt := range counts {
		c, err := strconv.Atoi(cnt)
		if err != nil {
			return contribution, err
		}
		contribution.Total += c
		contribution.CurrentStreak = calcCurrentStreak(c, contribution.CurrentStreak)
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
