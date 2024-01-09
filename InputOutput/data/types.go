package data

import "fmt"

func PrintType() {
	var a int = 5
	var b float64 = 3.14
	// type conversion explicit
	a = int(b)
	fmt.Println("Type conversion ", a, b)
}
