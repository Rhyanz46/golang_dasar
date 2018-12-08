package main

import (
	"fmt"
	"time"
)

func main() {
	defer finish()
	var barang = map[string]string{
		"arian": "coklat",
		"jihad": "cabe",
	}

	colorsMap := map[string]string{
		"AliceBlue":    "#F0F8FF",
		"AntiqueWhite": "#FAEBD7",
		"Aqua":         "#00FFFF",
		"Aquamarine":   "#7FFFD4",
		// NB: Go requires this trailing comma on the last element when initializing values over multiple lines
		"Azure": "#F0FFFF",
	}
	colorsSlice := colorsMap["AliceBlue"]

	fmt.Println("Inilah hasilnya : ", colorsSlice)
	// var a map[string]string = tampung(barang)
	// fmt.Println(a)
	var gg, dd = tampung(barang, "mantap jiwa", "mantap mantap", "mantap")
	fmt.Println(gg)
	fmt.Println(dd)
	fmt.Println(time.Since(time.Now()))

	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	type Android struct {
		Person string
		Model  string
	}

	go f(2)

	tt := Android{Person: "Perseonnn", Model: "Modelll"}
	fmt.Println(tt.Person)
}

func tampung(data map[string]string, nama ...string) (vv string, de string) {
	var isi = [2]string{
		"arian",
		"jihad",
	}
	for _, i := range isi {
		fmt.Println("nilai i = ", i, "nilai data = ", data[i])
	}
	vv = "mantapppppp"
	de = data[isi[0]]

	tambah := generators()
	fmt.Println(tambah())
	fmt.Println(tambah())
	fmt.Println(tambah())
	fmt.Println(tambah())
	fmt.Println(tambah())
	fmt.Println(tambah())

	fmt.Println(nama[2])
	return
}

func generators() func() int {
	i := 0
	return func() (jumlah int) {
		jumlah = factorial(i)
		i += 2
		return
	}
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func finish() {
	// panic("PANIC")
	str := recover()
	fmt.Println(str)
	fmt.Println("Finised ..")
}

func zero(xPtr *int) {
	*xPtr = 0
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
