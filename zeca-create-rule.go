package main

import "fmt"

type Pattern [3]string
type Probability map[string]float64

type Rule map[Pattern]Probability

var aa = []string{
	"#", "A", "C", "D", "E", "F", "G", "H", "I", "K", "L",
	"M", "N", "P", "Q", "R", "S", "T", "V", "Y", "W"}

var ss = []string{"-", "*", "|", "?"}

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

	for c := 0; c < len(cell); c++ {
		for ln := 0; ln < len(cell); ln++ {
			for rn := 0; rn < len(cell); rn++ {

				fmt.Printf("[%s %s %s]\n", cell[ln], cell[c], cell[rn])
				// if h != nil {
				//
				// } else {
				// 	rule[Pattern{aa[ln], aa[c], aa[rn]}] = Probability
				// }
			}
		}
	}
	return rule
}

func rule(hydro string) {
	// m := make([]string, r)
	// for i := 0; i < r; i++ {
	//
	// }
}

func main() {
	create("roseSpecialCharged")
}
