package main

import "fmt"

type Ninja struct {
	isAlive              bool
	health, power, speed int
}

func (n *Ninja) getAbilityLevel() int {
	return n.health * n.power * n.speed
}

func (n *Ninja) sacriface() {
	n.isAlive = false
}

func main() {
	kakashi := Ninja{
		isAlive: true,
		health:  10,
		power:   20,
		speed:   50,
	}
	fmt.Println(kakashi)
	fmt.Println(kakashi.getAbilityLevel())
	fmt.Println(kakashi.isAlive)
	kakashi.sacriface()
	fmt.Println(kakashi.isAlive)
}
