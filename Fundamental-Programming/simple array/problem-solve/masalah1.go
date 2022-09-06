package main

import ("fmt"
)

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
    // Write your code here
    var nilai1 int32 = 0
    var nilai2 int32 = 0

    for i := 0; i < len(a); i++ {
        if a[i] > b[i] {
            nilai1++
        } else if a[i] < b[i] {
            nilai2++
        }
    }
    result := []int32{nilai1, nilai2}
    return result
}

func main() {
    data := []int32{5, 6, 7}
    fmt.Println(compareTriplets(data))
}