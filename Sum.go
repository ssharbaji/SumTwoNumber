package main

import (
	"fmt"
)

func Sum(a, b int) int {
	return a + b
}

func main() {
	var a int
	var b int

	fmt.Println("Enter the 1st integer number:")
	fmt.Scan(&a)

	fmt.Println("Enter the 2nd integer number:")
	fmt.Scan(&b)

	fmt.Printf("The sum is: %v", Sum(a, b))

}
