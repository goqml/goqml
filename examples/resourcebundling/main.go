package main

import (
	"fmt"
	"os"
	"path"

	"github.com/goqml/goqml"
)

//go:generate rcc --binary ./resources.qrc -o main.rcc
func main() {
	app := goqml.NewQApplication()
	defer app.Delete()
	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	dirname, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	goqml.RegisterResource(path.Join(dirname, "examples/resourcebundling/main.rcc"))
	engine.LoadUrl(goqml.NewQUrl("qrc:///main.qml"))

	app.Exec()
}
