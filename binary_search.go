package main

import "fmt"

//The slice or array HAS TO BE SORTED!!!
func binarySearch(array []int, target int) int {
	l, r := 0, len(array)-1

	for l <= r {
		m := (l + r) / 2
		if array[m] == target {
			return m
		}

		if array[m] > target {
			r--
		} else {
			l++
		}
	}
	return -1
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(items, 100))
}
