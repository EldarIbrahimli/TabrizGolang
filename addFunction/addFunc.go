package main

import "fmt"

func sum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	var num1, num2 int
	fmt.Println("Enter two numbers for summation")
	fmt.Scan(&num1, &num2)
	//sum(num1, num2)
	fmt.Println("The summation of two numbers are:", sum(num1, num2))
}
