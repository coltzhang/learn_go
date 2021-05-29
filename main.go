package main

import "fmt"

var a int
var (
	b string
	c int
)

const (
	aa = 1
	bb = "bb"
)

func main() {
	a = 10
	b = "123"
	c = 20
	fmt.Println(a, b, c)
	var d int = 5
	var e = 6
	f := 7
	fmt.Println(d, e, f)

	fmt.Println(aa, bb)
}
