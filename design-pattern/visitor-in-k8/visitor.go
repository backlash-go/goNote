package main

import "fmt"

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func() error

type VisitorList []Visitor

func (v VisitorList) Visit(fn VisitorFunc) error {
	for i := range v {
		if err := v[i].Visit(func() error {
			fmt.Println("In VisitorList before fn")
			fn()
			fmt.Println("In VisitorList after fn")
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

type Visitor1 struct {
}

func (v Visitor1) Visit(fn VisitorFunc) error {
	fmt.Println("In Visitor1 before fn")
	fn()
	fmt.Println("In Visitor1 after fn")
	return nil
}

type Visitor2 struct {
	visitor Visitor
}

func (v Visitor2) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("In Visitor2 before fn")
		fn()
		fmt.Println("In Visitor2 after fn")
		return nil
	})
	return nil
}

type Visitor3 struct {
	visitor Visitor
}

func (v Visitor3) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("In Visitor3 before fn")
		fn()
		fmt.Println("In Visitor3 after fn")
		return nil
	})
	return nil
}

func main() {

	//组成这个结构 --> Visitor3{Visitor2{VisitorList{Visitor1}}}
	var visitor Visitor

	var visitors []Visitor

	visitor = Visitor1{}

	visitors = append(visitors, visitor)

	visitor = Visitor2{VisitorList(visitors)}

	visitor = Visitor3{visitor}

	visitor.Visit(func() error {
		fmt.Println("In visitFunc")
		return nil
	})

}
