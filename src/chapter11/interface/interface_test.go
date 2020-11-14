package _interface

import "testing"

/*
	Go接口
*/
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "Java Hello World"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

	p = new(JavaProgrammer)
	t.Logf(p.WriteHelloWorld())
}
