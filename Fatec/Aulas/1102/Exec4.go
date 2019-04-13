// Peso ideal

package main

import "fmt"

func main() {
	var altura, peso float64
	var sexo string

	fmt.Printf("Informe sua altura: ")
	fmt.Scanf("%f", &altura)
	fmt.Printf("Informe M para masculino ou F para feminino: ")
	fmt.Scanf("%c", &sexo)

	if sexo == "M" {
		peso = (72.7 * altura) - 58
		fmt.Println("Seu peso ideal é", peso)
	} else if sexo == "F" {
		peso = (62.1 * altura) - 58
		fmt.Println("Seu peso ideal é", peso)
	} else {
		fmt.Println("Umas das informações não é valida")
	}
}
