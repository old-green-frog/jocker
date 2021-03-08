package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"jocker/lex"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		node := lex.Node{Source: scanner.Text()}
		node.Init()

		fmt.Print(node)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
