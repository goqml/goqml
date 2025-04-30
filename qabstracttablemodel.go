package goqml

import "unsafe"

var qAbstractTableModelStaticMetaObjectInstance = NewQAbstractTableModelMetaObject()

type QAbstractTableModel struct {
	QAbstractItemModel
}

func (model *QAbstractTableModel) StaticMetaObject() *QMetaObject {
	return qAbstractTableModelStaticMetaObjectInstance
}

func (model *QAbstractTableModel) Parent(child *QModelIndex) *QModelIndex {
	index := dos.QAbstractTableModelParent(DosQAbstractTableModel(model.vptr), child.vptr)
	return NewQModelIndexFromOther(index, OwnershipTake)
}

func (model *QAbstractTableModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	index := dos.QAbstractTableModelIndex(DosQAbstractTableModel(model.vptr), row, column, parent.vptr)
	return NewQModelIndexFromOther(index, OwnershipTake)
}

func (model *QAbstractTableModel) Setup() {
	qAIMCallbacks := DosQAbstractItemModelCallbacks{
		RowCount:     DosRowCountCallback(qModelRowCountCallback),
		ColumnCount:  DosColumnCountCallback(qModelColumnCountCallback),
		Data:         DosDataCallback(qModelDataCallback),
		SetData:      DosSetDataCallback(qModelSetDataCallback),
		RoleNames:    DosRoleNamesCallback(qModelRoleNamesCallback),
		Flags:        DosFlagsCallback(qModelFlagsCallback),
		HeaderData:   DosHeaderDataCallback(qModelHeaderDataCallback),
		Index:        DosIndexCallback(qModelIndexCallback),
		Parent:       DosParentCallback(qModelParentCallback),
		HasChildren:  DosHasChildrenCallback(qModelHasChildrenCallback),
		CanFetchMore: DosCanFetchMoreCallback(qModelCanFetchMoreCallback),
		FetchMore:    DosFetchMoreCallback(qModelFetchMoreCallback),
	}
	model.vptr = DosQObject(dos.QAbstractTableModelCreate(unsafe.Pointer(model), model.StaticMetaObject().vptr, qObjectCallback, qAIMCallbacks))
}
