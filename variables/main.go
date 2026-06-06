package main

import "fmt"

func main() {
	fmt.Println("variables and types")

	var j int = 14
	var x float64 = -2.3
	var yo string = "Hi"
	var u uint = 14
	var symbol byte = 'H'
	var statement bool = true

	k := 34
	y := 3.7

	fmt.Println(j)
	fmt.Println(x)
	fmt.Println(yo)
	fmt.Println(u)
	fmt.Println(symbol)
	fmt.Println(statement)

	fmt.Println(float64(k) * y)
}
