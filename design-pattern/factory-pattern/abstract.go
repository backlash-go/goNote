package factory_pattern

type PersonSet interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	panic("implement me")
}

func newPerson(name string, age int) PersonSet {
	return person{name: name, age: age}
}







