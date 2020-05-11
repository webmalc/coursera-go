package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortNums(chunk []int, c chan []int) {
	fmt.Printf("Sorting in a goroutine: %v\n", chunk)
	sort.Ints(chunk)
	c <- chunk
}

func promptForNums() []int {
	nums := make([]int, 0, 12)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a series of numbers separated by space: ")
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

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func main() {
	const chunks = 4
	numbers := promptForNums()
	size := len(numbers) / chunks
	c := make(chan []int, chunks)

	for i := 0; i < chunks; i++ {
		if i != chunks-1 {
			go sortNums(numbers[i*size:(i+1)*size], c)
		} else {
			go sortNums(numbers[i*size:], c)
		}
	}
	parts := make([][]int, chunks)
	for i := 0; i < chunks; i++ {
		parts[i] = <-c
	}
	result := merge(merge(parts[0], parts[1]), merge(parts[2], parts[3]))
	fmt.Printf("Merged result: %v", result)
}
