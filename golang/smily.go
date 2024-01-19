package main

import (
	"fmt"
	"strings"
)

func CountSmilyFace(text []string) int {
	fmt.Print(text)
	count := 0
	for _, str := range text {
		if strings.Contains(str, ")") || strings.Contains(str, "D") {
			count += 1
		}
	}
	fmt.Println(" / count : ", count)

	return count
}
