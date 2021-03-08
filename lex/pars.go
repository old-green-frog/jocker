package lex

import (
	"strconv"
)

///////////////////
//SIMPLE PARSER
///////////////////

type pars struct {
	toks []element
}

func (p *pars) parse() element {

	huh := ""
	// fmt.Println(p.toks)
	if len(p.toks) == 1 && p.toks[0].elType == "STR" {
		return p.toks[0]
	}

	for in, tok := range p.toks {

		switch tok.elType {

		case "ADD", "SUB", "INT", "FL":
			continue

		case "MULT":
			{
				if p.toks[in-1].elType == "FL" || p.toks[in+1].elType == "FL" {

					p.consNum(&huh, in, "MULT", "FL")
					// fmt.Println(p.toks, " M")
					p.toks = removeEl(p.toks, in+1)
					p.toks = removeEl(p.toks, in-1)

					p.insert(element{"FL", huh}, in-1)
					p.toks = removeEl(p.toks, in)
					// fmt.Println(p.toks, " M")

				} else if p.toks[in-1].elType == "INT" && p.toks[in+1].elType == "INT" {

					p.consNum(&huh, in, "MULT", "INT")
					// fmt.Println(p.toks, " M")
					p.toks = removeEl(p.toks, in+1)
					p.toks = removeEl(p.toks, in-1)

					p.insert(element{"INT", huh}, in-1)
					p.toks = removeEl(p.toks, in)
					// fmt.Println(p.toks, " M")
				}
			}

		case "DIV":
			{
				if p.toks[in-1].elType == "FL" || p.toks[in+1].elType == "FL" {

					p.consNum(&huh, in, "DIV", "FL")
					// fmt.Println(p.toks, " M")
					p.toks = removeEl(p.toks, in+1)
					p.toks = removeEl(p.toks, in-1)

					p.insert(element{"FL", huh}, in-1)
					p.toks = removeEl(p.toks, in)
					// fmt.Println(p.toks, " M")

				} else if p.toks[in-1].elType == "INT" && p.toks[in+1].elType == "INT" {

					p.consNum(&huh, in, "DIV", "INT")
					// fmt.Println(p.toks, " M")
					p.toks = removeEl(p.toks, in+1)
					p.toks = removeEl(p.toks, in-1)

					p.insert(element{"INT", huh}, in-1)
					p.toks = removeEl(p.toks, in)
					// fmt.Println(p.toks, " M")
				}
			}
			if len(p.toks) == 1 {
				break
			} else {
				huh = ""
			}
		}
	}

	huh = ""

	for in, tok := range p.toks {

		switch tok.elType {

		case "INT", "FL", "STR":
			continue

		case "ADD":
			{
				if p.toks[in-1].elType == "FL" || p.toks[in+1].elType == "FL" {

					p.consNum(&huh, in, "ADD", "FL")
					// fmt.Println(p.toks, " A")
					p.toks[in+1] = element{"FL", huh}
					// fmt.Println(p.toks, " A")

				} else if p.toks[in-1].elType == "INT" && p.toks[in+1].elType == "INT" {

					p.consNum(&huh, in, "ADD", "INT")
					// fmt.Println(p.toks, " A")
					p.toks[in+1] = element{"INT", huh}
					// fmt.Println(p.toks, " A")
				} else if p.toks[in-1].elType == "STR" && p.toks[in+1].elType == "STR" {

					p.toks[in+1] = element{"STR", p.toks[in-1].elValue + p.toks[in+1].elValue}
					// fmt.Println(p.toks, " A")

					huh = p.toks[in+1].elValue
				}

			}

		case "SUB":
			{
				if p.toks[in-1].elType == "FL" || p.toks[in+1].elType == "FL" {

					p.consNum(&huh, in, "SUB", "FL")
					// fmt.Println(p.toks, " S")
					p.toks[in+1] = element{"FL", huh}
					// fmt.Println(p.toks, " S")

				} else if p.toks[in-1].elType == "INT" && p.toks[in+1].elType == "INT" {

					p.consNum(&huh, in, "SUB", "INT")
					// fmt.Println(p.toks, " S")
					p.toks[in+1] = element{"INT", huh}
					// fmt.Println(p.toks, " S")
				}

			}

		}
		// p.toks = removeEl(p.toks, in+1)
		// p.toks = removeEl(p.toks, in)

	}

	cont := p.toks[len(p.toks)-1].elValue

	_, err := strconv.Atoi(cont)
	if err != nil {
		if _, e := strconv.ParseFloat(cont, 32); e == nil {
			return element{"FL", cont}
		} else {
			return element{"STR", cont}
		}

	}
	return element{"INT", cont}
}
