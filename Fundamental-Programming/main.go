package main

import "fmt"

func main() {
	nilai := 80
	waktu := 4

	if nilai > 85 {
		fmt.Println("NIlai A")
		if waktu <= 4 {
			fmt.Println("Bagus")
		} else {
			fmt.Println("Standar")
		}
	}
	
	if nilai > 70 {
		fmt.Println("Baik")
	}
}