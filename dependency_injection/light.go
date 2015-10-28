package main

import (
	"fmt"
	"log"
)

type OnOff interface {
	On() error
	Off() error
}

// Light implements OnOff
type Light struct {
	age *int
}

func NewLight() Light {
	return Light{
		age: new(int),
	}
}

func (light Light) On() error {
	*light.age++
	if *light.age == 5 {
		return fmt.Errorf("The light broke")
	}
	fmt.Println("The light is on")
	return nil
}

func (light Light) Off() error {
	fmt.Println("The light is off")
	return nil
}

// Switch uses an OnOff
type Switch struct {
	isOn   bool
	outlet *OnOff
}

func NewSwitch(outlet *OnOff) Switch {
	return Switch{
		isOn:   false,
		outlet: outlet,
	}
}

func (s *Switch) Flip() error {
	s.isOn = !s.isOn
	if s.isOn {
		return (*s.outlet).On()
	}
	return (*s.outlet).Off()
}

func main() {

	var l OnOff
	l = NewLight()
	s := NewSwitch(&l)

	for {
		err := s.Flip()
		if err != nil {
			log.Fatal(err)
		}
	}
}
