package main

import (
	"github.com/shapled/goqml"
)

// @goqml
type Contact struct {
	goqml.QObject

	// @goqml.property("name")
	Name string
}

func NewContact() *Contact {
	contact := &Contact{Name: "InitialName"}
	contact.Setup(contact, contact.StaticMetaObject())
	return contact
}
