package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func prompt() (kind, method string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("An error has occurred!")
	}
	f := strings.Fields(strings.ToLower(text))

	return f[0], f[1]
}

// NewAnimal returns a new Animal object
func NewAnimal(kind string) *Animal {
	data := map[string]map[string]string{
		"cow":   {"food": "grass", "locomotion": "walk", "noise": "moo"},
		"bird":  {"food": "worms", "locomotion": "fly", "noise": "peep"},
		"snake": {"food": "mice", "locomotion": "slither", "noise": "hsss"},
	}
	if value, ok := data[kind]; ok {
		return &Animal{
			food:       value["food"],
			locomotion: value["locomotion"],
			noise:      value["noise"]}
	}
	panic("Invalid kind of animal!")
}

func main() {
	for {
		kind, method := prompt()
		animal := NewAnimal(kind)
		switch method {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			panic("Invalid action!")
		}
	}
}
