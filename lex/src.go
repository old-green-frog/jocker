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

//////////////////
//ELEMENTS
type element struct {
	elType  string
	elValue string
}

func (el element) String() string {
	return fmt.Sprintf("%s: %s", el.elType, el.elValue)
}

//optional
// func removeEl(s []element, i int) []element {

// 	copy(s[i:], s[i+1:])
// 	s[len(s)-1] = element{}
// 	s = s[:len(s)-1]

// 	return s
// }

// func getInt(str string) int {

// 	res, err := strconv.Atoi(str)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return res
// }

// func getFl(str string) float32 {

// 	resFloat, err := strconv.ParseFloat(str, 32)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return float32(resFloat)
// }

///////////////////
//NODES

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

	n.pos = -1
	n.curChar = "\n"
	n.adv()
}

//////////////////////////

func (n Node) build() []element {

	var trueS []element

	for n.curChar != "\n" {
		if strings.ContainsAny(n.curChar, "       ") {
			n.adv()
		} else if strings.ContainsAny(n.curChar, ADD) {
			trueS = append(trueS, element{"ADD", ADD})
			n.adv()
		} else if strings.ContainsAny(n.curChar, SUB) {
			trueS = append(trueS, element{"SUB", SUB})
			n.adv()
		} else if strings.ContainsAny(n.curChar, MULT) {
			trueS = append(trueS, element{"MULT", MULT})
			n.adv()
		} else if strings.ContainsAny(n.curChar, DIV) {
			trueS = append(trueS, element{"DIV", DIV})
			n.adv()
		} else if strings.ContainsAny(n.curChar, NUMS) {

			num := ""
			dotCount := 0

			for !strings.ContainsAny(n.curChar, " "+ADD+MULT+DIV+SUB) {

				if n.curChar == DOT {
					if dotCount >= 1 {
						return []element{{"ERROR", "To many dots!"}}
					}
					dotCount++
				}
				num = num + n.curChar
				n.adv()

				if n.curChar == "\n" {
					break
				}
			}

			if dotCount == 1 {
				trueS = append(trueS, element{"FL", num})
			} else {
				trueS = append(trueS, element{"INT", num})
			}

		}
	}
	//for testing:
	//fmt.Println(trueS[0])
	return trueS
}

// func (n Node) run() string {

// 	src := n.build()
// 	var res interface{}
// 	var trueRes interface{}

// 	for in, el := range src {

// 		if strings.ContainsAny(el.elType, "INT"+"FL") {
// 			continue
// 		} else {
// 			switch el.elType {

// 			case "ADD":
// 				{
// 					if src[in-1].elType == "FL" {
// 						if src[in+1].elType == "FL" {
// 							res = getFl(src[in+1].elValue) + getFl(src[in-1].elValue)
// 						} else {
// 							res = float32(getInt(src[in+1].elValue)) + getFl(src[in-1].elValue)
// 						}
// 					} else {
// 						if src[in+1].elType == "FL" {
// 							res = getFl(src[in+1].elValue) + float32(getInt(src[in-1].elValue))
// 						} else {
// 							res = getInt(src[in+1].elValue) + getInt(src[in-1].elValue)
// 						}
// 					}
// 				}

// 			}
// 			src = removeEl(src, in+1)
// 			src = removeEl(src, in)
// 			src = removeEl(src, in-1)

// 		}
// 	}
// 	return fmt.Sprintf("%v", res)
// }

func (n Node) String() string {
	return fmt.Sprintf("%v", n.build())
}

//for future.....

// res, err := strconv.Atoi(intString)
// if err != nil {
// 	panic(err)
// }
// fmt.Printf("Parsed integer: %d\n", res)

// resFloat, err := strconv.ParseFloat(floatString, 32)
// if err != nil {
// 	panic(err)
// }
// fmt.Printf("Parsed float: %.5f\n", resFloat)
