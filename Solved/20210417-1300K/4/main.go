package main

import (
	"fmt"
)

func main() {
	a := []int{13, 7, 2, 8, 3}
	b := Solution(a)
	fmt.Println(b)
}

func SolutionMy(A []int) int {
	for i := 2; i <= len(A); i++ {
		for j := 1; j <= len(A)-i+1; j++ {
			var tmp int
			for k := j - 1; k < i+j; k++ {
				if tmp == 0 {
					tmp = A[k]
					continue
				}
				if A[k]&tmp == 0 {
					return i - 1
				}
			}
		}
	}
	return 0
}

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	max := A[0]
	for _, val := range A {
		if val > max {
			max = val
		}
	}
	ans := 0
	tmp := 1
	for {
		tmpAns := 0
		for _, v := range A {
			if v&tmp > 0 {
				tmpAns++
			}
		}
		if tmpAns > ans {
			ans = tmpAns
		}
		if max/tmp == 0 {
			break
		}
		tmp *= 2
	}
	return ans
}
