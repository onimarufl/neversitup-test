package main

import "fmt"

func Manipulate(text string) []string {
	resp := test2(text)
	fmt.Println(fmt.Sprintf("With input %s : return %s", text, resp))

	return resp
}

func test2(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	strSlice := []string{}
	for i, c := range str {
		subStr := str[:i] + str[i+1:]
		permutation := test2(subStr)

		for _, p := range permutation {
			strSlice = append(strSlice, string(c)+string(p))
		}
	}

	resp := removeDuplicateStr(strSlice)

	return resp
}

func removeDuplicateStr(strSlice []string) []string {
	keys := make(map[string]bool)
	slice := []string{}
	for _, item := range strSlice {
		if _, value := keys[item]; !value {
			keys[item] = true
			slice = append(slice, item)
		}
	}
	return slice
}
