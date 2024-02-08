package main

import "fmt"

func main() {
	var input string
	// fmt.Scanln(&input)
	input = "radar"
	result := true
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			result = false
		}
	}

	fmt.Println(result)
}
