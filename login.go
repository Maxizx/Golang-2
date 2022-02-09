package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	login()
}

//Inicio login y rechazo del mismo
func login() {
	var user, pass string
	var usuario1 = [2]string{"pato", "patito"}
	var usuario2 = [2]string{"daniel", "danka"}
	var y string
	var x bool
	for i := x; i != true; {
		fmt.Println("quiere iniciar el programa? si/no")
		fmt.Scanln(&y)

		if y == "si" {
			fmt.Println("escriba su nombre de usuario")
			fmt.Scanln(&user)

			fmt.Println("escriba su contraseña")
			fmt.Scanln(&pass)

			if user == usuario1[0] || user == usuario2[0] {
				fmt.Println("__________________________")
				fmt.Println("usuario Existente         |")

				if pass == usuario1[1] || pass == usuario2[1] {
					fmt.Println("Contraseña correcta       |")
					fmt.Println("__________________________")
					acept()
					break

				}
				if pass != usuario1[1] || pass != usuario2[1] {
					fmt.Println("Contraseña incorrecta     |")
					fmt.Println("__________________________")

					b := []byte(user)

					err := ioutil.WriteFile("log", b, 0644)
					if err != nil {
						log.Fatal(err)
					}

					continue
				}
			}
			if user != usuario1[0] || user != usuario2[0] {
				fmt.Println("_________________")
				fmt.Println("usuario erroneo  |")
				fmt.Println("_________________")
				continue
			}
		}
		if y == "no" {
			break
		}
	}

}

func acept() {
	fmt.Println("----------------------------------")
	fmt.Println("Redirigiendo a la turnera")
	fmt.Println("----------------------------------")

	semana()
}

//Tuernera, elegir dia

func semana() {
	fmt.Println("elige un dia de la semana")
	fmt.Println("1. Lunes")
	fmt.Println("2. Martes")
	fmt.Println("3.Miercoles")
	fmt.Println("4. Jueves")
	fmt.Println("5. Viernes")
	fmt.Println("6. Sabado")
	fmt.Println("7. Domingo")

	var sem int
	fmt.Scanln(&sem)
	switch sem {
	case 1:
		fmt.Println("----------------------------------")
		fmt.Println("Lunes")

		hora()
	case 2:
		fmt.Println("----------------------------------")
		fmt.Println("Martes")

		hora()
	case 3:
		fmt.Println("----------------------------------")
		fmt.Println("Miercoles")

		hora()
	case 4:
		fmt.Println("----------------------------------")
		fmt.Println("Jueves")

		hora()
	case 5:
		fmt.Println("----------------------------------")
		fmt.Println("Viernes")

		hora()
	case 6:
		fmt.Println("----------------------------------")
		fmt.Println("Sabado")

		hora()
	case 7:
		fmt.Println("----------------------------------")
		fmt.Println("Domingo")

		hora()
	default:
		fmt.Println("Elige otro dia")

	}
}

//elegir turno
func hora() {

	fmt.Println("Elige un turno")
	fmt.Println("1. Mañana")
	fmt.Println("2. Tarde")
	fmt.Println("3. Noche")

	var turno int

	fmt.Scan(&turno)
	switch turno {
	case 1:
		fmt.Println("----------------------------------")
		fmt.Println("Mañana")

		Morning()
	case 2:
		fmt.Println("----------------------------------")
		fmt.Println("Tarde")
		after()
	case 3:
		fmt.Println("----------------------------------")
		fmt.Println("Noche")
		Night()
	default:
		fmt.Print("Sin turnos")
	}

}

//turnos y elegir hora
func Morning() {
	var mañana = [7]int{}
	var h int
	fmt.Println("Horarios disponibles a la mañana: ")
	for i, x := 6, 0; i < 13; i++ {

		fmt.Println(i)
		mañana[x] = i
		x++
	}
	for i := 0; i < 3; i++ {
		fmt.Scanln(&h)
		if i == 0 {
			continue
		}
		if h == mañana[0] || h == mañana[1] || h == mañana[2] || h == mañana[3] || h == mañana[4] || h == mañana[5] || h == mañana[6] {
			fmt.Printf("horario %v seleccionado", h)
			break
		}
		if h != mañana[0] || h != mañana[1] || h != mañana[2] || h != mañana[3] || h != mañana[4] || h != mañana[5] || h != mañana[6] {
			fmt.Println("horario inexistente")
			continue
		}

	}
}

func after() {
	var tarde = [7]int{}
	var h int
	fmt.Println("Horarios disponibles a la Tarde: ")
	for i, x := 14, 0; i < 21; i++ {

		fmt.Println(i)
		tarde[x] = i
		x++
	}
	for i := 0; i < 3; i++ {
		fmt.Scanln(&h)
		if i == 0 {
			continue
		}
		if h == tarde[0] || h == tarde[1] || h == tarde[2] || h == tarde[3] || h == tarde[4] || h == tarde[5] || h == tarde[6] {
			fmt.Printf("horario %v seleccionado", h)
			break
		}
		if h != tarde[0] || h != tarde[1] || h != tarde[2] || h != tarde[3] || h != tarde[4] || h != tarde[5] || h != tarde[6] {
			fmt.Println("horario inexistente")
			continue
		}

	}
}

func Night() {
	var h int
	var noche = [3]int{}

	fmt.Println("horarios disponibles")
	for i, x := 21, 0; i < 24; i++ {

		fmt.Println(i)
		noche[x] = i
		x++
	}

	for i := 0; i < 3; i++ {

		fmt.Scanln(&h)
		if i == 0 {
			continue
		}
		if h == noche[0] || h == noche[1] || h == noche[2] {
			fmt.Printf("horario %v seleccionado", h)
			break
		}
		if h != noche[0] || h != noche[1] || h != noche[2] {
			fmt.Println("horario inexistente")
			continue
		}

	}
}
