package main

import (
	"fmt"
)

const NAME string = "Patito"

func main() {
	foor()
	//ejer_13()
	//ba()
}

func ba() {
	fmt.Println("Escribe tus numeros.")
	var numeros [10]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&numeros[i])
	}

	elegir(numeros)
}

//seleccion de operacion
func elegir(numeros [10]int) {

	var selec int
	fmt.Println("Elije")
	fmt.Println("1. ver tus numeros")
	fmt.Println("2. Ver el numero mayor")
	fmt.Println("3. Ver el numero menor")
	fmt.Println("4. Ver el numero Repetido")
	fmt.Scan(&selec)
	switch selec {
	case 1:
		ejer_10(numeros)
	case 2:
		ejer_11(numeros)
	case 3:
		ejer_12(numeros)
	case 4:
		//ejer_13(numeros)
	default:
		fmt.Print("gracias")
	}

}

func ejer_10(valores [10]int) {
	for i := 0; i < len(valores); i++ {
		fmt.Println(valores[i])
	}

	for _, num := range valores {
		fmt.Println(num)
	}
}

func ejer_11(valores [10]int) {

	//numeros := [5]int{0, 4, 7, 10, 2}
	mayor := 0

	for i := 0; i < len(valores); i++ {
		if mayor < valores[i] {
			mayor = valores[i]
		}
	}

	fmt.Println("El numero más grande es el: ", mayor)

}

func ejer_12(num [10]int) {

	menor := 0

	for i := 0; i < len(num); i++ {
		if menor < num[i] {
			menor = num[i]
		}
	}

	fmt.Println("El numero más grande es el: ", menor)

}

//func ejer_13(num [10]int) {
func ejer_13() {

	num := [4]int{3, 1, 2, 0}

	var igual, a int

	fmt.Scan(&a)

	for i := 0; i < len(num); i++ {
		if a == num[i] {
			igual = num[i]

		}
	}

	fmt.Println("El numero que se repite es el: ", igual)
}

func foor() {

	//var numeros [10]int

	/*for i := 0; i < 10; i++ {
		fmt.Scan(&numeros[i])

	}*/

	var num = [3]int{}
	var igual, a int
	//fmt.Scan(&num[0], &num[1], &num[2], &num[3], &num[4], &num[5])
	fmt.Scan(&num[0])
	for i := 1; i < 3; i++ {
		fmt.Scan(&a)

		if a != num[i] {
			num[i] = a

		}
		for x := 0; a == num[x]; x++ {
			if a == num[x] {
				igual = a
			}

		}

	}
	fmt.Println("se repite", igual)
	fmt.Println("su array es", num)
}
