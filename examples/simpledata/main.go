package main

import "github.com/shapled/goqml"

func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	var qVar1 = goqml.NewIntQVariant(10)
	defer qVar1.Delete()

	var qVar2 = goqml.NewStringQVariant("Hello World")
	defer qVar2.Delete()

	var qVar3 = goqml.NewBoolQVariant(false)
	defer qVar3.Delete()

	var qVar4 = goqml.NewFloatQVariant(3.5)
	defer qVar4.Delete()

	engine.SetRootContextProperty("qVar1", qVar1)
	engine.SetRootContextProperty("qVar2", qVar2)
	engine.SetRootContextProperty("qVar3", qVar3)
	engine.SetRootContextProperty("qVar4", qVar4)
	engine.Load("examples/simpledata/main.qml")
	app.Exec()
}
