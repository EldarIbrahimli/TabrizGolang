package main

import "fmt"

func main() {
	var num1, num2, num3, sum int
	fmt.Println("Write three numbers")
	fmt.Scan(&num1, &num2, &num3)
	sum = num1 + num2 + num3
	fmt.Println(sum)

}
