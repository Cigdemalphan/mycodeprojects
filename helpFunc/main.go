package main

import "fmt"

func greet(name string)string{
	return "Hello,"+name+"!"

}

func main() {
	//Call the greet function and save its result
	message := greet("Jordon")
	//Print the returned message
	fmt.Println(message)
}