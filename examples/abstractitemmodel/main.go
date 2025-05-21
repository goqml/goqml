package main

import "github.com/goqml/goqml"

//go:generate go run ../../cmd/goqml gen -f ./mylistmodel.go
func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	myListModel := NewMyListModel()
	defer myListModel.Delete()

	variant := goqml.NewQVariant(myListModel)
	defer variant.Delete()

	engine.SetRootContextProperty("myListModel", variant)
	engine.Load("examples/abstractitemmodel/main.qml")
	app.Exec()
}
