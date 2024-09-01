package main

import "fmt"

// exercise 4.1, 4.2, 4.3

func sum(array []int) (int, int, int){
	total := 0
	quantity := 0
	max := 0
	if (len(array) == 0) {
		return 0, 0, 0
	} else {
		if array[0] > max {
			max = array[0]
		}
		total += array[0]
		quantity += 1
		array = append(array[1:])
		a, b, c := sum(array)
		total += a
		quantity += b
		if c > max {
			max = c
		}
	}
	return total, quantity, max
}


func main() {
	array := []int{3, 1, 2, 4, 10}
	total, quantity, max := sum(array)
	fmt.Println(total, quantity, max)
}