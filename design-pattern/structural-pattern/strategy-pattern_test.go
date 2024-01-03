package structural_pattern

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	operator := Operator{}
	operator.setStrategy(&add{})
	result := operator.calculate(100, 230)
	fmt.Println("add:", result)

	operator.setStrategy(&reduce{})
	result2 := operator.calculate(100, 230)
	fmt.Println("reduce:", result2)

}
