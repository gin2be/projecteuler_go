package main

func P3(param int) int {
	n := param

	// 偶数を奇数になるまで割る
	for n%2 == 0 {
		n /= 2
	}
	// 全部割り切れたら最大の素因数は2
	if n == 1 {
		return 2
	}

	answer := 3
	for {
		if n == 1 {
			return answer
		}
		if n%answer == 0 {
			n /= answer
		} else {
			answer += 2
		}
	}
	return answer
}
