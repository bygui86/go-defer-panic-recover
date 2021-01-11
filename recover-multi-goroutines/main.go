package main

import "fmt"

func main() {
	executePanic()

	fmt.Println("Main function completed") // this line will be executed
}

func executePanic() {
	defer recoveryFunction()

	panic("This is Panic Situation")

	fmt.Println("Panic function completed") // this line won't be executed
}

func recoveryFunction() {
	recoveryMessage := recover()

	if recoveryMessage != nil {
		fmt.Printf("Panic situation avoided: %s \n", recoveryMessage)
	}

	fmt.Println("Recovery function completed") // this line will be executed
}
