package main

import (
	"fmt"
	ArithmeticOp "go/ArithmeticOp"
)

func main() {

	var num1 float64
	var num2 float64

	fmt.Print("Enter num1 = ")
	fmt.Scanln(&num1)
	fmt.Print("Enter num2 = ")
	fmt.Scanln(&num2)
	fmt.Println("Addition is: ", ArithmeticOp.Sum(num1, num2))
	fmt.Println("Subtraction is: ", ArithmeticOp.Diff(num1, num2))
	fmt.Println("Multiplication is: ", ArithmeticOp.Multi(num1, num2))
	fmt.Println("Division is: ", ArithmeticOp.Div(num1, num2))
}
