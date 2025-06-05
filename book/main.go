package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your favorite book name")
	name, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("My favorite book name is ", name)
	// fmt.Println(reflect.TypeOf(name))
	fmt.Print("Enter the book ISBN number")
	favNum, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("The book ISBN number is ", favNum)
	fmt.Print("Enter book fun fact")
	funFact, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your fun fact is ", funFact)
}
