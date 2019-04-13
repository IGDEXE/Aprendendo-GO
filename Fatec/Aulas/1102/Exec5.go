// Multiplos de 13

package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if i%13 == 0 {
			fmt.Println(i)
		}
		i++
	}
}
