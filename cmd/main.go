package main

import (
	"fmt"
	"github.com/bgzxz/gomod/v2"
)

func main() {
	g, err := gomod.HelloWorld("zhang", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}
