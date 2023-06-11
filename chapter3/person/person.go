package person

import (
	"fmt"
	"time"
)

type Person struct {
	Dob       time.Time
	Email     string
	FirstName string
	LastName  string
	Location  string
}

func (p *Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

func (p *Person) PrintDetails() {
	fmt.Printf("[Date of Brith: %s, Email: %s, Location: %s]\n", p.Dob.String(), p.Email, p.Location)
}

func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}
