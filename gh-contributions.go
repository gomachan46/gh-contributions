package contributions

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Contribution struct {
	Username      string
	Start         string
	End           string
	Total         int
	CurrentStreak int
}

func Get(username string) (*Contribution, error) {
	contribution := &Contribution{Username: username}
	url := fmt.Sprintf("https://github.com/users/%s/contributions", contribution.Username)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return contribution, err
	}

	s, _ := doc.Find("rect").First().Attr("data-date")
	contribution.Start = s
	e, _ := doc.Find("rect").Last().Attr("data-date")
	contribution.End = e
	counts := doc.Find("rect").Map(func(_ int, s *goquery.Selection) string {
		c, _ := s.Attr("data-count")
		return c
	})
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
