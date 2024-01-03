package factory_pattern

import "fmt"

type Product interface {
    Use()
}


type ConcreteProductA struct {
    
}

func (c ConcreteProductA) Use() {
	fmt.Println("Product  A  Use")
}


type ConcreteProductB struct {

}

func (c ConcreteProductB) Use() {
	fmt.Println("Product  B  Use")
}




//define a factory interface


type Factory interface {
	CreateProduct() Product
}

type ConcreteFactoryA struct{}

func (c ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteFactoryB struct{}

// CreateProduct 实现了 Factory 接口的 CreateProduct 方法
func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}




















