package lex

import "strconv"

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
