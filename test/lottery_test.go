package test

import (
	"testing"
)

func TestDraw(t *testing.T) {
	participants := []string{
		"Alice",
		"Bob",
		"Charlie",
		"David",
		"Eve",
	}

	prizes := []Prize{
		{"Prize1", 1},
		{"Prize2", 1},
		{"Prize3", 2},
	}

	winners, nonWinners, err := drawLots(participants, prizes)
	if err != nil {
		t.Logf("Error: %v\n", err)
		return
	}
	t.Logf("Winners: %v\n", winners)
	t.Logf("NoWinners: %v\n", nonWinners)
}
