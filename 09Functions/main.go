package main

import (
	"fmt"
)

func main() {
	fmt.Println("🎇Welcome to Functions Topic🎇")
	fmt.Println("\nAddition :-", addition(34, 12))      //* This func have "int" return type
	fmt.Println("\nSubtraction :-", subtraction(13, 5)) //* This func have "int" return type
	x := 5
	fmt.Println("\nPass by value :-", passByValue(x))

	_, m := ignoreValue()
	fmt.Println("\nIgnore one value by underscore getting second value (string type):-", m) //! Imp:- if want to skip one return type then just put on an underscore. Also declare two return types 1st is "int" and 2nd is "string"

	n, _ := ignoreValue()
	fmt.Println("\nReturn first value from func (int type) :-", n)

	l, j := namedType(12)
	fmt.Println("\nValue from name type function :- ", "\nThis is an addition of 12 :-", l, "\nThis is sub from 12 :-", j) //* Also know as naked function

}

// ! Function with return type "int" with two parameters
func subtraction(x int, y int) int {
	return x - y
}

// ! Not necessary to declare each parameter type if they all are belongs to same type, You just have to declare at the end of last parameter Golang will understand all the parameters are belonging to same type
func addition(a, b int) int {
	return (a + b)
}

// ! Variables in Go are passed by value (except for a few data types we haven't covered yet). "Pass by value" means that when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the caller's original data.
func passByValue(m int) int {
	return m
}

// ! A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _
func ignoreValue() (int, string) {
	return 3, "Hello"
}

// ! Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function. Named return values are best thought of as a way to document the purpose of the returned values.
func namedType(v int) (x, y int) {
	x = v + 4
	y = v - 4
	return
}
