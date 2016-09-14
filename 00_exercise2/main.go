package main

import "fmt"

func main(){

	half := func (n int) (int, bool) {
		return n/2, n%2 == 0
	}

	fmt.Println(half(6))
	fmt.Println(largest(2,4,5,1,13,6,2,10))

	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))
	fmt.Println(foo())
}

func largest(x ...int) int {
	fmt.Printf("%T\n", x)
	var biggest int
	for _, w := range x {
		if w > biggest {
			biggest = w
		}
	}
	return biggest
}

func foo (numbers ...int) float64 {
	if numbers == nil {
		return 0
	}
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return float64(sum / len(numbers))

}
