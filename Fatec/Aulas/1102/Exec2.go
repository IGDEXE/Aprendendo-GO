// Calcula volume de um cilindro

package main

import "fmt"

func main() {
	var altura, volume float64
	const pi float64 = 3.14
	fmt.Printf("Informe a altura: ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Informe o raio: ")
	var raio float64
	fmt.Scanf("%f", &raio)

	volume = pi * (raio * raio) * altura
	fmt.Printf("O volume é igual a", volume)
}
