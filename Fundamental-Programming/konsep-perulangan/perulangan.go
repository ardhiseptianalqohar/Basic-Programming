package main

import "fmt"

func main() {
	/* Ada beberapa cara standar menggunakan (for)
	cara pertama dengan memasukan variabel counter perulangan beserta kondisinya
	setelah keyword
	*/


	/*angka 1 di cetak sebanyak 6 kali, perulangan ini hanya akan berjalan
    ketika variabel i bernilai dibawah angka (6), dngn ktentuan setiap kali
	perulangan.nilai variabel i akan di iterasi/ditambahkan satu,sama seperti
	i=i+1.karen i pada awalnya bernilai 0,maka perulangan akan berlangsung 5 kali
	yaitu bernilai (0,1,2,3,4,5)
*/
	// angka := 6
	// for i := 0; i < angka; i++ {
	// 	fmt.Println(1)
	// }


	/*cara ke 2 dngn menuliskan kondisi stelah keyword (for) hanya kondisi.deklarasi dan iterasi
	variabel counter tidak dituliskan stelah keywrd,hanya kondisi perulangannya saja.
	*/
	// i := 0
	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	/*penggunaan keywoard (for) tanpa argumen*/
	// cara ke 3 adalah (for) tanpa kondisi. dngn ini akan dihasilkan perulangan tanpa henti (sama dngn for true)
    // pemberhentian perulangan dilakukan dngn menggunakan key break
    // dalam perulangan tanpa henti dibawah, variabel (i) yg nilai awalnya 0 di inkrementasi.
	// ketika nilai (i) sudah mencapai 5, key break digunakan dan perulangan akan berhenti.
	
	// var i = 0

	// for {
	// 	fmt.Println("angka", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	/*Penggunaan key break & continue*/
	/*key break digunakan untuk mnghentikan scara paksa sebuah perulangan, sedangkan continue dipakai untuk memksa maju ke perulangan berikutnya*/

    /*perulangan mulai angka 1 hingga 10 dengan i sbgai variabel iterasi. Ketika i adalah gnjil(dapt dikthui i % 2, jika hasilnya 1, berarti ganjil)
	Maka akan dipaksa lanjut(continuee) ke perulangan berikutnya. Ketika i lebih besar dari 8, maka perulangan akan berhenti. Nilai akan ditampilkan*/
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}

	/*PERULANGAN BERSARANG*/
	/*Pada kode dibawah, untuk pertama kalinya fungsi -fmt.Print()- dipanggil tnpa disisipkan parameter. Cara sprti ini bisa digunakan untk menampilkan baris baru. 
	Kegunaanya sama seperti output dari statement -fmt.Print("\n")*/
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}


	/*PEMANFAATAN LABEL DALAM PERULANGAN*/
	/*di perulangan bersarang, break dn continue akan berlaku pada blok perulngan dimna ia digunkan saja.Ada cara agar kedua key ini bsa tertuju pada perulngan terluar
	atau perulngan trtentu. yaitu dngn memanfaatkan teknik pemberian label.*/

    /*tepat sebelum key -for- terluar, terdapat baris kode outerloop: .Maksud dari trsbut adlh disiapkan sebuh label bernama outerloop untk (for) dibawahnya. Nma label bisa digantik dngn nama lain
	(dan harus diakhiri dengn tanda titik dua("")/colon(:)). Pada (for) bagian dalam, terdapat seleksi kondisi untuk pengecekn nilai i. Ketika nilai trsbut sama dengan 3, 
	maka break dipanggil dengn target adalah perulangngn yg dilabeli outerloop, perulangan tersebut akan dihentikan*/
	outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Print("Matrik[", i, "][", j, "]", "\n")
		}
	}
}