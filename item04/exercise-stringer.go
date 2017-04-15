package main

import "fmt"

type IPAddr [4]byte

func main()  {
	addrs := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googledns": {8, 8, 8, 8},
	}
	for n, a := range addrs{
		fmt.Printf("%v: %v\n", n, a)
	}
}
