package main

import "fmt"

func fib() func() int{

	a, b := 0, 1
	i := 0
	return func() int {
		if i == 0 {
			i++
			return a
		} else if i == 1 {
			i++
			return b
		} else {
			i++
			a, b = b, a + b
			return b
		}


	}
}

func main() {

	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
