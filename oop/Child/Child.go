package Child

import (
	"sandbox-go/oop/Parent"
)

type Child struct {
	Parent.Parent
}

func New(firstname string, p *Parent.Parent) *Child {
	return &Child{*Parent.New(firstname, p.GetName())}
}

func (c Child) Sleep() string {
	return "early"
}

func (C Child) WakeUp() string {
	return "early"
}
