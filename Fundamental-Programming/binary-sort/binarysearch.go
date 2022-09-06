// contoh binary search, atau mencari angka

package main

import "fmt"

func binarySearch(data []int, x int) int {
	var kiri = 0
	var kanan = len(data) - 1
	for kiri <= kanan{
		var tengah = (kiri + kanan) / 2
		if x < data[tengah]{
			kanan = tengah - 1
		} else if x > data[tengah] {
			kiri = tengah + 1
		}else if data[tengah] == x {
			return tengah
		}
	}
	return -1
}

func linierSearch(elements []int, x int) int {
	var count  int
	for i := 0; i < len(elements); i++ {
		count++
		fmt.Println("count", count)
		if elements[i] == x {
			return i
		}
	}
	return count 
}

func main() {
	data := []int{1, 4, 7, 9, 10, 14, 17, 20, 24}
	fmt.Println(binarySearch(data, 24))
	fmt.Println("========")
	fmt.Println(linierSearch(data, 24))

}