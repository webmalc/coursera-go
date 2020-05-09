package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func prompt() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">")
	text, err := reader.ReadString('\n')
	if err != nil {
		return []string{}, err
	}
	f := strings.Fields(strings.ToLower(text))

	return f, nil
}

// NewAnimal returns a new Animal object
func NewAnimal(kind string) (Animal, error) {
	switch kind {
	case "cow":
		return &Cow{}, nil
	case "bird":
		return &Bird{}, nil
	case "snake":
		return &Snake{}, nil
	default:
		return nil, errors.New("invalid kind of animal")
	}
}

func RunCommand(animal Animal, command string) {
	switch command {
	case "move":
		animal.Move()
	case "eat":
		animal.Eat()
	case "speak":
		animal.Speak()
	default:
		printError()
	}
}

func printError() {
	fmt.Println("Invalid command!")
}

func main() {
	storage := make(map[string]Animal)
	for {
		params, err := prompt()
		if err != nil || len(params) != 3 {
			printError()
			continue
		}
		switch params[0] {
		case "newanimal":
			animal, err := NewAnimal(params[2])
			if err != nil {
				printError()
			}
			storage[params[1]] = animal
			fmt.Println("Created it!")
		case "query":
			if animal, ok := storage[params[1]]; ok {
				RunCommand(animal, params[2])
			} else {
				printError()
			}
		default:
			printError()
		}
	}
}
