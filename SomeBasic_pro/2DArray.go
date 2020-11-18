package main

import "fmt"

func main() {
	samp := [2][3]int{{3, 2, 4}, {6, 3, 1}}
	fmt.Printf("total no of rows is %d\n ", len(samp))
	fmt.Println("total col", len(samp[0]))
	fmt.Println("total no of element", len(samp)*len(samp[0]))

	fmt.Println("Traversing Array:")
	fmt.Println("Using for-range")
	for _, row := range samp {
		for _, val := range row {
			fmt.Print(" ", val)
		}
		fmt.Println()
	}

	fmt.Println("\nUsing for loop")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(" ", samp[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\nUsing for loop - Second way")
	for i := 0; i < len(samp); i++ {
		for j := 0; j < len(samp[i]); j++ {
			fmt.Print(" ", samp[i][j])
		}
		fmt.Println()
	}
}
