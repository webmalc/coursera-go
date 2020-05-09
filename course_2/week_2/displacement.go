package main

import (
	"fmt"
)

// PromptForNum gets number from the user input.
func PromptForNum(name string) float64 {
	var num float64
	fmt.Printf("Enter %s: ", name)
	_, err := fmt.Scanln(&num)
	if err != nil {
		panic("Invalid number.")
	}
	return num
}

// GenDisplaceFn generates the displacement function.
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
}

func main() {
	a := PromptForNum("acceleration")
	v := PromptForNum("velocity")
	s := PromptForNum("displacement")
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(PromptForNum("time")))
}
