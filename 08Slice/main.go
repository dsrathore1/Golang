package main

import (
	"fmt"
	"sort"
)

func main() {
	var helloMsg string = "\n🎇Hello Welcome to the golang Slice concept🎇"
	fmt.Println(helloMsg)

	var fruitList = []string{"Apple", "Mango", "Peach"} //! This is a "slice" data structure, we have to initialize it during creation, with dynamic length

	fmt.Printf("Type of this list is %T\n", fruitList)
	fmt.Println("\nYour fruit list is :- ", fruitList)

	//! Append from the back of the "slice"
	fruitList = append(fruitList, "Tomato", "Banana", "Grapes")
	fmt.Println("\nYour updated fruit list is :- ", fruitList)

	//! Now, we print the value from given starting point with index number
	//? Starting point is 1, it means leave 0th index element and print all the rest of the element present into the slice
	fruitList = append(fruitList[1:])
	fmt.Println("Your list from starting index 1", fruitList) //? Exclusively, "apple" will not countable

	//! Here, we will give both the point start and end.
	fruitList = append(fruitList[:3]) //! Last range is always non inclusive
	fmt.Println(fruitList)

	//! Another way to make slice

	highScores := make([]int, 4)

	highScores[0] = 44
	highScores[1] = 32
	highScores[2] = 54
	highScores[3] = 50

	// highScores[4] = 56  //* This will give an error because slice range exit

	fmt.Println("Used all the indexed now result is", highScores, "we cannot add more elements into slice by traditional method")

	//! In this method we can dynamically addon elements without get effected by already declared range of slice. Cause append method will relocate all the elements to different memory location and add new elements into it and this time we can append lot of values into an slice.

	highScores = append(highScores, 33, 443, 85)
	fmt.Println("After append into slice the final result is :-", highScores)

	//! We can sort this slice with built in package sort
	sort.Ints(highScores)
	fmt.Println("After sort a slice in increasing order the result look like :-", highScores)

	fmt.Println("This will tell slice is sorted or not by showing boolean value :-", sort.IntsAreSorted(highScores))

}
