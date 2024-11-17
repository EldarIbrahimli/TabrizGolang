package main

import (
	"fmt"
)

func main() {
	for {
		var num1 float64
		var theSign string
		var num2 float64
		var startOver string

		fmt.Println("\nWelcome to the calculator!") // welcome

		fmt.Println("Write the first number") // the first number is written
		fmt.Scan(&num1)

		fmt.Println("What sign do you choose? (+, -, /, *, %)?") // the sign is written
		fmt.Scan(&theSign)

		fmt.Println("Write the second number") // the second number is written
		fmt.Scan(&num2)

		switch theSign {
		case "+":
			fmt.Println("The answer is: \n", num1+num2) // adds numbers together

		case "-":
			fmt.Println("The answer is: \n", num1-num2) // subtracts the numbers

		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.") // division by zero gives an error
			} else {
				fmt.Printf("The answer is: %.2f\n", num1/num2) // divides the numbers
			}

		case "*":
			fmt.Println("The answer is: \n", num1*num2) // multiplies the numbers

		case "%":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.") // division by zero gives an error
			} else {
				percentage := (num1 / num2) * 100
				fmt.Printf("The percentage is: %.2f%%\n", percentage) // shows the num1 percentage of num2
			}

		default:
			fmt.Println("Invalid sign. Please choose one of +, -, /, *, %.") // if the sign is not valid
		}

		fmt.Println("Thank you for using our calculator! \nDo you want to start over? (yes or no):") // you can start over by typing yes
		fmt.Scan(&startOver)

		if startOver != "yes" {
			fmt.Println("Thank you for using the calculator. Goodbye!") // if anything other than yes is typed, the code ends
			break
		}
	}
}
