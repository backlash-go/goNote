package main

import "fmt"



func main() {
	//go f1()
	//go f2()
	//a := 10

	ss1 := make([]int,3,5)
	fmt.Printf("%v,%p,%p\n",ss1,ss1,&ss1)
	ss1 = append(ss1,1)
	fmt.Printf("%v,%p,%p\n",ss1,ss1,&ss1)


	ss1 = append(ss1,11)
	fmt.Printf("%v,%p,%p\n",ss1,ss1,&ss1)


	ss1 = append(ss1,111)
	fmt.Printf("%v,%p,%p\n",ss1,ss1,&ss1)



	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)

}
