package main

import "fmt"

func main() {
	f1 := te1(100)
	f1()

	for _, f := range te2() {
		f()
	}

	for _, f := range te3() {
		f()
	}

	a, b := te4(100)
	a()
	b()
}

func te1(x int) func() {
	fmt.Println(x)

	return func() {
		fmt.Println(&x, x)
	}

}

func te2() []func() {
	var a []func()
	for i := 0; i < 2; i++ {
		a = append(a, func() {
			fmt.Println(&i, i)
		})
	}

	return a
}

func te3() []func() {
	var a []func()
	for i := 0; i < 2; i++ {
		x := i
		a = append(a, func() {
			fmt.Println(&x, x)
		})
	}

	return a
}

func te4(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x += 100
		}, func() {
			fmt.Println(x)
		}

}
