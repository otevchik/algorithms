package main

import "fmt"

func removeElement(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}

func findSmallest(array []int) int{
	smallest := array[0]
	smallest_index := 0
	for i := 1; i < len(array); i++ {
		if array[i] < smallest {
			smallest = array[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(array []int, newArray *[]int) {
	arr := make([]int, len(array))
	copy(arr, array)
	for len(arr) > 0 {
		smallest := findSmallest(arr)
		*newArray = append(*newArray, arr[smallest])
		arr = removeElement(arr, smallest)
	}
}

func main() {
	array := []int {4, 2, 1, 5, 3}
	newArray := make([]int, 0, len(array))
	selectionSort(array, &newArray)
	fmt.Println(array, newArray)
}