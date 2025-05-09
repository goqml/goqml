package main

import (
	"fmt"

	"github.com/shapled/goqml"
)

// @goqml
type ApplicationLogic struct {
	goqml.QObject

	ContactList *ContactList
	App         *goqml.QApplication
}

func NewApplicationLogic(app *goqml.QApplication) *ApplicationLogic {
	al := &ApplicationLogic{
		ContactList: NewContactList(),
		App:         app,
	}
	al.Setup(al, al.StaticMetaObject())
	return al
}

func (al *ApplicationLogic) Delete() {
	al.QObject.Delete()
	al.ContactList.Delete()
}

// @goqml.property("contactList").getter
func (al *ApplicationLogic) getContactList() *goqml.QVariant {
	return goqml.NewQVariant(al.ContactList)
}

// @goqml.slot("onLoadTriggered")
func (al *ApplicationLogic) OnLoadTriggered() {
	fmt.Println("Load Triggered")
	al.ContactList.Add("John", "Doo")
}

// @goqml.slot("onSaveTriggered")
func (al *ApplicationLogic) OnSaveTriggered() {
	fmt.Println("Save Triggered")
}

// @goqml.slot("onExitTriggered")
func (al *ApplicationLogic) OnExitTriggered() {
	al.App.Quit()
}
