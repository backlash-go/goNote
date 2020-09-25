package main

import "fmt"

func main() {
	var digitalList1 = []int64{1, 2, 3, 4, 5}
	var digitalList2 = []int64{10, 20, 30, 40}

	fmt.Println(digitalList1, digitalList2)

	d1 := append(digitalList1, digitalList2...)
	fmt.Println(d1)

	d2 := append(digitalList1, 100, 200)
	fmt.Println(d2)

	hello("a")

	hellod(digitalList1...)

}

func hello(name ...string) {
	fmt.Printf("%T\n", name)
	return

}

func hellod(name ...int64) {
	fmt.Printf("%T\n", name)
	return

}
