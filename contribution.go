package gh_contributions

type Contribution struct {
	username      string
	from          string
	to            string
	total         int
	currentStreak int
	longestStreak int
}

func (c *Contribution) Username() string {
	return c.username
}

func (c *Contribution) From() string {
	return c.from
}

func (c *Contribution) To() string {
	return c.to
}

func (c *Contribution) Total() int {
	return c.total
}

func (c *Contribution) CurrentStreak() int {
	return c.currentStreak
}

func (c *Contribution) LongestStreak() int {
	return c.longestStreak
}
