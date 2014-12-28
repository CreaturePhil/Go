package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// include struct's name fields
	fmt.Printf("%+v\n", p)

	// Go syntax representation
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	// binary representation
	fmt.Printf("%b\n", 14)

	// character corresponding to integer
	fmt.Printf("%c\n", 33)

	// hex encoding
	fmt.Printf("%x\n", 456)

	// float
	fmt.Printf("%f\n", 78.9)

	// scientific notation
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	// double-quote strings as in Go source
	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "\"string\"")

	// pointer
	fmt.Printf("%p\n", &p)

	// padding left - right justify
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// padding right - left justify
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
