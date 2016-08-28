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
