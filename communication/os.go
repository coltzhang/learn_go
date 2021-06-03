package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}
}
