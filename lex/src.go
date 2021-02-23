package lex

import (
	"fmt"
	"strings"
)

///////////

const (
	ADD  = "+"
	SUB  = "-"
	MULT = "*"
	DIV  = "/"
	NUMS = "0123456789"
	DOT  = "."
)

type Node struct {
	Source  string
	pos     int
	curChar string
}

//////////////////

func (n *Node) adv() {
	n.pos++
	if n.pos < len(n.Source) {
		n.curChar = string(n.Source[n.pos])
	} else {
		n.curChar = "\n"
	}

}

/////////////////
//Initialize node
func (n *Node) Init() {

	// f := func(c rune) bool {
	// 	return !unicode.IsDigit(c) && !strings.ContainsAny(string(c), "+-*/.")
	// }

	// n.TrueS = strings.FieldsFunc(n.Source, f)
	// for _, value := range n.Source {
	// 	if strings.ContainsAny(string(value), " ") {
	// 		continue
	// 	} else if strings.ContainsAny(string(value), NUMS) {

	// 		// makeNum := func() string {
	// 		// 	numT := ""
	// 		// 	dotCount := 0

	// 		// 	for string(value) == io.EOF.Error() && strings.ContainsAny(string(value), NUMS+DOT) {
	// 		// 		if string(value) == DOT {
	// 		// 			if dotCount == 1 {
	// 		// 				break
	// 		// 			}
	// 		// 			dotCount++
	// 		// 			numT += DOT
	// 		// 		} else {
	// 		// 			numT += string(value)
	// 		// 		}
	// 		// 	}
	// 		// 	fmt.Println(numT)
	// 		// 	return numT
	// 		// }

	// 		numT := ""
	// 		dotCount := 0
	// 		for string(value) != io.EOF.Error() && strings.ContainsAny(string(value), NUMS+DOT) {
	// 			if string(value) == DOT {
	// 				if dotCount == 1 {
	// 					break
	// 				}
	// 				dotCount++
	// 				numT += DOT
	// 			} else {
	// 				numT += string(value)
	// 			}
	// 			continue
	// 		}

	// 		n.TrueS = append(n.TrueS, numT)

	// 	} else if strings.ContainsAny(string(value), ADD+SUB+DIV+MULT) {
	// 		n.TrueS = append(n.TrueS, string(value))
	// 	}
	//}

	n.pos = -1
	n.curChar = "\n"
	n.adv()
}

//////////////////////////

func (n Node) build() []string {

	var trueS []string

	for n.curChar != "\n" {
		if strings.ContainsAny(n.curChar, "       ") {
			n.adv()
		} else if strings.ContainsAny(n.curChar, ADD) {
			trueS = append(trueS, ADD)
			n.adv()
		} else if strings.ContainsAny(n.curChar, SUB) {
			trueS = append(trueS, SUB)
			n.adv()
		} else if strings.ContainsAny(n.curChar, MULT) {
			trueS = append(trueS, MULT)
			n.adv()
		} else if strings.ContainsAny(n.curChar, DIV) {
			trueS = append(trueS, DIV)
			n.adv()
		} else if strings.ContainsAny(n.curChar, NUMS) {

			num := ""

			for !strings.ContainsAny(n.curChar, " "+ADD+MULT+DIV+SUB) {

				num = num + n.curChar
				n.adv()

				if n.curChar == "\n" {
					break
				}
				// fmt.Printf("Source: %s\tpos: %d\tcurChar: %s\tnum: %s\n", n.Source, n.pos, n.curChar, num)
				//				fmt.Println(!strings.Contains(n.curChar, " \n"+ADD+MULT+DIV+SUB))
			}

			trueS = append(trueS, num)

		}
	}
	//for testing:
	//fmt.Println(trueS[0])
	return trueS
}

func (n Node) String() string {
	return fmt.Sprintf("%v", n.build())
}
