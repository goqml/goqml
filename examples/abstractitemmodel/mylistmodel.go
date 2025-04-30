package main

import "github.com/shapled/goqml"

const UserRoleName = goqml.UserRole + 1

type MyListModel struct {
	goqml.QAbstractListModel

	Names []string
}

func NewMyListModel() *MyListModel {
	model := &MyListModel{
		Names: []string{"John", "Max", "Paul", "Anna"},
	}
	model.Setup()
	return model
}

func (model *MyListModel) RowCount(index *goqml.QModelIndex) int {
	return len(model.Names)
}

func (model *MyListModel) Data(index *goqml.QModelIndex, role int) *goqml.QVariant {
	if !index.IsValid() {
		return nil
	}
	if index.Row() < 0 || index.Row() >= len(model.Names) {
		return nil
	}
	return goqml.NewQVariant(model.Names[index.Row()])
}

func (model *MyListModel) RoleNames() map[int]string {
	return map[int]string{UserRoleName: "name"}
}
