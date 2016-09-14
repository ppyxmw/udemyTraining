package main

import "fmt"

func main() {
	var name string
	fmt.Print("what is your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello dusty Mr", name)

	var small int
	var big int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Now enter a big number: ")
	fmt.Scan(&big)
	fmt.Println("The remainder of", big, "divided by", small, "equal"+
		" to", big%small)

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, "-- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "-- Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "-- Buzz")
		} else {
			fmt.Println(i)
		}
	}

	sum := 0

	for i := 0; i < 1000; i++ {

		if i%5 == 0 {
			sum += i
		} else if i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

}
