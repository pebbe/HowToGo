package main

import (
	"fmt"
	"sort"
)

type Sorter struct {
	s []int
	f func(int, int) bool
}

func (s *Sorter) Len() int {
	return len(s.s)
}

func (s *Sorter) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func (s *Sorter) Less(i, j int) bool {
	return s.f(s.s[i], s.s[j])
}

func main() {
	a := []int{3, 6, 8, 1, 2, 9, 7, 0, 4, 5}
	sort.Sort(&Sorter{a, func(a, b int) bool { return a > b }})
	fmt.Println(a)
	sort.Sort(&Sorter{a, func(a, b int) bool { return a < b }})
	fmt.Println(a)
}
