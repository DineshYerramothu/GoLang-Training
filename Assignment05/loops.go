package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 1. Using the for loop print 1 to 100 numbers
	fmt.Println("Answer 1")
	fmt.Println("Prining 1 to 100 numbers: ")
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	fmt.Print("\n \n")

	// 	2. Create a for loop using this syntax
	//      for condition { }
	// print out the odd numbers in 1 to 50.
	fmt.Println("Answer 2")
	fmt.Println("Odd numbers are:")
	i := 1
	for i <= 50 {
		i++
		if i%2 != 0 {
			fmt.Println(i)

		}
	}

	fmt.Print("\n \n")

	// 	3. Create a for loop using this syntax
	//     for { }
	// print out the even numbers in 1 to 50.
	fmt.Println("Answer 3")
	fmt.Println("Even numbers are:")
	j := 1
	for j <= 50 {
		j++
		if j%2 == 0 {
			fmt.Println(j)

		}
	}
	fmt.Print("\n \n")

	// 4. Print out the quotient for each number between 50 and 105 when it is divided by 6.
	fmt.Println("Answer 4")
	fmt.Println("Quotient for each number between 50 and 105 when it is divided by 6:")
	for i := 50.0; i <= 105.0; i++ {
		{
			var j float64
			j = i / 6.0
			fmt.Println("Quotient of ", i, "divided by 6 = ", j)
		}

	}
	fmt.Print("\n \n")

	// 	5. Read the user input.
	// If the input is equal to Golang tutorial print welcome else print end
	fmt.Println("Answer 5")
	fmt.Println("Enter text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println("Entered text is :", text)
	if text == "Golang tutorial" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}
	fmt.Print("\n \n")

	//  6. Write a program to print the numbers from 1 to 80.
	// But, for multiples of two print Golang instead of the number and for the multiples of four print tutorial.
	// For numbers which are multiples of both two and four print Golang tutorial.
	fmt.Println("Answer 6")
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println(i, "Golang Tutorial")
		} else if i%4 == 0 && i%2 == 0 {
			fmt.Println(i, "tutorial")
		} else if i%2 == 0 {
			fmt.Println(i, "Golang")
		}

	}

}
