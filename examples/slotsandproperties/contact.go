package main

import (
	"github.com/shapled/goqml"
)

type Contact struct {
	goqml.QObject[*Contact]

	// @goqml.property("name")
	Name string
}

func NewContact() *Contact {
	contact := &Contact{Name: "InitialName"}
	contact.Setup(contact, contact.StaticMetaObject())
	return contact
}
