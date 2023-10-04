package main

import (
	exported "Exported"
	"fmt"
)

// variadic function
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println(total)
}

// closure is a function that references var outside of its scope
func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursion - function calls itself
func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

func main() {
	fmt.Println("variadic")
	sum(1, 2, 3)
	a := []int{1, 2, 3, 4, 5, 6}
	sum(a...)

	next := closure()
	fmt.Println("closure")
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	new := closure()
	fmt.Println(new())

	fmt.Println("recursion")
	fmt.Println(recursive(7))

	//recursion implemented in closure
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println("recursion in closure")
	fmt.Println(fib(7))

	//Exported function call
	fmt.Println("exported funtion")
	exported.Hello("Naruto")

}
