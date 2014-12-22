package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(find_smallest_number(x))
}

func find_smallest_number(list []int) int {
	var smallest int = list[0]
	for i, v := range list {
		if v < smallest {
			smallest = list[i]
		}
	}
	return smallest
}
