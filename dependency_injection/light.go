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

func NewLight() Light {
	return Light{
		age: 0,
	}
}

func (light *Light) On() error {
	if light.age == 5 {
		return fmt.Errorf("The light broke")
	}
	light.age++
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

func NewSwitch(outlet OnOff) Switch {
	return Switch{
		isOn:   false,
		outlet: outlet,
	}
}

func (s *Switch) Flip() error {
	s.isOn = !s.isOn
	if s.isOn {
		return s.outlet.On()
	}
	return s.outlet.Off()
}

func main() {

	l := NewLight()
	s := NewSwitch(&l) // *Light implements OnOff

	for {
		err := s.Flip()
		if err != nil {
			log.Fatal(err)
		}
	}
}
