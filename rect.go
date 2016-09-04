package grasshopper

// Rect is contribution info of a day
type Rect struct {
	date  string
	count string
	fill  string
}

// Date returns a date of rect
func (r *Rect) Date() string {
	return r.date
}

// Count returns a contributions count of rect
func (r *Rect) Count() string {
	return r.count
}

// Fill returns a color of rect
func (r *Rect) Fill() string {
	return r.fill
}
