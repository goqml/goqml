package main

import (
	"fmt"

	"github.com/shapled/goqml"
)

type Contact struct {
	goqml.QObject
	Name string
}

func NewContact() *Contact {
	contact := &Contact{Name: "InitialName"}
	contact.Setup()
	return contact
}

func (contact *Contact) OnSlotCalled(slotName string, arguments []*goqml.QVariant) {
	fmt.Printf("slot called: %s, %#v", slotName, arguments)
}

//   proc getName*(self: Contact): string {.slot.} =
//     result = self.m_name

//   proc nameChanged*(self: Contact, name: string) {.signal.}

//   proc setName*(self: Contact, name: string) {.slot.} =
//     if self.m_name == name:
//       return
//     self.m_name = name
//     self.nameChanged(name)

//   QtProperty[string] name:
//     read = getName
//     write = setName
//     notify = nameChanged
