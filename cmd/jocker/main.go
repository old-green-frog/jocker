package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		fmt.Println(!unicode.IsNumber(c), !unicode.IsSymbol(c))
		return !unicode.IsNumber(c) && !unicode.IsSymbol(c)
	}

	sr := strings.FieldsFunc("1 + 2", f)
	fmt.Println(sr)
}
