package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "GoProgrammer"
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "JavaProgrammer"
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goGrog := new(GoProgrammer)
	javaGrog := &JavaProgrammer{}
	writeFirstProgrammer(goGrog)
	writeFirstProgrammer(javaGrog)
}
