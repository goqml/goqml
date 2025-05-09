package main

import "github.com/shapled/goqml"

// @goqml
type Contact struct {
	goqml.QObject

	// @goqml.property("firstName")
	FirstName string

	// @goqml.property("surname")
	LastName string
}

func NewContact() *Contact {
	contact := &Contact{
		FirstName: "first",
		LastName:  "last",
	}
	contact.Setup(contact, contact.StaticMetaObject())
	return contact
}
