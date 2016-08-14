package gh_contributions

import (
	"fmt"
	"strconv"

	"log"

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
		if c == 0 {
			contribution.CurrentStreak = 0
		} else {
			contribution.CurrentStreak++
		}
	}
	return contribution, nil
}

func Get(usernames []string) <-chan *Contribution {
	ch := make(chan *Contribution)
	for _, username := range usernames {
		go func(username string) {
			c, err := get(username)
			if err != nil {
				log.Fatalf("fail get contribution username: %s, err: %s", username, err.Error())
			}
			ch <- c
		}(username)
	}
	return ch
}
