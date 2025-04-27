package main

import (
	"github.com/shapled/goqml"
)

type Contact struct {
	goqml.QObject[*Contact]

	Name string

	// @goqml.property("name").emitter("nameChanged")
	nameChanged func(string)
}

// @goqml.property("name").getter("name")
func (c *Contact) GetName() string {
	return c.Name
}

// @goqml.property("name").setter("setName")
func (c *Contact) SetName(name string) {
	if name != c.Name {
		c.Name = name
		c.nameChanged(name)
	}
}

func NewContact() *Contact {
	contact := &Contact{Name: ""}
	contact.Setup(contact, contact.StaticMetaObject())
	return contact
}
