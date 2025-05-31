package main

import (
	"fmt"
	"golang/basic"
)

func main() {

	fmt.Println("ini bagian main")
	basic.Basic_sintax()

	// VARIABLE
	// fmt.Println(basic.Variable_func())
	fmt.Printf("Bulan ini = %s\n", basic.ThisMonth)
	fmt.Printf("Hewan 1: %s, hewan 2: %s\n", basic.Animal1, basic.Animal2)
	fmt.Println("-------------------------------------------------")

	// TIPE DATA
	fmt.Println(basic.Fruit)
	fmt.Println(basic.Angka1)
	fmt.Println(basic.Angka2)
	fmt.Printf("Nilai decimal = %.4f\n", basic.DecimalNumber) // %f (untuk format angka decimal)
	fmt.Printf("Hadir? %t\n", basic.Hadir)                    // %t (untuk format data boolean)
	fmt.Printf("Buah = %s\n", basic.Fruit)                    // %s (untuk format tipe data string)
	fmt.Println("-------------------------------------------------")

}
