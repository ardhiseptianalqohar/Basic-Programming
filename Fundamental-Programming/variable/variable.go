package main

import "fmt"

func main() {
	// var lastname string = "john"
	// firstname := "wick"

	// fmt.Printf("halo %s %s!\n", firstname, lastname)


	// deklarasi variabel tanpa menuliskan tipe data
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "sayhelo"
	fmt.Println(one, isFriday, twoPointTwo, say)

	// tipe data bool
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var nilai = (((2 + 6) % 3) * 4 - 2) / 3
	var isEqual = (nilai == 2)
	fmt.Printf("nilai %d (%t)", nilai, isEqual) // untuk memunculkan nilai bool menggunakan fmt.Prinf(), bisa juga gunakan layout (%t)

   // seleksi kondisi mengunakan keyword if, else if dan else
    point := 2
	if point == 10 {
		fmt.Println("nilai lulus")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda %d\n", point)
	}

	// variabel temporary pada if - else
	var angka = 10000.
	if ardi := angka / 100; ardi >= 100 {
		fmt.Printf("%.1f%s bagus !\n", ardi, "%")
	} else if ardi >= 70 {
		fmt.Printf("%.1f%s nice !\n", ardi, "%")
	} else {
		fmt.Printf("%.1f%s notbad !\n", ardi, "%")
	}

	// seleksi kondisi menggunakan key switch - case
	bebew := 6
	switch bebew {
	case 10:
		fmt.Println("Sangat bagus")
	case 7:
		fmt.Println("Lumayan")
	default:
		fmt.Println("tidak terlalu buruk") 
		// pada  kondisi ini atau case yang terpenuhi karena nilai var bebew tetap 6, 
	}   // ketika hal ini terjadi maka yg di panggil blok kondisi default


	// switch bisa digunakan dengan gaya if - else
	// ardhy := 6
	// switch {
	// case ardhy == 8:
	// 	fmt.Println("Sangat bagus")
	// case (ardhy < 8) && (ardhy > 3):
	// 	fmt.Println("Lumayan")
	// default:
	// 	{
	// 		fmt.Println("tidak terlalu buruk")
	// 	}
	// }


// penggunaan key falltrhough dalam switch, di gunakan untuk memaksa proses pengecekan ke satu case selanjutnya
	septian := 6
	switch {
	case septian == 8:
		fmt.Println("Sangat bagus")
	case (septian < 8) && (septian > 3):
		fmt.Println("Lumayan")
		fallthrough
	case septian < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("tidak terlalu buruk")
		}
	}

	// seleksi kondisi bersarang, yang berada dalam seleksi kondisi yg juga berada dalam seleksi kondisi dan seterusnya.
	//sleksi ini bisa di lakukan pada (if else, dan switch) atau kombinasi kedua nya
	Alqohar := 10 // input : 10, Output : 

	if Alqohar > 7 {
		switch Alqohar {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("nice")
		}
	}else {
		if Alqohar == 5 {
			fmt.Println("Not bad")
		} else if Alqohar == 3 {
			fmt.Println("keep trying")
		}else {
			fmt.Println("You can do it")
			if Alqohar == 0 {
				fmt.Println("try harde !")
			}
		}
	}


}