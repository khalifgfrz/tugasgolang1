package main

import (
	"fmt"
	"khalifgfrz/tugasgolang1/modules"
)

func main() {
	fmt.Println(modules.Pembulatan(4.37))
	fmt.Println(modules.Pembulatan(4.32))
	fmt.Println(modules.Pembulatan(4.324))

	deret := modules.DeretBilang{Value: 40}

	fmt.Println("Bilangan Prima:", deret.Prima())
	fmt.Println("Bilangan Ganjil:", deret.Ganjil())
	fmt.Println("Bilangan Genap:", deret.Genap())
	fmt.Println("Deret Fibonacci:", deret.Fibonacci())

	box := &modules.Box{Panjang: 3, Lebar: 4, Tinggi: 5}

    fmt.Printf("Luas: %f\n", box.Luas())
    fmt.Printf("Keliling: %f\n", box.Keliling())
    fmt.Printf("Volume: %f\n", box.Volume())
}