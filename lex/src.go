package lex

import (
	"fmt"
	"strings"

	"jocker/pkg"
)

///////////

const (
	ADD     = "+"
	SUB     = "-"
	MULT    = "*"
	DIV     = "/"
	NUMS    = "0123456789"
	DOT     = "."
	LETTERS = "abcdefghijklmnopqrstuvwxyz"
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

	n.pos = -1
	n.curChar = "\n"
	n.adv()
}

//////////////////////////

func (n Node) build() []element {

	var trueS []element
	err := error{}

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
						return []element{err.genErr("To many dots!")}
					}
					dotCount++
				}

				num = num + n.curChar
				n.adv()

				if n.curChar == "\n" {
					break
				}
			}

			// fmt.Println(num)

			if dotCount == 1 {
				trueS = append(trueS, element{"FL", num})
			} else {
				trueS = append(trueS, element{"INT", num})
			}

		} else if strings.Contains(n.curChar, "\"") {

			str := ""
			n.adv()

			for !strings.ContainsAny(n.curChar, "\"") {
				str = str + n.curChar
				n.adv()
			}

			n.adv()
			trueS = append(trueS, element{"STR", str})

			/////vars

		} else if n.curChar == "v" && n.curChar == string(n.Source[0]) {

			brr := ""
			for n.curChar != " " {
				// fmt.Println(1)
				brr = brr + n.curChar
				n.adv()
			}
			n.adv()

			// fmt.Println(brr)
			if brr == "var" {

				name := ""

				for !strings.ContainsAny(n.curChar, " =") {
					// fmt.Println(2)
					name = name + n.curChar
					n.adv()
				}

				for !strings.ContainsAny(n.curChar, NUMS) {
					n.adv()
				}

				// fmt.Println(n.Source)
				n.Source = n.Source[n.pos:]
				n.pos = 0
				// fmt.Println(n.Source)
				varDecl(stack, name, n.run())
				return []element{stack[name]}

			}

			continue

			///////////func

		} else if n.curChar == "p" && n.curChar == string(n.Source[0]) {

			brr := ""
			for n.curChar != " " {
				// fmt.Println(1)
				brr = brr + n.curChar
				n.adv()
			}
			n.adv()

			if brr == "println" {

				n.Source = n.Source[n.pos:]
				n.pos = 0

				res := n.run()

				pkg.Println(res.elValue)
			}

		} else if strings.ContainsAny(n.curChar, LETTERS) {

			word := ""

			for strings.ContainsAny(n.curChar, LETTERS) {
				word = word + n.curChar
				n.adv()
			}

			var cont interface{} = stack[word]

			if el, ok := cont.(element); ok {
				if elIsEmpty(el) {
					return []element{{"ERROR", "What?"}}
				}
				return []element{el}
			}
		}
	}
	//for testing:
	// fmt.Println(trueS)

	return trueS
}

func (n Node) run() element {

	prs := pars{}
	prs.toks = n.build()
	return prs.parse()
}

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
	n.run()
	return ""
}

////////////////
////////////////
////////////////
///VARIABLES
////////////////
////////////////
////////////////

////////////////
////////////////

var stack = make(map[string]element)

func varDecl(s map[string]element, name string, val interface{}) {

	rr := error{}
	switch tv := val.(type) {

	case element:
		s[name] = element{tv.elType, tv.elValue}
	default:
		fmt.Println(rr.genErr("Strange type here"))
		fmt.Printf("%T: %v\n", val, val)
	}
}
