package factory_pattern

import "fmt"

type Person struct {
	Name string
	Age int
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func (p Person) Greet()  {
	fmt.Printf("Hi! My name is %s", p.Name)
}