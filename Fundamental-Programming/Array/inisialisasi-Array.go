package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*INISIALISASI NILAI AWAL ARRAY*/
	/*Pengisian elemen array bisa dilakukan pada saat deklarasi variabel.Caranya dgn mnuliskn data elemen dlm kurung kurawal stlah tipe data
	dgn pembatas antar elemen adalah tanda (,) koma. Pengunaan fungsi fmt.Println() pada data array tanpa mengakses indeks trtentu, 
	akan mnghsilkan output dlm bntuk string
	dari semua array yg ada.Teknik ini biasa digunkan debugging data array. Fungsi len() dipakai untk mnghtung jmlh elemen sebuah array*/
	
	var fruits = [4]string{"apel", "anggur", "mangga", "melon"}
	
	fmt.Println("jumnlah elemen \t\t", len(fruits))
	fmt.Println("jumnlah elemen \t\t", fruits)



	/*INISIALISASI NILAI AWAL ARRAY TANPA JUMLAH ELEMEN*/
	/*deklarasi array yg nilainya diset awal, boleh tidak di tliskan jmlh lebar arraynya,cukup ganti dgn tanda 3 titik( ... )
	jmlah elemen akan di kalkulasi scara otomatis mnysuaikan data elemen yg di isikan.
	variabel numbers akan secara otomatis memiliki jmlh elemen 5, krna pada saat deklarasi disiapkan 5 buah elemen.*/
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("data array \t:", len(numbers))



	/*Array Multidimensi*/
	/*Array multidimensi adalah array yang tiap elemennya juga berupa array (dan bisa seterusnya, tergantung ke dalaman dimensinya).
    Cara deklarasi array multidimensi secara umum sama dengan cara deklarasi array biasa, dengan cara menuliskan data array dimensi
    selanjutnya sebagai elemen array dimensi sebelumnya.
    Khusus untuk array yang merupakan sub dimensi atau elemen, boleh tidak dituliskan jumlah datanya.
    Kedua array di bawah memiliki elemen yang sama.*/
	//  var numbers1  = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	//  var numbers2  = [2][3]int {{3, 2, 3}, {3, 4, 5}} 

	//  fmt.Println("numbers1", numbers1)
	//  fmt.Println("numbers2", numbers2)



	/*Perulangan Elemen Array Menggunakan Keyword for*/
	/*Keyword for dan array memiliki hubungan yang sangat erat.
	Dengan memanfaatkan perulangan menggunakan keyword ini, elemen elemen dalam array bisa didapat.
    Ada beberapa cara yang bisa digunakan untuk me-looping data array, yg pertama adalah dengan memanfaatkan variabel iterasi
    perulangan untuk mengakses elemen berdasarkan indeks-nya.

	Perulangan di bawah dijalankan sebanyak jumlah elemen array fruits (bisa diketahui dari kondisi i < len(fruits ). Di tiap
    perulangan, elemen array diakses lewat variabel iterasi i .*/
    var fruite  = [4]string{"apel", "mangga", "pisang", "melon"}
	for i := 0; i < len(fruite); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruite[i])
	}



	/*Perulangan Elemen Array Menggunakan Keyword for - range*/
    /*Ada cara yang lebih sederhana me-looping data array, 
	dengan menggunakan keyword for - range.Array fruits diambil elemen-nya secara berurutan. 
	Nilai tiap elemen ditampung variabel oleh fruit (tanpa huruf s), sedangkan
    indeks nya ditampung variabel i.
	Output program di bawah, sama dengan output program sebelumnya, hanya cara yang digunakan berbeda.*/
	var buah  = [4]string{"apel", "mangga", "pisang", "melon"}

	for i, buah := range buah {
		fmt.Printf("elemen %d : %s\n", i, buah)
	}



    /*Penggunaan Variabel Underscore _ Dalam for - range*/
	/*Kadang kala ketika looping menggunakan for - range , ada kemungkinan di mana data yang dibutuhkan adalah elemen-nya saja,
    indeks-nya tidak. Sedangkan kode di atas, range mengembalikan 2 data, yaitu indeks dan elemen.
    Seperti yang sudah diketahui, bahwa di Go tidak memperbolehkan adanya variabel yang menganggur atau tidak dipakai. Jika
    dipaksakan, error akan muncul
    */

    // var buah1 = [4]string{"apple", "grape", "banana", "melon"}
    // for i, fruit := range buah1 { // di variabel (i) menganggur dan tidak terpakai
    // fmt.Printf("nama buah : %s\n", fruit)
    // }

    /*  Di sinilah salah satu kegunaan variabel pengangguran, atau underscore ( _ ). Tampung saja nilai yang tidak ingin digunakan ke
    underscore.
   */
var buah2 = [4]string{"apple", "grape", "banana", "melon"}
for _, fruit := range buah2 { // for (i) nya di ganti dengan tanda underscore
fmt.Printf("nama buah : %s\n", fruit)
}



    /*Alokasi Elemen Array Menggunakan Keyword make*/
    /*Deklarasi sekaligus alokasi data array juga bisa dilakukan lewat keyword make .
    Parameter pertama keyword make diisi dengan tipe data elemen array yang diinginkan, parameter kedua adalah jumlah elemennya.
    Pada kode di atas, variabel fruits tercetak sebagai array string dengan alokasi 2 slot.
    */

var buah3  = make([]string, 2)
buah3[0] = "anggur"
buah3[1] = "semangka"

fmt.Println(buah3)


var bilangan [4] int // array zero value nya yaitu = 0
fmt.Println(reflect.ValueOf(bilangan).Kind())
bilangan[0] = 10
//bilangan[1] = 20
bilangan[2] = 15
//bilangan[3] = 14
fmt.Println(bilangan)
fmt.Println(len(bilangan)) // len itu untuk menghitung panjang/size array

var country [3] string 
country[0] = "indonesia"
country[1] = "malaysia"
country[2] = "singapura"
fmt.Println(len(country)) 


const Index = 0
var data [3] int = [3] int{Index, 20, 30}
fmt.Println(data)

fmt.Println("=====================")
	fmt.Println("INI ADALAH ARRAY")
	var coba = [...]int{10, 20, 30, 40, 50,60, 70, 80, 90, 100} // ttik titik itu akan menghitung otomatis sendiri sistem nya, di saat kita tidak tahu isi banyak nya data
	fmt.Println(len(coba))
    fmt.Println("=====================")

	var bilangan5 [10]int
	fmt.Println(reflect.ValueOf(bilangan5).Kind())
	bilangan[0] = 10
	bilangan[2] = 15
	bilangan[len(bilangan)-1] = 30 // bilangan -1 = 30
	fmt.Println(bilangan)
	fmt.Println(len(bilangan)) // output nya 10

}