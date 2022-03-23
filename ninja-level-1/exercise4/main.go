package main

import "fmt"

type Year int

func main() {
	var x Year
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 1942
	fmt.Printf("%v\n", x)
}
