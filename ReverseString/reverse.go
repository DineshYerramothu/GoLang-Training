package main

import (
	"fmt"
	"sort"
)

func main() {

	rev := []string{"The", "Lord", "Of", "The", "Rings"}
	fmt.Println("String is: ", rev)
	fmt.Println("Reverse of a string is: ", sort.Sort(sort.Reverse(sort.IntSlice(rev))))

}
