// cp437
package main

import (
	"fmt"

	"gbademo.go/demos/cp437/font"
)

func main() {
	for i := 0; i <= 255; i++ {
		fmt.Printf("ASCII code %v\n", i)
		for j := 0; j < 8; j++ {
			row := font.Char8x8[i][j]
			for k := 7; k >= 0; k-- {
				if row&(1<<k) > 0 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println("|")
		}
		fmt.Println("--------+")
	}
}
