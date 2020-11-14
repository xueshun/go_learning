package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Print(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Dog Wang...")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("Dog SpeakTo ", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("hello")
}

/*
	测试类型转换
*/
func TestTypeConversion(t *testing.T) {

	//var dog Pet = new(Dog)
	//var dog Pet = Pet(new(Dog))
	// var dog Dog = new(Dog)

}
