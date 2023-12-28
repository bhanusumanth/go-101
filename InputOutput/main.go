package main

import "fmt"

func main() {
	print("Hello Input Output!Module!")
	var message = "Hello World"
	var str string = "Jack Reacher"
	// constants
	const pi = 3.14
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	println("\nInteger default:", i, i8, i16, i32, i64)
	var f float32
	var f64 float64
	println("Float default: ", f, f64)
	var b bool
	println("Boolean Default", b)
	var s string
	println("String Default:", s)
	println(message, str, pi)
	printData()
}

func printData() {
	fmt.Print("printing Data")
	// If the function is pubic, its first letter will be uppercase
	// If first letter is lower case, its private to the package
	// If we need to export this function to other package, then we need to make first letter uppercase
}
