package main

import (
	"fmt"
)

func Sum(a, b int) int {
	return a + b
}

func main() {
	var aNum int
	var bNum int

	fmt.Println("Enter the 1st integer number:")
	fmt.Scan(&aNum)

	fmt.Println("Enter the 2nd integer number:")
	fmt.Scan(&bNum)

	fmt.Printf("The sum is: %v", Sum(aNum, bNum))

}
