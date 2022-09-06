package main

import "fmt"

func main() {
	data := []int{3, 2, 8, 1}
	for i := 0; i < len(data); i++{
		for j := i + 1; j < len(data); j++{
			fmt.Println("datai", data[i], "dataj", data[j])
			if data[i] > data[j]{
				fmt.Println("tukar")
			} else {
				fmt.Println("tetap")
			}
		}
}



}