package main

import (
	"testing"
)

func TestP2(t *testing.T) {
	patterns := []struct {
		param int
		want  int
	}{
		{10, 10},
		{50, 44},
		{4000000, 4613732},
	}

	for _, pattern := range patterns {
		param := pattern.param
		want := pattern.want
		got := P2(param)

		if got != want {
			t.Errorf("P2(%d) = %d, should be %d", param, got, want)
		}
	}
}
