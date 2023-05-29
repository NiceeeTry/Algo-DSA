package main

import (
	"fmt"
	"sort"
)

// По данным n отрезкам необходимо найти множество точек минимального размера, для которого каждый из отрезков содержит хотя бы одну из точек.
// Input :
// 3
// 1 3
// 2 5
// 3 6
// Output:
// 1
// 3
// In:
// 4
// 4 7
// 1 3
// 2 5
// 5 6
// Out:
// 2
// 3 6
func main() {
	var n int
	fmt.Scan(&n)
	var arr [][]int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		arr = append(arr, []int{a, b})
	}
	// fmt.Println(arr)
	sort.Slice(arr, func(i, j int) bool {
		// Compare the first element of each sub-slice
		return arr[i][1] < arr[j][1]
	})
	// fmt.Println(arr)
	zero := arr[0][1]
	c := 0
	var res []int
	for i := 1; i < n; i++ {
		if arr[i][0] > zero {
			res = append(res, zero)
			zero = arr[i][1]
			c++
			continue
		}
		// else if i == n-1 {
		// 	res = append(res, zero)
		// 	c++
		// 	continue
		// }
	}
	c++
	res = append(res, zero)
	fmt.Println(c)
	for _, v := range res {
		fmt.Print(v, " ")
	}
	// fmt.Println(arr)

}
