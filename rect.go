package grasshopper

type rect struct {
	date  string
	count string
	fill  string
}

func (r *rect) Date() string {
	return r.date
}

func (r *rect) Count() string {
	return r.count
}

func (r *rect) Fill() string {
	return r.fill
}
