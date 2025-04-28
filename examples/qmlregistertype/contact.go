package main

import "github.com/shapled/goqml"

type Contact struct {
	goqml.QObject

	// @goqml.property("firstName")
	FirstName string

	// @goqml.property("lastName")
	LastName string
}

func NewContact() *Contact {
	contact := &Contact{FirstName: "", LastName: ""}
	contact.Setup(contact, contact.StaticMetaObject())
	return contact
}
