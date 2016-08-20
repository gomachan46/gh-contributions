package gh_contributions

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
