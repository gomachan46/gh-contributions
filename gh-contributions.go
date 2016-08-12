package contributions

import (
	"fmt"
	"strconv"

	"time"

	"github.com/PuerkitoBio/goquery"
)

func Get(username string) (string, error) {
	url := fmt.Sprintf("https://github.com/users/%s/contributions", username)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	var count, streak int
	start, _ := doc.Find("rect").First().Attr("data-date")
	doc.Find("rect").Each(func(_ int, s *goquery.Selection) {
		dc, _ := s.Attr("data-count")
		c, err := strconv.Atoi(dc)
		if err != nil {
			return
		}

		count += c
		if c == 0 {
			streak = 0
		} else {
			streak++
		}
	})
	return fmt.Sprintf("%s,%s,%d,%d\n", start, time.Now().Format("2006-01-02"), count, streak), nil
}
