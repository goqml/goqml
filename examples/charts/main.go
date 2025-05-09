package main

import "github.com/shapled/goqml"

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
	engine.Load("examples/charts/main.qml")
	app.Exec()
}
