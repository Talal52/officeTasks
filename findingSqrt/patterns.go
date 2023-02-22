package main

import "fmt"

func main() {
	fmt.Println("Enter length of triangle Name: ")
	var first int
	fmt.Scanln(&first)

	for i := 0; i < first; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
