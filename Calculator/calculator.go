package main

import "fmt"

func main() {
	//did not do it myself could not understand by myself
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	sum := 0
	for number != 0 {
		sum += number % 10 // Add the last digit to sum
		number /= 10       // Remove the last digit
	}

	fmt.Println("Sum of digits:", sum)

	/*var num1, num2, num3, sum int
	fmt.Println("Write three numbers")
	fmt.Scan(&num1, &num2, &num3)
	sum = num1 + num2 + num3
	fmt.Println(sum)*/

}
