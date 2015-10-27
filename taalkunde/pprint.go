package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "vpx(v(_,vform,[vc([],inf,_,[],_,_)|_],[np(nom,agr,(u;thi;ninvje;fir),(year;temp_meas_mod;temp;refl;norm;meas_mod;het_nform;er;cleft_het),_,_,_)],_,_),fin,[np(nom,sg&de&indef,thi,norm,nwh,[],PER)],[])"

	fmt.Print("\nVERSION 1:\n\n")
	pprint1(s)
	fmt.Print("\nVERSION 1b:\n\n")
	pprint1b(s)
	fmt.Print("\nVERSION 2:\n\n")
	pprint2(s)

}

func pprint1(s string) {

	stack := make([]int, len(s))
	stack[0] = 0
	sp := 0
	op := 0
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			fmt.Print(string(r))
			sp++
			op++
			stack[sp] = op
		case ')', ']', '}':
			fmt.Print(string(r))
			if sp > 0 {
				sp--
			}
			op++
		case ',':
			op = stack[sp]
			fmt.Print(string(r), "\n", strings.Repeat(" ", op))
		default:
			fmt.Print(string(r))
			op++
		}
	}
	fmt.Println()
}

func pprint1b(s string) {

	stack := make([]int, len(s))
	stack[0] = 0
	sp := 0
	op := 0
	for i, r := range s {
		switch r {
		case '(', '[', '{':
			if i < len(s)-1 && (s[i+1] == ')' || s[i+1] == ']' || s[i+1] == '}') {
				fmt.Print(string(r))
				op++
			} else {
				fmt.Print(string(r), " ")
				op += 2
			}
			sp++
			stack[sp] = op
		case ')', ']', '}':
			if i > 0 && (s[i-1] == '(' || s[i-1] == '[' || s[i-1] == '{') {
				fmt.Print(string(r))
				op++
			} else {
				fmt.Print(" ", string(r))
				op += 2
			}
			if sp > 0 {
				sp--
			}
		case ',':
			op = stack[sp]
			fmt.Print(string(r), "\n", strings.Repeat(" ", op))
		default:
			fmt.Print(string(r))
			op++
		}
	}
	fmt.Println()
}

func pprint2(s string) {
	lvl := 0
	for i, r := range s {
		switch r {
		case '(', '[', '{':
			lvl += 4
			if i < len(s)-1 && (s[i+1] == ')' || s[i+1] == ']' || s[i+1] == '}') {
				fmt.Print(string(r))
			} else {
				fmt.Print(string(r), "\n", strings.Repeat(" ", lvl))
			}
		case ')', ']', '}':
			if lvl > 0 {
				lvl -= 4
			}
			if i > 0 && (s[i-1] == '(' || s[i-1] == '[' || s[i-1] == '{') {
				fmt.Print(string(r))
			} else {
				fmt.Print("\n", strings.Repeat(" ", lvl), string(r))
			}
		case ',':
			fmt.Print(string(r), "\n", strings.Repeat(" ", lvl))
		default:
			fmt.Print(string(r))
		}
	}
	fmt.Println()
}
