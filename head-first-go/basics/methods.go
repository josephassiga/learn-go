package main

import (
	"fmt"
	"math"
)

// MyFloat is a variable
type MyFloat float64

// Abs is a function
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Vertex is a structure to define coordonates
type Vertex struct {
	X, Y float64
}

//Abs is a function to retrive thhe Absolute value of a Vertex.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(v.Abs())
}
