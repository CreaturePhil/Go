package main

import "fmt"

func newtonMethod(x float64) float64 {
	z, delta := 1.0, 1.0

	for delta > 0.0001 {
		z0 := z
		z = z0 - ((z0*z0 - x) / (2 * z0))
		if diff := z - z0; diff < 0 {
			delta = -diff
		} else {
			delta = diff
		}
	}

	return z
}

func main() {
	fmt.Println(newtonMethod(1))
	fmt.Println(newtonMethod(2))
	fmt.Println(newtonMethod(3))
	fmt.Println(newtonMethod(4))
	fmt.Println(newtonMethod(5))
}
