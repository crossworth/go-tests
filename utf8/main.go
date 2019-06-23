package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	opt := "opção"

	fmt.Println(len(opt))
	fmt.Println(utf8.RuneCount([]byte(opt)))
}
