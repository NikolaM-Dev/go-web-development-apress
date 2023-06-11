package main

import (
	"time"

	"github.com/NikolaM-Dev/go-web-development-apress/chapter3/person"
)

func main() {
	nikola := &person.Person{
		FirstName: "Nikola",
		LastName:  "Dev",
		Dob:       time.Date(2000, time.October, 7, 0, 0, 0, 0, time.UTC),
		Email:     "nikola@gmail.com",
		Location:  "Colombia",
	}

	nikola.PrintName()
	nikola.PrintDetails()

	nikola.ChangeLocation("New York")
	nikola.PrintDetails()

	alex := &person.Admin{
		Person: person.Person{
			Dob:       time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			Email:     "alex@gmail.com",
			FirstName: "Alex",
			LastName:  "John",
			Location:  "New York",
		},
		Roles: []string{
			"Manage Tasks",
			"Manage Team",
		},
	}

	shiju := &person.Member{
		Person: person.Person{
			Dob:       time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			Email:     "shiju@gmail.com",
			FirstName: "Shiju",
			LastName:  "Varghese",
			Location:  "Kochi",
		},
		Skills: []string{
			"Docker",
			"Go",
			"Kubernetes",
		},
	}

	alex.PrintName()
	alex.PrintDetails()

	shiju.PrintName()
	shiju.PrintDetails()
}
