//Assignment04

package main

import "fmt"

func main() {

	// 1.
	// Create an ARRAY which holds 5 VALUES of TYPE int
	// ● assign VALUES to each index position.
	// ● Using format printing ○ print out the TYPE of the array
	var arr [5]int
	arr[0] = 2
	arr[1] = 3
	arr[2] = 4
	arr[3] = 5
	arr[4] = 6
	fmt.Println("\nAnswer 1")
	fmt.Println("Array is : ", arr)
	fmt.Printf("%T", arr)
	fmt.Print("\n \n")

	// 2.
	// Create a SLICE of TYPE int
	// ● assign 10 VALUES
	//  ● Using format printing ○ print out the TYPE of the slice
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Answer 2")
	fmt.Println("Slice is : ", slice)
	fmt.Printf("%T", slice)
	fmt.Print("\n \n")

	//3.
	// Using the code from the previous example,
	//       use SLICING to create the following new slices which are then printed:
	//   ● [42 43 44 45 46]
	//   ● [47 48 49 50 51]
	//   ● [44 45 46 47 48]
	//   ● [43 44 45 46 47]
	fmt.Println("Answer 3")
	range1 := slice[0:5]
	fmt.Println(range1)
	range2 := slice[5:10]
	fmt.Println(range2)
	range3 := slice[2:7]
	fmt.Println(range3)
	range4 := slice[1:6]
	fmt.Println(range4)
	fmt.Print("\n \n")

	// 4.
	// Start with this slice
	//   ○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//   . ● append to that slice this value
	//   ○ 52
	//   ● print out the slice
	//       in ONE STATEMENT append to that slice these values
	//  ○ 53
	//  ○ 54
	//  ○ 55
	//   ● print out the slice
	//   ● append to the x slice the below slice
	// ○ y := []int{56, 57, 58, 59, 60}
	//   ● print out the slice
	fmt.Println("Answer 4")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Print("\n \n")

	// 5.
	// 	  Write a program to pass a pointer to an array as an argument to the function
	//    a. Create an array of size 4, data type string
	//    b. Create a function with name updateThirdElement and update the value of 3rd index to Texas
	//    c. As an input to the function updateThirdElement pass pointer to an array to function updatearray
	arrState := [4]string{"Washington", "Alaska", "Ohio", "Kansas"}
	fmt.Println(arrState)
	updateThirdElement(&arrState)
	fmt.Println(arrState)

	// 6.
	// Create calculator app using pointers

	// num1 := 10
	// num2 := 20

	var num1 float64
	var num2 float64

	fmt.Println("Enter num1 value: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter num2 value: ")
	fmt.Scanln(&num2)

	fmt.Println("Addition of num1 and num2 is: ", add(&num1, &num2))
	fmt.Println("Subtraction of num1 and num2 is: ", sub(&num1, &num2))
	fmt.Println("Multiplication of num1 and num2 is: ", mul(&num1, &num2))
	fmt.Println("Division of num1 and num2 is: ", div(&num1, &num2))

}

func add(num1 *float64, num2 *float64) float64 {

	return *num1 + *num2
}

func sub(num1 *float64, num2 *float64) float64 {

	return *num1 - *num2
}

func mul(num1 *float64, num2 *float64) float64 {

	return *num1 * *num2
}

func div(num1 *float64, num2 *float64) float64 {

	return *num1 / *num2
}

func updateThirdElement(thirdIndex *[4]string) {

	(*thirdIndex)[3] = "Texas"
}
