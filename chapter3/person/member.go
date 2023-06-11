package person

import "fmt"

type Member struct {
	Person
	Skills []string
}

func (m Member) PrintDetails() {
	m.Person.PrintDetails()

	fmt.Println("Skills:")

	for _, s := range m.Skills {
		fmt.Println(s)
	}
}
