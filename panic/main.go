package main

import "fmt"

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}

func executePanic() {
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely") // this line won't be executed
}
