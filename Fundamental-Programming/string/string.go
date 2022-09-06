package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main()  {
	//1. len string (menghitung huruf yg ada di sentence ada = 7)
	sentence := "alterra"
	lensentence := len(sentence)
	fmt.Println(lensentence)
	fmt.Println("=====================")

	//compare string
	star1 := "abc"
	star2 := "abd"
	fmt.Println(star1 == star2) // false
	fmt.Println("=====================")

	//contains
	var str = "something"
	var substr = "some"
	res := strings.Contains(str, substr) 
	fmt.Println(res) //true
	fmt.Println("=====================")

	//
	for i := 1; i <= 5; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print(" ") 
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ") ; 
		}
		fmt.Println()
	}

	fmt.Println("=====================")
	fmt.Println("INI ADALAH ARRAY")
	var data = [...]int{10, 20, 30, 40, 50,60, 70, 80, 90, 100} // ttik titik itu akan menghitung otomatis sendiri sistem nya, di saat kita tidak tahu isi banyak nya data
	fmt.Println(len(data))
    fmt.Println("=====================")

	var bilangan [10]int
	fmt.Println(reflect.ValueOf(bilangan).Kind())
	bilangan[0] = 10
	bilangan[2] = 15
	bilangan[len(bilangan)-1] = 30 // bilangan -1 = 30
	fmt.Println(bilangan)
	fmt.Println(len(bilangan)) // output nya 10
}