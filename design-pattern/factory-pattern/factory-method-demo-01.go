package factory_pattern

import "fmt"

type People struct {
    name string
	age int
}

func NewPeopleFactory(age int) func(name string) People  {
	return func(name string) People {
		return People{
			name: name,
			age: age,
		}
	}
}

func T()  {
	newBaby := NewPeopleFactory(18)
	baby := newBaby("backlash")
	fmt.Println(baby)
}



