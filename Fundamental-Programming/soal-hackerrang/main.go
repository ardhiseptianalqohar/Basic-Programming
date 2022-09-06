package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
    // Write your code here
    for i :=int32(0); i < n;i++ {
        s := strings.Repeat(" ",int(n-i)-1) + strings.Repeat("#",int(i)+1)
        fmt.Println(s)
    }

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}


// package main

// import "fmt"


// func plusMinus(arr []int32) {
//     // Write your code here
//     ratioPositive, ratioNegative, ratioZero := 0.00, 0.00, 0.00
//     var length float64 = float64(len(arr))
//     for i := 0; i < len(arr); i++{
//         if arr[i] > 0 {
//             ratioPositive++
//         }
//         if arr[i] < 0 {
//             ratioNegative++
//         }
//         if arr[i] == 0 {
//             ratioZero++
//         }
//     }
//     ratioPositive = ratioPositive / length
//     ratioNegative = ratioNegative / length
//     ratioZero = ratioZero / length
//     fmt.Printf("%.6f \n", ratioPositive)
//     fmt.Printf("%.6f \n", ratioNegative)
//     fmt.Printf("%.6f \n", ratioZero)

// }

// func main()  {
// 	// fmt.Println(plusMinus([]))
// 	fmt.Println()
// }

// package main

// import (
//     "bufio"
//     "fmt"
//     "io"
//     "os"
//     "strconv"
//     "strings"
// )

// /*
//  * Complete the 'miniMaxSum' function below.
//  *
//  * The function accepts INTEGER_ARRAY arr as parameter.
//  */

// func miniMaxSum(arr []int32) {
//     // Write your code here
    
//     var min, max int32 = arr[0], arr[0]
//     var sum int64 = 0
    
//     for i :=0; i < len(arr); i++ {
//     if arr[i] > max {
//         max = arr[i]
//     } 
//     if arr[i] < min {
//         min = arr[i]
//     }  
//     sum += int64(arr[i]) 
//     }
//     fmt.Println(sum-int64(max), sum-int64(min))
// }

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

//     var arr []int32

//     for i := 0; i < 5; i++ {
//         arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
//         checkError(err)
//         arrItem := int32(arrItemTemp)
//         arr = append(arr, arrItem)
//     }

//     miniMaxSum(arr)
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }
