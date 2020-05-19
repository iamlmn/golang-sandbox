package main

import ("fmt")

func add(x float64, y float64) float64 {
	return x+y
}

func multiple(a,b string) (string,string) {
	return a,b
}

func main() {
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5

	// var num1, num2 float64 = 5.6,9.5

	num1,num2 := 5.6, 9.5 // we can actually not assign typing, and just define variables like we might in a dynamically-typed language, using := as the assignment operator.

	fmt.Println(add(num1,num2))

	w1,w2 := multiple("Hey","there")
	fmt.Println(w1,w2)
	
	var a int = 62
	var b float64 = float64(a)

	x := a // x will be a int

	// Type inference works too.
	var x float32
	y := x // y is float32 type
}