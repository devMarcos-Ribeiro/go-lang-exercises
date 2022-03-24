// Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import "fmt"

const (
	x int = 42
	y     = 43
)

func main() {
	fmt.Printf("%d\t%d\n", x, y)
	fmt.Printf("%T\t%T", x, y)
}
