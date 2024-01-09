package main

import "fmt"

func main() {
	square, cube := calculateSquareCube(14)
	fmt.Println(square, cube)
	n := 0
	pointerDemo(&n)
}

func pointerDemo(ptr *int) {
	*ptr = *ptr + 100
	switch *ptr {
	case 100:
		fmt.Printf("100")
		fallthrough
	case 200:
		fmt.Printf("200")
		fallthrough
	case 300:
		fmt.Printf("300")
		fallthrough
	default:
		fmt.Printf("Default")
	}
}

func calculateSquareCube(n int) (int, int) {
	if n != 100 {

	}
	return n * n, n * n * n
}
