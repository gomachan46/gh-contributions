package grasshopper

// Contribution is GitHub contribution data struct
type Contribution struct {
	username      string
	from          string
	to            string
	total         int
	currentStreak int
	longestStreak int
	rects         []*rect
}

func (c *Contribution) updateStreak(count int) *Contribution {
	if count == 0 {
		c.currentStreak = 0
	} else {
		c.currentStreak++
	}
	if c.currentStreak > c.longestStreak {
		c.longestStreak = c.currentStreak
	}
	return c
}

// Username returns a username
// e.g. gomachan46
func (c *Contribution) Username() string {
	return c.username
}

// From returns a sum up started date
// e.g. 2015-08-26
func (c *Contribution) From() string {
	return c.from
}

// To returns a sum up ended date
// e.g. 2016-08-26
func (c *Contribution) To() string {
	return c.to
}

// Total returns a total contribution count during the period covered
func (c *Contribution) Total() int {
	return c.total
}

// CurrentStreak returns a current streak during the period covered
func (c *Contribution) CurrentStreak() int {
	return c.currentStreak
}

// LongestStreak returns a longest streak during the period covered
func (c *Contribution) LongestStreak() int {
	return c.longestStreak
}

func (c *Contribution) Rects() []*rect {
	return c.rects
}
