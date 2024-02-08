package main

import "fmt"

func main() {
	arr := []int{3, 5, 1, 9, 2}
	if len(arr) == 0 {
		return
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println(max)
}
