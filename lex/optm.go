package lex

import (
	"fmt"
	"strconv"
)

////error structure
type error struct {
	details string
}

func (e error) genErr(d string) element {
	e.details = d
	return element{"ERROR", e.details}
}

func removeEl(s []element, i int) []element {

	copy(s[i:], s[i+1:])
	s[len(s)-1] = element{}
	s = s[:len(s)-1]

	return s
}

func getInt(str string) int {

	res, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return res
}

func getFl(str string) float32 {

	resFloat, err := strconv.ParseFloat(str, 32)

	if err != nil {
		panic(err)
	}

	return float32(resFloat)
}

func (p *pars) insert(el element, ind int) {

	// last := len(p.toks) - 1
	// p.toks = append(p.toks, p.toks[last])

	// copy(p.toks[2:], p.toks[1:last])

	// p.toks[ind] = el

	p.toks = append(p.toks, element{})
	copy(p.toks[ind+1:], p.toks[ind:])
	p.toks[ind] = el

	// if len(a) == index { // nil or empty slice or after last element
	// 	return append(a, value)
	// }
	// a = append(a[:index+1], a[index:]...) // index < len(a)
	// a[index] = value
	// return a

	///////////////////////////////////
	// i := sort.SearchStrings(ss, s)
	// ss = append(ss, "")
	// copy(ss[i+1:], ss[i:])
	// ss[i] = s
	// return ss
}

////////////////////////
////IMPORTANT PART OF PARSER
////////////////////////
func (p pars) consNum(s *string, ind int, dst string, t string) {

	switch dst {
	case "ADD":
		if t == "INT" {
			res := getInt(p.toks[ind-1].elValue) + getInt(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%d", getInt(*s)+getInt(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%d", res)
			}
		} else if t == "FL" {
			res := getFl(p.toks[ind-1].elValue) + getFl(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%f", getFl(*s)+getFl(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%f", res)
			}
		}

	case "SUB":
		if t == "INT" {
			res := getInt(p.toks[ind-1].elValue) - getInt(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%d", getInt(*s)-getInt(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%d", res)
			}
		} else if t == "FL" {
			res := getFl(p.toks[ind-1].elValue) - getFl(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%f", getFl(*s)-getFl(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%f", res)
			}
		}

	case "MULT":
		if t == "INT" {
			res := getInt(p.toks[ind-1].elValue) * getInt(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%d", getInt(*s)*getInt(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%d", res)
			}
		} else if t == "FL" {
			res := getFl(p.toks[ind-1].elValue) * getFl(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%f", getFl(*s)*getFl(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%f", res)
			}
		}

	case "DIV":
		if t == "INT" {
			res := getInt(p.toks[ind-1].elValue) / getInt(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%d", getInt(*s)/getInt(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%d", res)
			}
		} else if t == "FL" {
			res := getFl(p.toks[ind-1].elValue) / getFl(p.toks[ind+1].elValue)
			if *s != "" {
				*s = fmt.Sprintf("%f", getFl(*s)/getFl(p.toks[ind+1].elValue))
			} else {
				*s = fmt.Sprintf("%f", res)
			}
		}
	}
}
