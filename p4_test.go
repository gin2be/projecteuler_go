package main

import (
	"testing"
)

func TestP4(t *testing.T) {
	patterns := []struct {
		param int
		want  int
	}{
		{2, 9009},
		{3, 906609},
	}

	for _, pattern := range patterns {
		param := pattern.param
		want := pattern.want
		got := P4(param)

		if got != want {
			t.Errorf("P4(%d) = %v, should be %v", param, got, want)
		}
	}
}
