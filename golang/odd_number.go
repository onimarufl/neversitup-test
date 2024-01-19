package main

import "fmt"

func FindOddNumber(input []int) int {
	if len(input) == 1 {
		fmt.Println(fmt.Sprintf("should return %d, because it occurs %d time (which is odd).", input[0], len(input)))
		return len(input)
	}

	duplicate := make(map[int]int)
	for _, v := range input {
		_, exist := duplicate[v]
		if exist {
			duplicate[v] += 1
		} else {
			duplicate[v] = 1
		}
	}

	for k, v := range duplicate {
		if v%2 != 0 {
			fmt.Println(fmt.Sprintf("should return %d, because it occurs %d time (which is odd).", k, v))
			return v
		}
	}

	return 0
}
