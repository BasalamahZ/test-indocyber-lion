package main

import "fmt"

func main() {
	var i, j, input int
	// fmt.Scanln(&input)
	input = 5
	sym := "*"

	for i = 1; i <= input; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%s ", sym)
		}
		fmt.Println()
	}
}