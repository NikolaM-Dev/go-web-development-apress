package main

import (
	"time"

	"github.com/NikolaM-Dev/go-web-development-apress/chapter3/person"
	"github.com/NikolaM-Dev/go-web-development-apress/chapter3/team"
)

func main() {
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

	chris := &person.Member{
		Person: person.Person{
			FirstName: "Chris",
			LastName:  "Martin",
			Dob:       time.Date(1978, time.March, 15, 0, 0, 0, 0, time.UTC),
			Email:     "chris@email.com",
			Location:  "Santa Clara",
		},
		Skills: []string{
			"Go",
			"Docker",
		},
	}

	team := team.Team{
		Name:        "Go",
		Description: "Golang CoE",
		Users:       []person.User{alex, shiju, chris},
	}

	team.GetTeamDetails()
}
