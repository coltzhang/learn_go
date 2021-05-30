package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

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

	g := "测试"
	g = "abcdefghijk"
	fmt.Println([]rune(g)) // 4字节
	fmt.Println([]byte(g)) // 1字节

	s := "qqqqq" +
		"wwwww"
	fmt.Println(s)

	//var err error
	//fmt.Println(err)

	if err := os.Chmod("111", 0664); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	for k, v := range "12345" {
		fmt.Printf("%d %c\n", k, v)
	}
	print("%d", 11)

	str := "dsjkdshdjsdh....js测试"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n %d", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)), len(str))
}
