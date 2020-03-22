package main

import (
	"testing"
)

func TestP3(t *testing.T) {
	patterns := []struct {
		param int
		want  int
	}{
		{4, 2},
		{20, 5},
		{30, 5},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, pattern := range patterns {
		param := pattern.param
		want := pattern.want
		got := P3(param)

		if got != want {
			t.Errorf("P3(%d) = %v, should be %v", param, got, want)
		}
	}
}
