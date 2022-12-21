package main

import (
	"fmt"
	ArithmticOp "go/ArithmeticOp"
)

func main() {

	var num1 float64
	var num2 float64

	fmt.Print("Enter num1 = ")
	fmt.Scanln(&num1)
	fmt.Print("Enter num2 = ")
	fmt.Scanln(&num2)
	fmt.Println("Addition is: ", ArithmticOp.Sum(num1, num2))
	fmt.Println("Subtraction is: ", ArithmticOp.Diff(num1, num2))
	fmt.Println("Multiplication is: ", ArithmticOp.Multi(num1, num2))
	fmt.Println("Division is: ", ArithmticOp.Div(num1, num2))
}
