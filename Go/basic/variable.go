package basic

import "strconv"

func Variable_func() string {

	var name string = "Maulana Yusuf" // cara 1, menggunakan var
	age := 20                         // cara 1, menggunakan simbol :=

	// multiple variable
	var a, b, c = "a", "b", "c"
	height, weight := 175, 70

	data := name + " " + strconv.Itoa(age) + a + b + c + strconv.Itoa(height) + strconv.Itoa(weight)

	return data
}

// constanta

const ThisMonth string = "Mei"
const (
	Animal1 string = "Gajah"
	Animal2 string = "Angsa"
)
