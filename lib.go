package main

import "math"

type Stream chan int

func MakeInt(min, max int) Stream {
	s := make(Stream)
	go func() {
		for i := min; i <= max; i++ {
			s <- i
		}
		close(s)
	}()
	return s
}

func StreamFilter(fn func(int) bool, in Stream) Stream {
	s := make(Stream)
	go func() {
		for {
			v, ok := <-in
			if !ok {
				break
			}
			if fn(v) {
				s <- v
			}
		}
		close(s)
	}()
	return s
}

func MakeSieveStream(n int) Stream {
	s := make(Stream)
	filter := func(a int, in Stream) Stream {
		return StreamFilter(func(b int) bool { return b%a != 0 }, in)
	}
	go func() {
		in := MakeInt(2, n)
		for {
			x, ok := <-in
			if !ok {
				break
			}
			s <- x
			if x*x <= n {
				in = filter(x, in)
			}
		}
		close(s)
	}()
	return s
}

func GetPrimeFactors(n int) []int {
	factors := make([]int, 0)
	limit := int(math.Sqrt(float64(n)))

	s := MakeSieveStream(limit)
	for {
		p, ok := <-s
		if !ok || p*p > n {
			break
		}
		for {
			if n%p != 0 {
				break
			}
			factors = append(factors, p)
			n /= p
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}
