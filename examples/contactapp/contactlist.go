package main

import "github.com/goqml/goqml"

const (
	FirstName = goqml.UserRole + 1
	LastName  = goqml.UserRole + 2
)

// @goqml
type ContactList struct {
	goqml.QAbstractListModel

	Contacts []*Contact
}

func NewContactList() *ContactList {
	cl := &ContactList{}
	cl.Setup(cl, cl.StaticMetaObject())
	return cl
}

func (cl *ContactList) Delete() {
	cl.QAbstractListModel.Delete()
	for _, contact := range cl.Contacts {
		contact.Delete()
	}
}

func (cl *ContactList) RowCount(index *goqml.QModelIndex) int {
	return len(cl.Contacts)
}

func (cl *ContactList) Data(index *goqml.QModelIndex, role int) *goqml.QVariant {
	if !index.IsValid() || index.Row() < 0 || index.Row() >= len(cl.Contacts) {
		return nil
	}
	contact := cl.Contacts[index.Row()]
	switch role {
	case FirstName:
		return goqml.NewQVariant(contact.FirstName)
	case LastName:
		return goqml.NewQVariant(contact.LastName)
	}
	return nil
}

func (cl *ContactList) RoleNames() map[int]string {
	return map[int]string{FirstName: "firstName", LastName: "lastName"}
}

// @goqml.slot("add")
func (cl *ContactList) Add(name string, surname string) {
	contact := NewContact()
	contact.FirstName = name
	contact.LastName = surname
	cl.BeginInsertRows(goqml.NewQModelIndex(), len(cl.Contacts), len(cl.Contacts))
	cl.Contacts = append(cl.Contacts, contact)
	cl.EndInsertRows()
}

// @goqml.slot("del")
func (cl *ContactList) Del(pos int) {
	if pos < 0 || pos >= len(cl.Contacts) {
		return
	}
	cl.BeginRemoveRows(goqml.NewQModelIndex(), pos, pos)
	cl.Contacts = append(cl.Contacts[:pos], cl.Contacts[pos+1:]...)
	cl.EndRemoveRows()
}
