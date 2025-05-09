package main

import "github.com/shapled/goqml"

//go:generate go run ../../cmd/goqml gen -f ./contact.go
//go:generate go run ../../cmd/goqml gen -f ./contactlist.go
//go:generate go run ../../cmd/goqml gen -f ./applicationlogin.go
func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	logic := NewApplicationLogic(app)
	defer logic.Delete()

	variant := goqml.NewQVariant(logic)
	defer variant.Delete()

	engine.SetRootContextProperty("logic", variant)
	engine.Load("examples/contactapp/main.qml")
	app.Exec()
}
