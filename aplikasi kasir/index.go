package main

import "fmt"

type Ucapan struct {
	selamat_datang  string
	selamat_tinggal string
}
dd := Ucapan{selamat_datang: "Welcome To My Home", selamat_tinggal: "Good bye.."}

func main() {
	var firstArray [10]int
	_ = dd
	_ = firstArray

	// When declaring an array literal, we can use the [...] syntax to have Go infer the length of the literal
	// For example, this array will have a length of 3
	anotherArray := [...]int{21, 42, -21}
	fmt.Println(anotherArray)
	fmt.Println(dd.selamat_datang)
}
