package main

import "fmt"

var array = []int{1, 2, 3, 5, 10, 22, 33, 55}

func bynary_search(arr []int, x int) bool {
	low := 0
	high := len(arr) - 1
	mid := (high - low) / 2

	for low != high {
		if arr[mid] == x {
			return true
		}

		if arr[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}

		mid /= 2
	}

	return false
}

func main() {
	fmt.Printf("bynary_search(array, 2): %v\n", bynary_search(array, 2))
}
