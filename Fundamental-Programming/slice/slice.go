package main

import "fmt"

func main()  {
	var bilangan1 = []int {1,2,3,4,5} // menyatukan bilangan1 dan bilangan2 dengn (append)
	var bilangan2 = []int {6,7,8,9,10}
	bilangan1 = append(bilangan1, bilangan2...)
	fmt.Println(bilangan1)

	var prime  = [5]int{2, 3, 5, 7, 11}
	var part_primes []int =prime[1:4] //memotong dari sebuah array --> [3, 5, 7]
	fmt.Println(part_primes)
	

	bilangan1 = append(bilangan1, 10)	
	fmt.Println("====================")
	fmt.Println(bilangan1)
	fmt.Println(len(bilangan1))
	fmt.Println("====================")

	var nilai1  = []int{1, 2, 3}
	var nilai2  = []int{4, 5, 6}
	var nilai3 []int

	nilai3 = append(nilai1[1:], nilai2[1:]...) // menggabungkan 2 slice
	fmt.Println(nilai3)
	fmt.Println(nilai3[0])

}