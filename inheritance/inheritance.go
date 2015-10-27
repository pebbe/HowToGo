package main

import (
	"fmt"
)

// base 'class'

type vehicler interface {
	getCapacity() int
	hasWings() bool
}

type vehicle struct {
	capacity int
}

func (v vehicle) getCapacity() int {
	return v.capacity
}

func (v vehicle) hasWings() bool {
	return false
}

func vehicleInfo(v vehicler) {
	if v.hasWings() {
		fmt.Printf("Persons: %v\nWings: %v\n\n", v.getCapacity(), v.(plane).numWings())
	} else {
		fmt.Printf("Persons: %v\nCan't fly\n\n", v.getCapacity())
	}
}

// derived 'class' car

type car struct {
	vehicle
}

// derived 'class' plane

type plane struct {
	vehicle
	wings int
}

func (p plane) hasWings() bool {
	return true
}

func (p plane) numWings() int {
	return p.wings
}

// main

func main() {
	var ford = car{vehicle{capacity: 5}}
	var audi = car{vehicle{capacity: 4}}
	var jet = plane{vehicle: vehicle{capacity: 200}, wings: 2}

	var e = make([]vehicler, 0, 100)
	e = append(e, ford)
	e = append(e, jet)
	e = append(e, audi)

	fmt.Printf("%#v\n\n", e)

	for _, i := range e {
		vehicleInfo(i)
	}
}
