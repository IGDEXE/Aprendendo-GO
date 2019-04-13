// Receber 3 notas e calcular a media

package main

import "fmt"

func main() {
	var nota1, nota2, nota3, media float64
	var noticia string
	fmt.Printf("Informe a primeira nota: ")
	fmt.Scanf("%f", &nota1)
	fmt.Printf("Informe a segunda nota: ")
	fmt.Scanf("%f", &nota2)
	fmt.Printf("Informe a terceira nota: ")
	fmt.Scanf("%f", &nota3)

	media = (nota1 + nota2 + nota3) / 3
	if media >= 6 {
		noticia = "aprovado"
	} else {
		noticia = "reprovado"
	}
	fmt.Println("O aluno foi", noticia, " com media", media)
}
