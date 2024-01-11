package main

import "fmt"

func main() {
	var helloMsg string = "\n🎇Hello Welcome to the golang pointer concept🎇"
	fmt.Println(helloMsg)

	var myNum int = 23

	var ptr = &myNum

	fmt.Println("\nValue of actual pointer is", ptr, "<-- This is an memory location where", myNum, "<-- is save")
	fmt.Println("\nVale of pointer is", *ptr, "<-- this is an value on that location where this pointer points")

	*ptr = *ptr + 2
	fmt.Println("\nNew value is :-", *ptr, "<-- we only manipulate into value of pointer and that will not change the referenced it will change real value as well")
}
