package team

import (
	"fmt"

	"github.com/NikolaM-Dev/go-web-development-apress/chapter3/person"
)

type Team struct {
	Description string
	Name        string
	Users       []person.User
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members:")

	for _, u := range t.Users {
		u.PrintName()
		u.PrintDetails()
	}
}
