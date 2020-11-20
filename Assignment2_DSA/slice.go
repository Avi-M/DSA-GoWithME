package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 3, 5, 6} // declare int slice
	slice = append(slice, 8)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	fmt.Println(slice[:3])

	slice2 := make([]int, 10) // define with capacity
	fmt.Println(cap(slice2))
	// define 2 dim slice
	twoDSlice1 := make([][]int, 2)
	twoDSlice1[0] = []int{1, 2, 3}
	twoDSlice1[1] = []int{4, 5, 6}
	fmt.Println()
	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice1))
	fmt.Printf("Number of columns in arsliceray: %d\n", len(twoDSlice1[0]))
	fmt.Printf("Total number of elements in slice: %d\n", len(twoDSlice1)*len(twoDSlice1[0]))
	fmt.Println("First Slice")
	for _, row := range twoDSlice1 {
		for _, val := range row {
			fmt.Print(" ", val)
		}
		fmt.Println()
	}
}
