// print the Character of a number based on ASCII Table

package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
