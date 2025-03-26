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