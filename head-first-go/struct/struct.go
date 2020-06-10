package main

import "fmt"

type part struct {
	description string
	count int
}

type car struct {
	name string
	topSpeed float64
}

func main()  {
	var subscriber struct{
		name string
		rate float64
		active bool
	}
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true

	fmt.Println("Name:",subscriber.name)
	fmt.Println("Monthly rate:",subscriber.rate)
	fmt.Println("Active?",subscriber.active)

	var porshe car
	porshe.name = "Porshe 911 R"
	porshe.topSpeed = 323
	fmt.Println("Name:",porshe.name)
	fmt.Println("Top speed:",porshe.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:",bolts.description)
	fmt.Println("Count:",bolts.count)
}
