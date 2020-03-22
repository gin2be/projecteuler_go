package main

func P2(param int) int {
	f1, f2 := 0, 1
	f3 := f1 + f2
	sum := 0
	for f3 <= param {
		if f3%2 == 0 {
			sum += f3
		}
		f1, f2, f3 = f2, f3, f2+f3
	}
	return sum
}
