package main

import "fmt"

func ardhy5(jam int) string{
	var temp string
	if jam < 20 {
		temp = "Selamat Pagi"
	} else if jam < 25 {
		temp = "Selamat Siang"
	}else if jam < 30 {
		temp = "Selamat Sore"
	} else if jam < 35 {
		temp = "Subuh"
	}else {
	    temp = "Mari sholat"
	}
	return temp
}

func main()  {
	fmt.Println(ardhy5(8))
	
}