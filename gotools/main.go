package main

import (
	"flag"
	"fmt"
	"unicode"
)

func main() {

	var name string

	flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
	flag.Parse()

	fmt.Printf("name: %s \n", name)
	s := "ab"
	fmt.Println(string(unicode.ToUpper(rune(s[0]))) + s[1:])
}
