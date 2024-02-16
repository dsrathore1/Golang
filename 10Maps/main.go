package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	lang := make(map[int]string)

	lang[1] = "JS"
	lang[2] = "PY"
	lang[3] = "TS"
	lang[4] = "JAVA"

	fmt.Println(lang)

	fmt.Println("Another type of declaration")

	name := map[int]string{
		11 : "Rahul",
		12 : "Joe",
		13 : "Kevin",
		14 : "Roman",
		15 : "John",
	}
	fmt.Println(name)
}
