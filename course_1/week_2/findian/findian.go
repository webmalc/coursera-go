package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(text string) bool {
	const (
		first  = 'i'
		middle = 'a'
		last   = 'n'
	)
	text = strings.ToLower(strings.TrimSpace(text))
	l := len(text)

	if l == 0 {
		return false
	}
	if text[0] != first {
		return false
	}
	if text[l-1] != last {
		return false
	}

	return strings.ContainsRune(text, middle)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("An error has occurred!")
	}
	if check(text) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
