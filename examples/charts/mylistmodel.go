package main

import (
	"math/rand"

	"github.com/goqml/goqml"
)

type Point struct {
	X int
	Y int
}

// @goqml
type MyListModel struct {
	goqml.QAbstractListModel

	Points []*Point
	MaxX   int
	MaxY   int

	// @goqml.property("maxX").emitter
	MaxXChanged func(value int)

	// @goqml.property("maxY").emitter
	MaxYChanged func(value int)
}

func NewMyListModel() *MyListModel {
	model := &MyListModel{
		Points: []*Point{},
		MaxX:   0,
		MaxY:   50,
	}
	model.Setup(model, model.StaticMetaObject())
	model.AddRandomPoint()
	return model
}

func (model *MyListModel) RowCount(index *goqml.QModelIndex) int {
	return len(model.Points)
}

func (model *MyListModel) ColumnCount(index *goqml.QModelIndex) int {
	return 2
}

func (model *MyListModel) Data(index *goqml.QModelIndex, role int) *goqml.QVariant {
	if !index.IsValid() || index.Row() < 0 || index.Row() >= model.RowCount(nil) || index.Column() < 0 || index.Column() >= model.ColumnCount(nil) {
		return nil
	}
	if role == 0 {
		point := model.Points[index.Row()]
		if index.Column() == 0 {
			return goqml.NewQVariant(point.X)
		} else if index.Column() == 1 {
			return goqml.NewQVariant(point.Y)
		}
	}
	return nil
}

// @goqml.property("maxX").getter
func (model *MyListModel) GetMaxX() int {
	return model.MaxX
}

// @goqml.property("maxY").getter
func (model *MyListModel) GetMaxY() int {
	return model.MaxY
}

// @goqml.slot("addRandomPoint")
func (model *MyListModel) AddRandomPoint() {
	pos := len(model.Points)
	model.BeginInsertRows(goqml.NewQModelIndex(), pos, pos)
	x := model.MaxX + 1
	y := rand.Int() % 50
	if x > model.MaxX {
		model.MaxX = x
		model.MaxXChanged(x)
	}
	if y > model.MaxY {
		model.MaxY = y
		model.MaxYChanged(y)
	}
	model.Points = append(model.Points, &Point{X: x, Y: y})
	model.EndInsertRows()
}
