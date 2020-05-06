package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := make([]int, 0, 3)
	var input string
	for !strings.EqualFold(input, "x") {
		_, err := fmt.Scan(&input)
		if err == nil {
			k, err := strconv.Atoi(input)
			if err == nil {
				ints = append(ints, k)
				sort.Ints(ints)
				fmt.Println(ints)
			}
		}
	}
}
