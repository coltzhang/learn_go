package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err == nil {
		fmt.Println(string(b))
	}
}
