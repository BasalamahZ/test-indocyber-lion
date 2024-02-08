package main

import "fmt"

func factorial(input int) int {
	if input== 0 {
		return 1
	}
	return input * factorial(input-1)
}

func main() {
	var input int
	// fmt.Scanln(&input)
	input = 5
	result := factorial(input)
	fmt.Println("result:", result)
}
