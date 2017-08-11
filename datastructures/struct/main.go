package main

import "fmt"

type Person struct {
	Fname string
	Lname string
	Age   int
}

type Licence struct {
	Person
	Valid          bool
	ExpirationDate string
}

func (p Person) fullName() string {
	return p.Fname + p.Lname
}

func (l Licence) changeValidation() bool {
	if l.Valid {
		return false
	}
	return true

}

func main() {
	person1 := Person{"Boss", "Boss", 69}
	fmt.Println(person1.Fname, person1.Lname, person1.Age)

	personWithLicence := Licence{
		Person: Person{
			Fname: "Boss",
			Lname: "Boss",
			Age:   69,
		},
		Valid:          true,
		ExpirationDate: "2069-01",
	}
	fmt.Println(personWithLicence.fullName(), personWithLicence.ExpirationDate, personWithLicence.Valid)
	personWithLicence.Valid = personWithLicence.changeValidation()
	fmt.Println(personWithLicence.Valid)
}

func test() {

}
