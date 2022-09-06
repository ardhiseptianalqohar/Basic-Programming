package main  
import (  
 "fmt"  
)  
  
func CheckPrime(number int) {  
 isPrime := true  
 if number == 0 || number == 1 {  
  fmt.Printf(" %d is not a  prime number\n", number)  
 } else {  
  for i := 2; i <= number/2; i++ {  
   if number%i == 0 {  
    fmt.Printf(" %d is not a  prime number\n", number)  
    isPrime = false  
    break  
   }  
  }  
  if isPrime == true {  
   fmt.Printf(" %d is a prime number\n", number)  
  }  
 }  
}  
func main() {  
 CheckPrime(2)  // not a prime  
 CheckPrime(3)  // not a prime  
 CheckPrime(11)  // a prime  
 CheckPrime(13) // not a prime  
  
}  