package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v  years)", p.Name, p.Age)
}

func main()  {
	a := Person{"Anny", 15}
	z := Person{"Bob", 14}
	fmt.Println(a, z)
}