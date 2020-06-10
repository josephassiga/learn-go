package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{Name: "Arthur Dent", Age: 45}
	z := Person{Name: "Zaphod Beeblebrox", Age: 9801}
	fmt.Println(a, z)
}
