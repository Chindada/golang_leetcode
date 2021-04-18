package main

func main() {
	S := "baaaaa"
	Solution(S)
}

func Solution(S string) int {
	ans := 0
	for i := 0; i < len(S); i++ {
		j := i
		for j < len(S) && S[j] == S[i] {
			j++
		}

		ans += (j - i) / 3
		i = j - 1
	}
	return ans
}
