package modules

import (
	"math"
)

type DeretBilang struct {
	Value int
}

func (d *DeretBilang) Prima() []int {
	var primes []int
	for i := 2; i <= d.Value; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (d *DeretBilang) Ganjil() []int {
	var odds []int
	for i := 1; i <= d.Value; i += 2 {
		odds = append(odds, i)
	}
	return odds
}

func (d *DeretBilang) Genap() []int {
	var evens []int
	for i := 2; i <= d.Value; i += 2 {
		evens = append(evens, i)
	}
	return evens
}

func (d *DeretBilang) Fibonacci() []int {
	var fibs []int
	a, b := 0, 1
	for a <= d.Value {
		fibs = append(fibs, a)
		a, b = b, a+b
	}
	return fibs
}
