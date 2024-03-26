package main

import (
	"fmt"
	"strconv"
)

func main() {

	var totalNotas int
	var somaNotas float64

	fmt.Println("Calculadora de Médias")
	fmt.Print("Quantas notas você deseja calcular?: ")

	fmt.Scan(&totalNotas)

	if totalNotas <= 0 {
		fmt.Println("Numero de notas invalido, escolha um valor válido")
		return
	}

	for i := 1; i <= totalNotas; i++ {
		var inputNota string
		fmt.Println("Digite a nota", i, ":")
		fmt.Scan(&inputNota)

		nota, err := strconv.ParseFloat(inputNota, 64)
		if err != nil {
			fmt.Println("Nota invalida. Por favor, digite um número válido.")
			i--
			continue
		}
		somaNotas += nota

	}
	media := somaNotas / float64(totalNotas)

	fmt.Printf("A média das notas é: %.2f", media)
}
