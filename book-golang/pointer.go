package main

import "fmt"

func zeroValue(x int) {
	x = 0
}

func zeroReference(xPtr *int) {
	*xPtr = 0
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x := 5
	zeroValue(5)
	fmt.Println("x by value:", x) // x is still 5

	zeroReference(&x)
	fmt.Println("x by reference:", x) // x is 0

	y := 2
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	swap(&x, &y)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
