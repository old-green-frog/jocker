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
		}
		huh = ""
	}

	for in, tok := range p.toks {

		switch tok.elType {

		case "INT", "FL":
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

	_, err := strconv.Atoi(huh)
	if err != nil {
		return element{"FL", huh}
	}
	return element{"INT", huh}
}
