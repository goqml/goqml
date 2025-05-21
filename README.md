Developing, **not** ready for production.

## Preinstall

[DOtherSide](https://github.com/filcuc/DOtherSide)

## Install

```bash
go get github.com/goqml/goqml@latest
```

## Demo

main.go

```golang
package main

import (
	"github.com/goqml/goqml"
)

func main() {
	app := goqml.NewQApplication()
	defer app.Delete()

	engine := goqml.NewQQmlApplicationEngine()
	defer engine.Delete()

	engine.Load("./main.qml")
	app.Exec()
}
```

main.qml

```qml
import QtQuick 2.15
import QtQuick.Window 2.15

Window {
    width: 400
    height: 200
    visible: true
    title: "QML Hello World"

    Rectangle {
        anchors.centerIn:  parent
        width: 200
        height: 50
        color: "lightblue"

        Text {
            anchors.centerIn:  parent
            text: "Hello World!"
            font.pixelSize:  24
            color: "navy"
        }
    }
}
```

### More examples

[examples](./examples/)
