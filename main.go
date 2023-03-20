package main

type Animal interface {
	walk() string
	getName() string
}

type Cat struct {
	Name   string
	Weight int
}

func (cat *Cat) walk() string {
	return "Cat is walking"
}

func (cat *Cat) getName() string {
	return "Cat's name is name"
}

func main() {
	cat := Cat{Name: "Cat", Weight: 15}

	PrintAnimal(&cat)
}

func PrintAnimal(animal Animal) {
	println(animal.getName())
}
