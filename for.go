package main

import (
	"fmt"
)

func main() {
	var Numeros = [2]int{}

	fmt.Println("escribe tu numero")
	fmt.Scan(&Numeros[0], &Numeros[1])
	//base(Numeros)
	ejer4(Numeros)
}

func base(Numeros [2]int) {

	//var limit = 10

	for i := Numeros[0]; i < Numeros[1]; i++ {
		fmt.Println(i)
	}

}

func ejer3(Numeros [2]int) {

	for i := Numeros[0]; i < Numeros[1]; i += 3 {
		fmt.Println(i)

	}
}

func ejer4(Numeros [2]int) {
	var limit int
	limit = 2 * Numeros[0]
	for i := Numeros[0]; i < limit; i++ {
		fmt.Println(i)

	}
}
