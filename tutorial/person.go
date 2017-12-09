package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstname, lastname, email, location string
	dob                                  time.Time
}

type Admin struct {
	Person
	Roles []string
}

type Member struct {
	Person
	Skills []string
}

func (p *Person) getFirstName() {
	fmt.Printf("\n%s is the name\n", p.firstname)
}

func (p *Person) setFirstname(newName string) {
	p.firstname = newName
}

func main() {
	var p Person
	p.firstname = "A_first_name"
	p.lastname = "A_last_name"
	p.email = "foo@bar"
	p.location = "Kenilworth"
	p.dob = time.Date(1908, time.February, 13, 0, 0, 0, 0, time.UTC)
	//p.setFirstname("George")
	var adminUser Admin
	adminUser.firstname = "Admin's Name"
	adminUser.Roles = []string{"role1,role2,role3"}
	adminUser.getFirstName()
	p.getFirstName()
}
