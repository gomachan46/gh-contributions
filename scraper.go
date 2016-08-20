package gh_contributions

type Rect struct {
	date  string
	count int
}

type Scraper struct {
	username string
	rects    []*Rect
}
