package main

import "fmt"

func main() {

	var age int
	fmt.Println("Enter your age")

	count, err := fmt.Scanf("%d", &age)

	fmt.Println(count)
	fmt.Println(err)
	fmt.Println(age)
}
