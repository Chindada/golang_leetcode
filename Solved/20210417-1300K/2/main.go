package main

func Solution(S string) int {
	ans := 0
	count := 0
	for _, val := range S {
		if val == 97 {
			count++
		} else {
			count = 0
		}
		if count == 3 {
			return -1
		}
	}
	tmp := 0
	for _, val := range S {
		if val != 97 {
			ans += 2 - tmp
			tmp = 0
		} else {
			tmp++
		}
	}
	ans += 2 - tmp
	return ans
}
