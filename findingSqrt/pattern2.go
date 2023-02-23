package main

import "fmt"

func main() {
	fmt.Println("Enter length of triangle: ")
	var first int
	fmt.Scanln(&first)

	for i := 0; i < first; i++ {
		for j := 0; j < first-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
