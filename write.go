package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func write() {
	var y string
	fmt.Println("Escribe informacion")
	b2 := bufio.NewReader(os.Stdin)
	r, _ := b2.ReadString('\n')
	fmt.Println("Escribe el nombre del archivo")

	//fmt.Scanln(&r)
	fmt.Scanln(&y)
	b := []byte(r)
	err := ioutil.WriteFile(y, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Guardado con exito")

}

func write2() {
	b2 := bufio.NewReader(os.Stdin)
	r, _ := b2.ReadString('\n')

	fmt.Print(r)
}
