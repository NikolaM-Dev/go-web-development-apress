package person

import "fmt"

type Admin struct {
	Person
	Roles []string
}

func (a Admin) PrintDetails() {
	a.Person.PrintDetails()

	fmt.Println("Admin Roles:")
	for _, r := range a.Roles {
		fmt.Println(r)
	}
}
