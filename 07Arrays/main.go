package main

import "fmt"

func main() {
	var helloMsg string = "\n🎇Hello Welcome to the golang Arrays concept🎇"
	fmt.Println(helloMsg)

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Custard-Apple"

	fmt.Println("\nYour list of fruits is", fruitList)
	fmt.Println("Size of the array fruitList is", len(fruitList))

	var numList = [5]int{3, 5, 1}

	fmt.Println("\nYour integer array is", numList)
	fmt.Print("Size of a numList array is ", len(numList))

	//! No matter how many value you input it will always shows maximum capacity of an array and there is no concept of dynamic array type, we always have to declare the size of and array before initialize an array
}
