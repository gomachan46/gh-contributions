package gh_contributions

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Rect struct {
	date  string
	count string
}

func (r *Rect) Date() string {
	return r.date
}

func (r *Rect) Count() string {
	return r.count
}

func scrape(username string) ([]*Rect, error) {
	url := fmt.Sprintf("https://github.com/users/%s/contributions", username)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	var rects []*Rect
	doc.Find("rect").Each(func(_ int, s *goquery.Selection) {
		d, _ := s.Attr("data-date")
		c, _ := s.Attr("data-count")
		rects = append(rects, &Rect{date: d, count: c})
	})

	return rects, nil
}
