package main

import "fmt"

func main() {
	fmt.Print("Enter length of triangle: ")
	var first int
	fmt.Scanln(&first)

	for i := 0; i < first; i++ {
		for j := 0; j < i+1; j++ {
			// for k := 0; k < j+1; k++ {
			// 	fmt.Print("")
			// }
			fmt.Print(" ")
		}
			for k :=0 ; k < first; k++ {
				fmt.Print("* ")
			}
			fmt.Println("")
	}
}
