package grasshopper

import "testing"

func TestContribution_Username(t *testing.T) {
	expected := "user"
	c := &Contribution{username: expected}
	if c.Username() != expected {
		t.Errorf("fail Username() got: %s, want: %s", c.Username(), expected)
	}
}

func TestContribution_From(t *testing.T) {
	expected := "2015-08-26"
	c := &Contribution{from: expected}
	if c.From() != expected {
		t.Errorf("fail From() got: %s, want: %s", c.From(), expected)
	}
}

func TestContribution_To(t *testing.T) {
	expected := "2016-08-26"
	c := &Contribution{to: expected}
	if c.To() != expected {
		t.Errorf("fail To() got: %s, want: %s", c.To(), expected)
	}
}

func TestContribution_Total(t *testing.T) {
	expected := 100
	c := &Contribution{total: expected}
	if c.Total() != expected {
		t.Errorf("fail Total() got: %s, want: %s", c.Total(), expected)
	}
}

func TestContribution_CurrentStreak(t *testing.T) {
	expected := 10
	c := &Contribution{currentStreak: expected}
	if c.CurrentStreak() != expected {
		t.Errorf("fail CurrentStreak() got: %s, want: %s", c.CurrentStreak(), expected)
	}
}

func TestContribution_LongestStreak(t *testing.T) {
	expected := 50
	c := &Contribution{longestStreak: expected}
	if c.LongestStreak() != expected {
		t.Errorf("fail LongestStreak() got: %s, want: %s", c.LongestStreak(), expected)
	}
}
