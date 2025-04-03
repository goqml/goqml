package main

import "github.com/shapled/goqml"

func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	contact := NewContact()
	defer contact.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	variant := goqml.NewQVariantQObject(contact)
	defer variant.Delete()

	engine.SetRootContextProperty("contact", variant)
	engine.Load("examples/slotsandproperties/main.qml")
	app.Exec()
}
