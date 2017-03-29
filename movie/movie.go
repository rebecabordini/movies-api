package movie

import "fmt"

type Movie struct {
	Title string
	Director Person
	Year string
	Stars []Person
}

type Person struct {
	Name string
	Age int
}

func main() {
	p:= Person{
		Name: "Sofia Copolla",
		Age: 42,
	}
	p1:= Person{
		Name: "Scarlett Johansson",
		Age: 33,
	}
	p2:= Person{
		Name: "Rebecca Hall",
		Age: 35,
	}
	m:= Movie{
		Title: "Vicky Cristina Barcelona",
		Director: p,
		Year: "2005",
		Stars: []Person{p1, p2},
	}
	fmt.Println(m.Stars)
	fmt.Println(m.Director)
}
