package main

import (
	"fmt"
	"math"
	"sort"
)

func minDiff(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr [j]
	})

	var result int32
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i] - arr[i+1]
		result += int32(math.Abs(float64(sum)))
	} 
	return result
}

func main()  {
	fmt.Println(minDiff([]int32{1, 3, 3, 2, 4})) // 3
	fmt.Println(minDiff([]int32{5, 1, 3, 7, 3})) // 6
}