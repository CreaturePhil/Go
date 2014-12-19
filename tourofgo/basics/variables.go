package main

import "fmt"

var c, python, java bool // package level

func main() {
	var i int // function level
	fmt.Println(i, c, python, java)
}
