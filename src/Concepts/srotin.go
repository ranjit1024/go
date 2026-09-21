package main

import (
	"fmt"
	"slices"
)

func demo_sort() {
	roll_num := []int{3, 4, 5, 6, 7, 1}
	slices.Sort(roll_num)
	fmt.Println(roll_num)
	fmt.Println("Sroting....")
}
