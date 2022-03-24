/*
	Using the following operators, write expressions and assign their values to variables:
	g. ==
	h. <=
	i. >=
	j. !=
	k. <
	l. >
	Now print each of the variables
*/

package main

import "fmt"

func main() {
	a := 42 == 42
	b := 42 <= 35
	c := 42 >= 2
	d := 42 != 42
	e := 42 < 42
	f := 42 > 41

	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v", a, b, c, d, e, f)
}
