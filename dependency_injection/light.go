package main

import (
	"fmt"
	"log"
)

type OnOff interface {
	On() error
	Off() error
}

type Light struct {
	age int
}

func (light *Light) On() error {
	light.age++
	if light.age == 5 {
		return fmt.Errorf("The light broke")
	}
	fmt.Println("The light is on")
	return nil
}

func (light *Light) Off() error {
	fmt.Println("The light is off")
	return nil
}

type Switch struct {
	isOn   bool
	outlet OnOff
}

func (s *Switch) Flip() error {
	s.isOn = !s.isOn
	if s.isOn {
		return s.outlet.On()
	}
	return s.outlet.Off()
}

func main() {

	l := Light{}
	s := Switch{outlet: &l} // *Light implements OnOff

	for {
		err := s.Flip()
		if err != nil {
			log.Fatal(err)
		}
	}
}
