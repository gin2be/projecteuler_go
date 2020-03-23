package main

import (
	"reflect"
	"testing"
)

func TestMakeInt(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := make([]int, 0)
	s := MakeInt(1, 10)
	for {
		v, ok := <-s
		if !ok {
			break
		}
		got = append(got, v)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MakeInt(1, 10) = %v, want: %v", got, want)
	}
}

func TestStreamFilter(t *testing.T) {
	want := []int{2, 4, 6, 8, 10}
	got := make([]int, 0)
	isEven := func(x int) bool {
		return x%2 == 0
	}
	s := StreamFilter(isEven, MakeInt(1, 10))
	for {
		v, ok := <-s
		if !ok {
			break
		}
		got = append(got, v)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("StreamFilter(isEven, MakeInt(1, 10)) = %v, want: %v", got, want)
	}
}

func TestMakeSieveStream(t *testing.T) {
	patterns := []struct {
		param int
		want  []int
	}{
		{0, []int{}},
		{1, []int{}},
		{2, []int{2}},
		{4, []int{2, 3}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, pattern := range patterns {
		param := pattern.param
		want := pattern.want
		got := make([]int, 0)
		s := MakeSieveStream(param)
		for {
			v, ok := <-s
			if !ok {
				break
			}
			got = append(got, v)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("MakePrimes(%v) = %v, want: %v", param, want, got)
		}
	}
}

func TestGetPrimeFactors(t *testing.T) {
	patterns := []struct {
		param int
		want  []int
	}{
		// primes
		{2, []int{2}},
		{3, []int{3}},
		{11, []int{11}},
		{13, []int{13}},
		{29, []int{29}},
		// numbers
		{8, []int{2, 2, 2}},
		{10, []int{2, 5}},
		{100, []int{2, 2, 5, 5}},
	}

	for _, pattern := range patterns {
		param := pattern.param
		want := pattern.want
		got := GetPrimeFactors(param)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetPrimeFactors(%v) = %v, want: %v", param, want, got)
		}
	}
}
