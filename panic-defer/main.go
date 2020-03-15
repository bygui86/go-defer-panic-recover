package main

import "fmt"

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}

func executePanic() {
	defer recoveryFunction()
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely") // this line won't be executed
}

func recoveryFunction() {
	fmt.Println("This is recovery function...")
}
