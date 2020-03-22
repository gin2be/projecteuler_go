package main

import (
	"testing"
)

func TestP1Channel(t *testing.T) {
	patterns := []struct {
		param int
		want  int
	}{
		{10, 23},
		{1000, 233168},
	}

	for _, pattern := range patterns {
		param := pattern.param
		want := pattern.want
		got := P1Channel(param)

		if got != want {
			t.Errorf("P1Channel(%d) = %d, should be %d", param, got, want)
		}
	}
}

func TestP1Normal(t *testing.T) {
	patterns := []struct {
		param int
		want  int
	}{
		{10, 23},
		{1000, 233168},
	}

	for _, pattern := range patterns {
		param := pattern.param
		want := pattern.want
		got := P1Normal(param)

		if got != want {
			t.Errorf("P1Normal(%d) = %d, should be %d", param, got, want)
		}
	}
}
