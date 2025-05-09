package goqml

import "unsafe"

var qAbstractListModelStaticMetaObjectInstance = NewQAbstractListModelMetaObject()

type QAbstractListModel struct {
	QAbstractItemModel
}

func (model *QAbstractListModel) StaticMetaObject() *QMetaObject {
	return qAbstractListModelStaticMetaObjectInstance
}

func (model *QAbstractListModel) ColumnCount(index *QModelIndex) int {
	return dos.QAbstractListModelColumnCount(DosQAbstractListModel(model.vptr), index.vptr)
}

func (model *QAbstractListModel) Parent(child *QModelIndex) *QModelIndex {
	index := dos.QAbstractListModelParent(DosQAbstractListModel(model.vptr), child.vptr)
	return NewQModelIndexFromOther(index, OwnershipTake)
}

func (model *QAbstractListModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	index := dos.QAbstractListModelIndex(DosQAbstractListModel(model.vptr), row, column, parent.vptr)
	return NewQModelIndexFromOther(index, OwnershipTake)
}

func (model *QAbstractListModel) Setup(inst IQAbstractItemModel, meta *QMetaObject) {
	qAIMCallbacks := &DosQAbstractItemModelCallbacks{
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
	model.vptr = DosQObject(dos.QAbstractListModelCreate(unsafe.Pointer(&inst), meta.vptr, qIQAbstractItemModel, qAIMCallbacks))
}
