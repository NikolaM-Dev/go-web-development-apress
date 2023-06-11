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
}
