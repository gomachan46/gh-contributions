package grasshopper

import "testing"

func TestRect_Date(t *testing.T) {
	expected := "2016-09-07"
	r := &Rect{date: expected}
	if r.Date() != expected {
		t.Errorf("fail Date() got: %s, want: %s", r.Date(), expected)
	}
}
