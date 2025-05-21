package main

import "github.com/goqml/goqml"

//go:generate go run ../../cmd/goqml gen -f ./contact.go
func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	contact := NewContact()
	defer contact.Delete()

	variant := goqml.NewQVariantQObject(contact)
	defer variant.Delete()

	engine.SetRootContextProperty("contact", variant)
	engine.Load("examples/slotsandproperties/main.qml")
	app.Exec()
}
