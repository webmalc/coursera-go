package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a filename: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("An error has occurred!")
	}
	return strings.TrimSpace(text)
}

func getNames(filename string) *[]name {
	result := make([]name, 0)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		result = append(result, name{fname: parts[0], lname: parts[1]})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &result
}

func printNames(names *[]name) {
	for _, name := range *names {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}
}

func main() {
	filename := prompt()
	names := getNames(filename)
	printNames(names)
}
