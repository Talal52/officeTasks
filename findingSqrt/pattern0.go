package main

import "fmt"

func main() {
	row := 5

	for i := 0; i < row; i++ {

		for j := 0; j < row-i; j++ {
			fmt.Print("  ")
		}

		for k := 0; k <= i*2; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
