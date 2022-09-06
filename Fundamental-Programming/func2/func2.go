package main

import "fmt"

func ardhy(jam int) {
	if jam < 20 {
		fmt.Println("Selamat Pagi")
	} else if jam < 25 {
		fmt.Println("Selamat Siang")
	}else if jam < 30 {
		fmt.Println("Selamat Sore")
	} else if jam < 35 {
		fmt.Println("Subuh")
	} else {
		fmt.Println("Mari sholat")
	}
}

func ardhireturn(jam int) string {
		if jam < 20 {
			return "Selamat Pagi"
		} else if jam < 25 {
			return "Selamat Siang"
		}else if jam < 30 {
			return "Selamat Sore"
		} else if jam < 35 {
			 return "Subuh"
		} else {
			return "Mari sholat"
		}
	}

func main()  {
	ardhy (100)
	fmt.Println(ardhireturn(2))
	
}


