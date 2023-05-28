package main

import "fmt"

func interpolatedSearch(arr []int, target int) int {
	l, r := 0, len(arr)-1

	return (l + ((target-arr[l])/(arr[r]-arr[l]))*(r-l))
}

func main() {
	array := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	fmt.Println(interpolatedSearch(array, 14))
}
