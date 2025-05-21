package main

import (
	"github.com/goqml/goqml"
)

//go:generate go run ../../cmd/goqml gen -f ./contact.go
func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	goqml.QmlRegisterType("ContactModule", 1, 0, "Contact", func() *Contact { return NewContact() })

	engine.Load("examples/qmlregistertype/main.qml")
	app.Exec()
}
