package main

import "fmt"
import "math"

// Square type with side
// Implements Shape interface
type Square struct {
	side float64
}

// Adding area func to the type Square
func (z Square) area() float64 {
	return z.side * z.side
}

// Circle type with radius
// Implements Shape interface
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Shape interface
// For Square, Circle
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

type Person struct {
	name string
	age  int
}

type Dog struct {
	name       string
	age        int
	tailLength int
}

type Cat struct {
	name       string
	age        int
	tailLength int
}

func (person Person) say() string {
	return "I'm human, I can talk"
}

func (dog Dog) say() string {
	return "Woof!"
}

func (cat Cat) say() string {
	return "Meow..."
}

type LivingCreature interface {
	say() string
}

func say(lc LivingCreature) {
	fmt.Println(lc.say())
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)

	person := Person{"Boss", 69}
	cat := Cat{"Kitty", 2, 13}
	dog := Dog{"Doggy", 3, 15}

	say(person)
	say(dog)
	say(cat)
}
