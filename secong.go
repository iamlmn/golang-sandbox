package main

import ("fmt"
"time"
"math/rand")

func foo(){
	// fmt.Println("The square root of 4 is",math.Square(4))
}
func random() {
	fmt.Println("A number from 1-100", rand.Intn(100))
}
func main() {
	//fmt.Println("The square root of 4 is",math.Sqrt(4))
	fmt.Println("It is currently:",time.Now())
	random()
}