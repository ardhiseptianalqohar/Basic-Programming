package main

import ("fmt")

func main()  {
	// caesar chiper
	name := "abc"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i]+3)
	}
}


// ===========================
// func DrawXYZ(n int) string {
// 	var result string
// 	var temp int
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= n; j++ {
// 			temp++
// 			// fmt.Print(temp, " ")
// 			if temp%3 == 0 {
// 				fmt.Print(" X")
// 				// result = result + " X"
				
// 			} else if temp%2 == 1 {
// 				fmt.Print(" Y")
// 				// result = result + " Y"
// 			} else if temp%2 == 0 {
// 				fmt.Print(" Z")
// 				// result = result + " Z"
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// 	return result
// }

// func main()  {
// 	fmt.Println(DrawXYZ(3))
// 	fmt.Println(DrawXYZ(5))
//}