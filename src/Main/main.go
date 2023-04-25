package main

import (
	"fmt"
	"main/sort"
)

func main() {
	arr := []int{6, 5, 3, 2, 1, 4}
	sorted := sort.Sleepsort(arr)
	fmt.Println("Sorted array is, ", sorted)
}
