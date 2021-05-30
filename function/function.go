package main

import (
	"fmt"
	"strconv"
)

func func1() (int, int) {
	return 1, 2
}

func func2() {
	println("begin")
	for i := 1; i < 5; i++ {
		defer println(i)
	}
	println("end")
}

func func3(arg ...int) (sum int) { // 变参传递的实际是切片
	for _, val := range arg {
		sum += val
	}
	return
}

func func4() {
	a := func() {
		println("niming function")
	}
	a()
}

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(num int) {
	if s.i > 9 {
		return
	}
	s.data[s.i] = num
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

// fmt.Printf("\%v") 可以打印实现了 String() 接口的任何值（%v）
func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str += "[" + strconv.Itoa(i) + ": " + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func max(param []int) int {
	maxNum := param[0]
	for _, v := range param {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

type Website struct {
	Name string
}

func main() {
	_, b := func1()
	println(b)
	func2()
	num := []int{1, 2, 3, 4}
	sum := func3(num...)
	println("sum:", sum) // 变参传递需要指明...
	func4()

	s := new(stack)
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Printf("%+v\n", s)
	println(s.pop())
	println(s.pop())
	println(s.pop())

	num1 := []int{1, 4, 2, 5, 3}
	num2 := make([]int, 5)
	num2[0] = 1
	num2[1] = 6
	num2[2] = 3
	println("max: ", max(num1))
	println("max: ", max(num2))

	// 定义结构体变量
	var site = Website{Name: "studygolang"}
	fmt.Printf("%v\n", site)
	fmt.Printf("%+v\n", site)
	fmt.Printf("%#v\n", site)
	fmt.Printf("%T\n", site)

}
