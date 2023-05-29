package main

import (
	"fmt"
	"sort"
)

// Input:
// 3 50
// 60 20
// 100 50
// 120 30
// Output:
// 180.000
func main() {
	m := make(map[float64]float64)
	var arr []float64
	var n, v float64
	fmt.Scan(&n, &v)
	for i := 0; i < int(n); i++ {
		var p, vol float64
		fmt.Scan(&p, &vol)
		m[p/vol] = vol
		arr = append(arr, float64(p/vol))
	}
	// fmt.Println(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	// fmt.Println(arr, m)
	i := 0
	var res float64
	for v > 0 && i < len(arr) {
		vol := m[arr[i]]
		if v >= vol {
			v -= vol
			res += float64(vol) * arr[i]
		} else {
			res += float64(v) * arr[i]
			v = 0
		}
		i++

	}
	fmt.Printf("%.3f", res)

}
