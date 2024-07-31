package modules

type Box struct {
	Panjang, Lebar, Tinggi float64
}

func (b *Box) Luas() float64 {
	return b.Panjang * b.Lebar
}

func (b *Box) Keliling() float64 {
	return 2 * (b.Panjang + b.Lebar)
}

func (b *Box) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}
