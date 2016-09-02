package collate

import "testing"

func TestCollate(t *testing.T) {
	if len(SupportedLangs()) == 0 {
		t.Fatal("expected something greater than zero")
	}
	less := Index("en")
	if !less("a", "b") {
		t.Fatal("expected true, got false")
	}
}
