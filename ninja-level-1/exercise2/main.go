package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Print(x, y, z)
	// The compiler assigned ZERO Values to the variables, since they haven't been initialized.
}
