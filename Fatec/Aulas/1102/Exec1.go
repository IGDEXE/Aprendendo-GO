// Escrever um algoritmo que receba um inteiro e some ele com o antecessor

package main

import "fmt"

func main() {
	var num, ant, res float64
	fmt.Printf("Informe um numero: ")
	fmt.Scanf("%f", &num)
	ant = num - 1
	res = ant + num
	fmt.Println("O numero é", res)
}
