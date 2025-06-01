package main

import (
	"fmt"
	"golang/basic"
)

func main() {

	fmt.Println("ini bagian main")
	basic.Basic_sintax()
	fmt.Println("-------------------------------------------------")

	// VARIABLE
	// fmt.Println(basic.Variable_func())
	fmt.Println("Hasil dari Belajar Variable")
	fmt.Printf("Bulan ini = %s\n", basic.ThisMonth)
	fmt.Printf("Hewan 1: %s, hewan 2: %s\n", basic.Animal1, basic.Animal2)
	fmt.Println("-------------------------------------------------")

	// TIPE DATA
	fmt.Println("Hasil Dari Belajar Tipe Data")
	fmt.Println(basic.Fruit)
	fmt.Println(basic.Angka1)
	fmt.Println(basic.Angka2)
	fmt.Printf("Nilai decimal = %.4f\n", basic.DecimalNumber) // %f (untuk format angka decimal)
	fmt.Printf("Hadir? %t\n", basic.Hadir)                    // %t (untuk format data boolean)
	fmt.Printf("Buah = %s\n", basic.Fruit)                    // %s (untuk format tipe data string)
	fmt.Println("-------------------------------------------------")

	// OPERATOR
	fmt.Println("Hasil Dari Belajar Operator")
	fmt.Println("----- operator aritmatika -----")
	fmt.Printf("Hasil Operasi Tambah: %d\n", basic.OperatorTambah)                      // %d (untuk format number)
	fmt.Printf("Hasil Operasi Kurang: %d\n", basic.OperatorKurang)                      // %d (untuk format number)
	fmt.Printf("Hasil Operasi Kali: %d\n", basic.OperatorKali)                          // %d (untuk format number)
	fmt.Printf("Hasil Operasi Bagi: %d\n", basic.OperatorBagi)                          // %d (untuk format number)
	fmt.Printf("Hasil Operasi Modulus: %d\n", basic.OperatorModulus)                    // %d (untuk format number)
	fmt.Printf("Hasil Operasi Perhitungan Sederhana: %d\n", basic.PerhitunganSederhana) // %d (untuk format number)
	fmt.Println("----- operator perbandingan -----")
	fmt.Printf("Hasil Operator Lebih Besar Dari: %t\n", basic.LebihBesarDari)
	fmt.Printf("Hasil Operator Lebih Kecil Dari: %t\n", basic.LebihKecilDari)
	fmt.Printf("Hasil Operator Lebih Besar Dari Sama Dengan: %t\n", basic.LebihBesarDariSamadengan)
	fmt.Printf("Hasil Operator Lebih Kecil Dari Sama Dengan: %t\n", basic.LebihKecilDariSamadengan)
	fmt.Printf("Hasil Operator Sama Dengan: %t\n", basic.Samadengan)
	fmt.Printf("Hasil Operator Tidak Sama Dengan: %t\n", basic.TidakSamadengan)
	fmt.Println("-------------------------------------------------")

}
