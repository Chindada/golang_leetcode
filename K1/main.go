package main

import "fmt"

func main() {
	s := "a"
	fmt.Println([]byte(s))
	input := "dog"
	k1(input)

}

func k1(input string) int {
	var ans, tmp int
	for i, val := range input {
		if val == 97 {
			if tmp < 2 {
				tmp += 1
			} else {
				return -1
			}
		}
		if val != 97 {
			tmp = 0
			if i < len(input)-2 {
				if byte(input[i+1]) != 97 {
					ans += 1
				} else {
					continue
				}
			}
		}
	}
	fmt.Println(ans)
	return 0
}
