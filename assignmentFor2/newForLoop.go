package main

import "fmt"

func main() {

	fmt.Println("Sums of items up to 5 are:")

	var sum int

	for i := 1; i <= 5; i++ {

		sum += i

	}
	fmt.Println(sum)

}
