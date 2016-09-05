package grasshopper

// Rect is contribution info of a day
type Rect struct {
	date  string
	count string
	fill  string
}

// Date returns a date of Rect
func (r *Rect) Date() string {
	return r.date
}

// Count returns a contributions count of Rect
func (r *Rect) Count() string {
	return r.count
}

// Fill returns a color of Rect
func (r *Rect) Fill() string {
	return r.fill
}
