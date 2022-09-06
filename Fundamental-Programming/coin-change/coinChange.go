package main

import (
	"fmt"
	"sort"
)

func coinChange(money int, coinvalue []int) []int {
	sort.SliceStable(coinvalue, func(i, j int) bool {
		return coinvalue[i] > coinvalue[j]
	})
	result := []int{}
	for _, coin := range coinvalue {
		if money >= coin {
			for money >= coin {
				result = append(result, coin)
				money = money - coin
			}	
		} else {
			continue
		}
	}
	return result

}


/* koin : 10, 25, 5, 1
uang : 42
42 = 1..... seterusnya sampai (42) kali dengan hasil yg sama dengan value = (42)
42 = 5, 5, 5, 5, 5, 5, 5, 5, 1, 1 dengan (10) keping koin yg di dapat dengan value yg sama yaitu = (42)
42 = 10,10,10,10,1,1 dengan (6) kali keping koin = 42
42 = 10,10,10,5,5,1,1 dengan (7) keping koin = 42
*/
func main() {
	coins := []int{10, 25, 5, 1}
	fmt.Println(coinChange(42, coins))
}