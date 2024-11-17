package main

import "fmt"

func arraySum(arr [100]int, size int) {
	sum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	fmt.Printf("Sum = %d\n", sum)
}
func arrayPrint(arr [100]int, size int) {
	fmt.Print("Your array: [")
	for i := 0; i < size; i++ {
		if i == size-1 {
			fmt.Printf("%d", arr[i])
		} else {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println("]")
}

func main() {
	var arr [100]int
	var size int

	fmt.Print("How many numbers do you want to enter (max 100)? ")
	fmt.Scan(&size)

	if size <= 0 || size > 100 {
		fmt.Println("Please enter a size between 1 and 100")
		return
	}

	for i := 0; i < size; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&arr[i])
	}
	arrayPrint(arr, size)
	arraySum(arr, size)
}
