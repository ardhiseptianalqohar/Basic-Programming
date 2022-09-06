package main

import (
	"fmt"
	"math"
)

func calculateCircle(diameter float64) (float64, float64, string) {
	var luas = math.Pi * math.Pow(diameter/2, 2)
	var keliling = math.Pi * diameter
	return keliling, luas, "Sukses"

}
func perkalian(a, b int) (mul int) {
	mul = a * b
	return mul
}

func main()  {
	var data float64 = 10.0
	kel, ls, str := calculateCircle(data)
	fmt.Println("keliling :", kel)
	fmt.Println("luas :", ls)
	fmt.Println("str :", str)

	// var hasil = perkalian(10, 20)
	fmt.Println(perkalian(10, 20))
}