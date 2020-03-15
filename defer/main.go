package main

import "fmt"

func returnMessage() {
	fmt.Println("This is Simple Defer Function Execution")
}

func main() {
	defer returnMessage()
	fmt.Println("This is line One")
	fmt.Println("This is line Two")
	fmt.Println("This is line Three")
}
