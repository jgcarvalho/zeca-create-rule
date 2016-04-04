package main

import (
	"flag"
	"fmt"
)

type Pattern [3]string
type Probability map[string]float64

type Rule map[Pattern]Probability

var aa = []string{
	"#", "A", "C", "D", "E", "F", "G", "H", "I", "K", "L",
	"M", "N", "P", "Q", "R", "S", "T", "V", "Y", "W"}

var ss = []string{"_", "*", "|", "?"}

// rose: polar or nonpolar
var rose = map[string]string{
	"A": "n", "C": "n", "V": "n", "I": "n", "L": "n", "M": "n", "F": "n",
	"W": "n", "G": "p", "S": "p", "T": "p", "H": "p", "Y": "p", "P": "p",
	"D": "p", "N": "p", "E": "p", "Q": "p", "K": "p", "R": "p"}

// rose special: polar, nonpolar, Gly and Pro
var roseSpecial = map[string]string{
	"A": "n", "C": "n", "V": "n", "I": "n", "L": "n", "M": "n", "F": "n",
	"W": "n", "G": "G", "S": "p", "T": "p", "H": "p", "Y": "p", "P": "P",
	"D": "p", "N": "p", "E": "p", "Q": "p", "K": "p", "R": "p"}

//rose special charged: polar, nonpolar, Gly, Pro, positives and negatives
var roseSpecialCharged = map[string]string{
	"A": "n", "C": "n", "V": "n", "I": "n", "L": "n", "M": "n", "F": "n",
	"W": "n", "G": "G", "S": "p", "T": "p", "H": "p", "Y": "p", "P": "P",
	"D": "-", "N": "p", "E": "-", "Q": "p", "K": "+", "R": "+"}

func create(hydro string) Rule {
	rule := make(Rule)

	var cell []string

	switch hydro {
	case "rose":
		cell = make([]string, len(aa)+3*2)
		for i := 0; i < len(aa); i++ {
			cell[i] = aa[i] + rose[aa[i]]
		}
		i := len(aa)
		for _, v := range []string{"n", "p"} {
			for _, u := range ss[:len(ss)-1] {
				cell[i] = u + v
				i++
			}
		}

	case "roseSpecial":

		cell = make([]string, len(aa)+3*4)
		for i := 0; i < len(aa); i++ {
			cell[i] = aa[i] + roseSpecial[aa[i]]
		}
		i := len(aa)
		for _, v := range []string{"n", "p", "G", "P"} {
			for _, u := range ss[:len(ss)-1] {
				cell[i] = u + v
				i++
			}
		}

	case "roseSpecialCharged":
		cell = make([]string, len(aa)+3*6)
		for i := 0; i < len(aa); i++ {
			cell[i] = aa[i] + roseSpecialCharged[aa[i]]
		}
		i := len(aa)
		for _, v := range []string{"n", "p", "G", "P", "+", "-"} {
			for _, u := range ss[:len(ss)-1] {
				cell[i] = u + v
				i++
			}
		}
	case "SSAA":
		cell = make([]string, len(aa)+3*20)
		for i := 0; i < len(aa); i++ {
			cell[i] = aa[i]
		}
		i := len(aa)
		for _, v := range aa[1:] {
			for _, u := range ss[:len(ss)-1] {
				cell[i] = u + v
				i++
			}
		}

	default:
		cell = make([]string, len(aa)+3*1)
		for i := 0; i < len(aa); i++ {
			cell[i] = aa[i]
		}
		i := len(aa)
		for _, u := range ss[:len(ss)-1] {
			cell[i] = u
			i++
		}
	}
	// count := 0
	for c := 0; c < len(cell); c++ {
		for ln := 0; ln < len(cell); ln++ {
			for rn := 0; rn < len(cell); rn++ {
				pt := Pattern{cell[ln], cell[c], cell[rn]}
				var prob Probability
				fmt.Printf("[ %s ][ %s ][ %s ] -> ", cell[ln], cell[c], cell[rn])
				if cell[c] == "#" {
					prob = Probability{"_": 0.0, "*": 0.0, "|": 0.0, "?": 1.0}
					fmt.Printf("{ _ : 0.0, * : 0.0, | : 0.0, ? : 1.0 }\n")
				} else if len(cell[c]) == 1 {
					prob = Probability{"_": 0.25, "*": 0.25, "|": 0.25, "?": 0.25}
					fmt.Printf("{ _ : 0.25, * : 0.25, | : 0.25, ? : 0.25 }\n")
				} else {
					prob = Probability{"_" + cell[c][1:]: 0.25, "*" + cell[c][1:]: 0.25, "|" + cell[c][1:]: 0.25, "?" + cell[c][1:]: 0.25}
					fmt.Printf("{ _%s : 0.25, *%s : 0.25, |%s : 0.25, ?%s : 0.25 }\n", cell[c][1:], cell[c][1:], cell[c][1:], cell[c][1:])
				}
				rule[pt] = prob
				// count++
			}
		}
	}
	// fmt.Println(count)
	return rule
}

//
// func (r *Rule) String() string {
// 	var toprint string
// 	for _, v := range *r {
// 		toprint += fmt.Sprintf("[%s][%s][%s] -> {} ")
// 	}
//
// }

func main() {
	h := flag.String("h", "", "Hydrophobicity pattern [rose, roseSpecial, roseSpecialCharged]")
	flag.Parse()
	create(*h)
}
