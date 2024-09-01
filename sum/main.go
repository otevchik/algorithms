package main

import "fmt"

// exercise 4.1

func sum(array []int) int{
	total := 0
	if (len(array) == 0) {
		return 0
	} else if (len(array) == 1) {
		return array[0]
	} else {
		total += array[0]
		array = append(array[1:])
		total += sum(array)
	}
	return total
}

func main() {
	array := []int{3, 1, 2}
	total := sum(array)
	fmt.Println(total)
}