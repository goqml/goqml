package main

import (
	"github.com/shapled/goqml"
)

type Contact struct {
	goqml.QObject
	Name string
}

func NewContact() *Contact {
	contact := &Contact{Name: "InitialName"}
	contact.Setup(goqml.NewQMetaObject(
		goqml.RootMetaObject,
		"Contact",
		nil,
		nil,
		[]*goqml.PropertyDefinition{
			goqml.NewPropertyDefinition("name", contact.getNameSlot, contact.setNameSlot, contact.nameChangedSignal),
		},
	))
	return contact
}

func (contact *Contact) getNameSlot() string {
	return contact.Name
}

func (contact *Contact) setNameSlot(name string) {
	if contact.Name == name {
		return
	}
	contact.Name = name
	contact.nameChangedSignal(name)
}

func (contact *Contact) nameChangedSignal(name string) {
	contact.Emit("notify_name", goqml.NewQVariantString(name))
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
