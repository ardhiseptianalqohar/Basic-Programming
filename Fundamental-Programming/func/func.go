package main

import (
	"fmt"
)

func main()  {
	CetakHello()
	cetakApaKabar()
	var datasapa = data ()
	fmt.Println(datasapa)
	var tempint = returnint()
	fmt.Println(tempint) 
	fmt.Println(returnint())
	var angka = 100
	withParam(angka)
	
}


func CetakHello()  { // model func pascal case
	fmt.Println("hello")
	
}

func cetakApaKabar()  { //model func camel case dan ini type style (var snake_case string)
	fmt.Println("Apa kabar")
	
}

func data() string{
	var data string = "20 30 40 50 60 70 80 90"
	return data

	
}

func returnint () int{
	var i = 0 
	i = i + 20
	return i
}

func withParam(data1 int) {
	fmt.Println("data1" ,data1)
	var temp = data1 + 10
	fmt.Println(temp)


}