package main

import (
	"fmt"
	"sort"
)

func minimizeBias(ratings []int32) int32 {
	var bias int32  
	sort.SliceStable(ratings, func(i, j int) bool {
		return ratings[i] < ratings[j]
	})
	return bias
}

func main()  {
	fmt.Println(minimizeBias([]int32{4,2, 5, 1}))
	
}