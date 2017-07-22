package main

import (
	"fmt"
)

type Comparable interface {
	Less(Comparable) bool
}

type ComparableSlice []Comparable

type MyType int

func main() {

	mt := []MyType{5, 3, 5, 7, 8, 3, 5, 2}

	// This won't compile
	//cs := []Comparable(mt)

	// Instead, you have to make a copy
	cs := make([]Comparable, len(mt))
	for i, v := range mt {
		cs[i] = v
	}

	fmt.Println(ComparableSlice(cs).Max())
}

func (t ComparableSlice) Max() Comparable {
	if len(t) == 0 {
		var v Comparable
		return v
	}

	max := t[0]
	for _, v := range t[1:] {
		if max.Less(v) {
			max = v
		}
	}
	return max
}

func (t MyType) Less(t2 Comparable) bool {
	return t < t2.(MyType)
}
