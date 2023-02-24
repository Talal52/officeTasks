package main

import "fmt"

func main() {
	row := 10

	fmt.Println("\nLeft Angled Star Pattern")

	for i := 0; i <= row; i++ {

		for j := 0; j < i; j++ {
			fmt.Print("  ")
		}

		for k := 0; k < row-i; k++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
