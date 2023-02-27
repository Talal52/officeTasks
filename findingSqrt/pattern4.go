package main

import "fmt"

func main() {
	row := 5

	for i := 0; i < row; i++ {

		for j := 0; j < i; j++ {
			fmt.Print("  ")
		}

		for k := 0; k < 4-i; k++ {
			fmt.Print("* ")
		}
		for k := 0; k < 5-i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
