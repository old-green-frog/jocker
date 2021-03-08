package main

import (
	"bufio"
	"fmt"
	"os"

	"jocker/lex"
)

func main() {

	fmt.Println("Welcome to the IShell!")
	// pkg.Println("Hello")

	for {
		fmt.Print("--> ")

		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		if sc.Text() == "exit" {
			os.Exit(0)
		}
		node := lex.Node{Source: sc.Text()}
		node.Init()

		fmt.Println(node)
	}
}
