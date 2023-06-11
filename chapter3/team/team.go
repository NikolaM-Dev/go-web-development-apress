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
