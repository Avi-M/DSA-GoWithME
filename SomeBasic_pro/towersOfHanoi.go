package main

import "fmt"

func towerOfHanoi(n int, from_rod, to_rod, aux_rod string) {

	if n == 1 {
		fmt.Printf("\n Move disk 1 from rod %s to rod %s", from_rod, to_rod)
		return
	}
	towerOfHanoi(n-1, from_rod, aux_rod, to_rod)
	fmt.Printf("\n Move disk %d from rod %s to rod %s", n, from_rod, to_rod)
	towerOfHanoi(n-1, aux_rod, to_rod, from_rod)
}

func main() {
	n := 4
	towerOfHanoi(n, "A", "C", "B")
}

/*The pattern I followed :
Shift 'n-1' disks from 'A' to 'B'.
Shift last disk from 'A' to 'C'.
Shift 'n-1' disks from 'B' to 'C'.*/
