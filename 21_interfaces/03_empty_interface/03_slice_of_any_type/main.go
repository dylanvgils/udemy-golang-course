package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

// Function accepts value of any type
func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	animals := []interface{}{fido, fifi}
	fmt.Println(animals)
}
