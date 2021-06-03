package main

import "fmt"

func main() {
	num := []int{1, 2, 3}
	add := make([]int, 4)
	add[0] = 1
	add[1] = 2
	num = append(num, 5, 4, 7)
	num = append(num, add...)
	fmt.Println(num)
}
