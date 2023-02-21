package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greet: ")
	log.SetFlags(0)

	message, err := Greet("Jimbo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
