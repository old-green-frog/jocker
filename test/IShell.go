package main

import (
	"bufio"
	"fmt"
	"os"

	"jocker/lex"
)

func main() {

	fmt.Println("Welcome to the IShell!")

	for {
		fmt.Print("--> ")

		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		node := lex.Node{Source: sc.Text()}
		node.Init()

		fmt.Println(node)
	}
}
