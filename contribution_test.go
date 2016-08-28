package grasshopper

import "testing"

func TestContribution_Username(t *testing.T) {
	expected := "user"
	c := &Contribution{username: expected}
	if c.Username() != expected {
		t.Errorf("fail Username() got: %s, want: %s", c.Username(), expected)
	}
}
