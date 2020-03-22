package main

func P1Channel(param int) int {
	sum := 0

	ch := make(chan int)
	go func() {
		for i := 1; i < param; i++ {
			if i%3 == 0 || i%5 == 0 {
				ch <- i
			}
		}
		close(ch)
	}()

	for s := range ch {
		sum += s
	}
	return sum
}

func P1Normal(param int) int {
	sum := 0
	for i := 1; i < param; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
