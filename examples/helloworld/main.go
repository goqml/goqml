package main

import (
	"github.com/goqml/goqml"
)

func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	engine.Load("examples/helloworld/main.qml")
	app.Exec()
}
