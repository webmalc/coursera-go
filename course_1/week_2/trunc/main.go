package main

import "fmt"

func main() {
	var num float64
	fmt.Println("Enter a float number:")
	_, err := fmt.Scanln(&num)
	if err != nil {
		panic("Invalid number.")
	}
	fmt.Println(int(num))
}
