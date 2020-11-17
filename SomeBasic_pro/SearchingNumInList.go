package main

import "fmt"

func NumInList(lst []int, ele int) bool {
	for _, val := range lst {
		if val == ele {
			return true
		}
	}
	return false
}

func main() {
	lst := []int{34, 5, 6, 7, 78, 23}
	fmt.Println(NumInList(lst, 6))
}
