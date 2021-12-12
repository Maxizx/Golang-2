package main

import (
	"fmt"
)

const NAME string = "Patito"

func main() {
	obtenerElMayor()
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
		//ejer_11(n1, n2, n3, n4)
	case 3:
		//ejer_12(n1, n2, n3, n4)
	case 4:
		//ejer_13(n1, n2, n3, n4)
	default:
		//fmt.Print("gracias")
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

func ejer_11(n1, n2, n3, n4 int) {

	//for i := 0; i < 10; i++ {
	//sum += i
	//}
	//var i int
	var num = [...]int{n1, n2, n3, n4}
	numM := num[0] //Asignar variable al primer elemento
	for _, numero := range num {
		if numero > numM {
			numM = numero
		}
	}

	fmt.Printf("El numero más grande es el %v", numM)

}

func ejer_12(n1, n2, n3, n4 int) {

	//var num = [...]int{n1, n2, n3, n4}
	//, n5, n6, n7, n8, n9, n0}

	//for i := 0; i < 10; i++ {
	//sum += i
	//}
	//var i int
	var num = [...]int{n1, n2, n3, n4}
	numM := num[0] //Asignar variable al primer elemento
	for _, numero := range num {
		if numero < numM {
			numM = numero
		}
	}
	fmt.Printf("El numero más chico es el %v", numM)

}

func ejer_13(n1, n2, n3, n4 int) {

	var num = [...]int{n1, n2, n3, n4}
	//fmt.Print(num)

	numI := num[0] //Asignar variable al primer elemento
	for _, numero := range num {
		if numero != numI {
			fmt.Print(numI)

		}
	}
	//fmt.Printf("El numero repeditdo es %v", numI)
}

func obtenerElMayor() {
	numeros := [5]int{0, 4, 7, 10, 2}
	mayor := 0

	for i := 0; i < len(numeros); i++ {
		if mayor < numeros[i] {
			mayor = numeros[i]
		}
	}

	fmt.Println(mayor)
}
