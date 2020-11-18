package main

import "fmt"

func main() {
	arr := [3][4]int{{1, 1, 1, 1}, {1, 2, 3, 4}, {5, 6, 7, 8}}
	var sum int
	sum = 0
	for _, row := range arr {
		for _, val := range row {
			sum += val
		}
	}
	fmt.Print(" ", sum)
}
