package grasshopper

import "testing"

func TestRect_Date(t *testing.T) {
	expected := "2016-09-07"
	r := &Rect{date: expected}
	if r.Date() != expected {
		t.Errorf("fail Date() got: %s, want: %s", r.Date(), expected)
	}
}

func TestRect_Count(t *testing.T) {
	expected := "10"
	r := &Rect{count: expected}
	if r.Count() != expected {
		t.Errorf("fail Count() got: %s, want: %s", r.Count(), expected)
	}
}

func TestRect_Fill(t *testing.T) {
	expected := "#111111"
	r := &Rect{fill: expected}
	if r.Fill() != expected {
		t.Errorf("fail Fill() got: %s, want: %s", r.Fill(), expected)
	}
}
