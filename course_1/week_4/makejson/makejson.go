package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func prompt(name string, m map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s: ", name)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("An error has occurred!")
	}
	m[name] = strings.TrimSpace(text)
}

func main() {
	m := make(map[string]string)
	prompt("name", m)
	prompt("address", m)
	data, err := json.Marshal(m)
	if err != nil {
		panic("An error has occurred!")
	}
	fmt.Println(string(data))
}
