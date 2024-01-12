package main

import "fmt"

type Laptop struct {
	model  string
	cpu    string
	price  float32
	memory uint32
}

func (lappy Laptop) PrintLaptop() {
	fmt.Printf("%s %s %f %d\n", lappy.model, lappy.cpu, lappy.price, lappy.memory)
}

func main() {
	var dell Laptop = Laptop{"inspiron", "i3", 600.0, 8}
	dell.PrintLaptop()
	var operation string
	var num1, num2 int
	fmt.Println("Calculator")
	fmt.Printf("============\n")
	fmt.Printf("\n operation : (add, sub, mul, div):")
	fmt.Scanf("%s\n", &operation)
	fmt.Printf("\nNumber 1: ")
	fmt.Scanf("%d\n", &num1)
	fmt.Printf("\nNumber 2: ")
	fmt.Scanf("%d\n", &num2)
	switch operation {
	case "add":
		fmt.Printf("\n %d", num1+num2)
	case "sub":
		fmt.Printf("\n %d", num1-num2)
	case "mul":
		fmt.Printf("\n %d", num1*num2)
	case "div":
		fmt.Printf("\n %d", num1/num2)
	default:
		fmt.Printf("\nInvalid operation")
	}
}
