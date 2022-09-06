package main

import (
	"fmt"
	"strings"
)

func main() {
	// var data int = 5
	// for i := 0; i < data; i++ {
	// 	if data == 25 {
	// 		fmt.Println("selamat sore")
	// 	} else if data > 50 {
	// 		fmt.Println("selamat malam")
	// 	}else{
	// 		fmt.Println("selamat pagi")
	// 	}
	// } 

    //  var data int = 1
	// for i := 1; i <= 5; i++ {
	// 	fmt.Printf("%d",data)
	// }

	/*Fungsi strings.Contains()
	Dipakai untuk deteksi apakah string (parameter kedua)
	 merupakan bagian dari string lain (parameter pertama). Nilai kembaliannya berupa bool.
	*/
		var isExists = strings.Contains("john wick", "wick")
		fmt.Println(isExists)


		/*Fungsi strings.Count()
		Memiliki kegunaan untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama).
		 Nilai kembalian fungsi ini adalah jumlah karakternya.
		 Nilai yang dikembalikan 2, karena pada string "ethan hunt" terdapat dua buah karakter "t".
         */
		
		var howMany = strings.Count("ethan hunt", "u")
		fmt.Println(howMany) // 2


		/*Fungsi strings.Index()
Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).*/

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2
/*String "ha" berada pada posisi ke 2 dalam string "ethan hunt" (indeks dimulai dari 0).
 Jika diketemukan dua substring, maka yang diambil adalah yang pertama, contoh:*/

var index2 = strings.Index("ethan hunt", "n")
fmt.Println(index2) // 4
/*String "n" berada pada indeks 4 dan 8. Yang dikembalikan adalah yang paling kiri (paling kecil), yaitu 4.*/
		


	}

