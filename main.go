package main

import (
	"fmt"
	"khalifgfrz/tugasgolang1/modules"
)

func main() {
	// Nomor 1
	fmt.Println(modules.Pembulatan(4.37))
	fmt.Println(modules.Pembulatan(4.32))
	fmt.Println(modules.Pembulatan(4.324))

	// Nomor 2
	// deret := modules.DeretBilang{Value: 40}
	// fmt.Println("Bilangan Prima:", deret.Prima())
	// fmt.Println("Bilangan Ganjil:", deret.Ganjil())
	// fmt.Println("Bilangan Genap:", deret.Genap())
	// fmt.Println("Deret Fibonacci:", deret.Fibonacci())

	// Nomor 3
	// box := &modules.Box{Panjang: 3, Lebar: 4, Tinggi: 5}
    // fmt.Printf("Luas: %.2f\n", box.Luas())
    // fmt.Printf("Keliling: %.2f\n", box.Keliling())
    // fmt.Printf("Volume: %.2f\n", box.Volume())
}