package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	person := Person{
		Name: "alice",
		Age:  12,
	}
	computing(&person)
	fmt.Println(person)
}

func computing(person *Person) {
	person.Name = "john"
	(*person).Country = "Korea"
	computing2(person)
}
func computing2(person *Person) {
	person.Age = 80
}
