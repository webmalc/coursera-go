package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PromptForNums() []int {
	nums := make([]int, 0, 10)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the numbers: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("An error has occurred!")
	}
	for _, s := range strings.Fields(text) {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic("Invalid number!")
		}
		nums = append(nums, i)
	}
	return nums
}

// BubbleSort sort the slice
func BubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		swapped := false
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// Swap swaps the neighboring elements
func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}

func main() {
	nums := PromptForNums()
	BubbleSort(nums)
	fmt.Println(nums)
}
