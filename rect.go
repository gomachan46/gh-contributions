package grasshopper

type rect struct {
	date  string
	count string
	fill  string
}

// Date returns a date of rect
func (r *rect) Date() string {
	return r.date
}

// Count returns a contributions count of rect
func (r *rect) Count() string {
	return r.count
}

// Fill returns a color of rect
func (r *rect) Fill() string {
	return r.fill
}
