package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\nWelcome to time study of golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println("\nNow, you will get exact date, day and time -->", presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2021, time.April, 2, 12, 30, 32, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println("\nLet's format date by I want")
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}
