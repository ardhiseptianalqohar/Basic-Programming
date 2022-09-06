// contoh soal problem 4 - Total Maksimun dari deret bilangan (tanpa recursive)

package main

import "fmt"

func MaxSequence(arr []int) int {
	var max = -999999
	for i, _ := range arr {
        var sum int
		fmt.Println("arr", arr[i:])
		for _, value := range arr[i:]{
			sum += value
			//fmt.Println("sum", sum)
			
		}

	}
	return max
	
}

func main()  {
	fmt.Println(MaxSequence([]int {-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	
}