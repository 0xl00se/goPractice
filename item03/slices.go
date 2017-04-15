package main

import "fmt"

func main() {

	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p==", p[1:3])

	for i := 0;i < len(p); i++{
		fmt.Printf("p[%d]==%d\n", i, p[i])
	}
}
