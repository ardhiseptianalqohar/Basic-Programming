// contoh soal dari hackkerrank

package main

import (
	"fmt"
)

func simpleArraySum(ar []int32) int32  {
	var temp int32
	for i := 0; i < len(ar); i++ {  // input : 6 dan (1, 2, 3, 4, 10, 11)
		fmt.Println(ar[i])
		temp += ar[i]
		
	}
	return temp
	

}

func main()  {
	fmt.Println()
}

