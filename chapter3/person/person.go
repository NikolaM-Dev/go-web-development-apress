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
