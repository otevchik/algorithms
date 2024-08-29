package main

import "fmt"

func binary_search(list []string, target string) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if (list[mid] == target) {
			return mid
		} else if (list[mid] > target) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	list := []string {"1", "2", "3", "4", "5"}
	target := "4"
	fmt.Println(binary_search(list, target))
}