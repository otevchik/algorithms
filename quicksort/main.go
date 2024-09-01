package main

import "fmt"

func quicksort(array []int) []int{
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		less := make([]int, 0, 10)
		for i := 1; i < len(array); i++ {
			if array[i] <= pivot {
				less = append(less, array[i])
			}
		}

		greater := make([]int, 0, 10)
		for i := 1; i < len(array); i++ {
			if array[i] > pivot {
				greater = append(greater, array[i])
			}
		}
		array1 := make([]int, 0, len(array))
		array1 = quicksort(less)
		array1 = append(append(array1, pivot), quicksort(greater)...)
		return array1
	}
}

func main() {
	array := []int{10, 5, 2, 3, 0, 14}
	array = quicksort(array)
	fmt.Println(array)
}