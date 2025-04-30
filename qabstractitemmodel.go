package goqml

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

type QtItemFlag int

const (
	QtItemFlagNone             QtItemFlag = 0
	QtItemFlagIsSelectable     QtItemFlag = 1
	QtItemFlagIsEditable       QtItemFlag = 2
	QtItemFlagIsDragEnabled    QtItemFlag = 4
	QtItemFlagIsDropEnabled    QtItemFlag = 8
	QtItemFlagIsUserCheckable  QtItemFlag = 16
	QtItemFlagIsEnabled        QtItemFlag = 32
	QtItemFlagIsTristate       QtItemFlag = 64
	QtItemFlagNeverHasChildren QtItemFlag = 12

	UserRole = 0x100
)

var (
	rootAbstractItemModelMetaObject = NewQAbstractItemModelMetaObject()

	qModelRowCountCallback = purego.NewCallback(func(ptr unsafe.Pointer, rawIndex DosQModelIndex, result *int) {
		model := *(*IQAbstractItemModel)(ptr)
		index := NewQModelIndexFromOther(rawIndex, OwnershipClone)
		*result = model.RowCount(index)
	})

	qModelColumnCountCallback = purego.NewCallback(func(ptr unsafe.Pointer, rawIndex DosQModelIndex, result *int) {
		model := *(*IQAbstractItemModel)(ptr)
		index := NewQModelIndexFromOther(rawIndex, OwnershipClone)
		*result = model.ColumnCount(index)
	})

	qModelDataCallback = purego.NewCallback(func(ptr unsafe.Pointer, rawIndex DosQModelIndex, role int, result DosQVariant) {
		model := *(*IQAbstractItemModel)(ptr)
		index := NewQModelIndexFromOther(rawIndex, OwnershipClone)
		qvar := model.Data(index, role)
		if qvar != nil {
			dos.QVariantAssign(result, qvar.vptr)
			qvar.Delete()
		}
	})

	qModelSetDataCallback = purego.NewCallback(func(ptr unsafe.Pointer, rawIndex DosQModelIndex, rawValue DosQVariant, role int, result *bool) {
		model := *(*IQAbstractItemModel)(ptr)
		index := NewQModelIndexFromOther(rawIndex, OwnershipClone)
		qvar := NewQVariantFrom(rawValue, OwnershipClone)
		*result = model.SetData(index, qvar, role)
	})

	qModelRoleNamesCallback = purego.NewCallback(func(ptr unsafe.Pointer, result DosQHashIntByteArray) {
		model := *(*IQAbstractItemModel)(ptr)
		roleNames := model.RoleNames()
		for k, v := range roleNames {
			dos.QHashIntByteArrayInsert(result, k, v)
		}
	})

	qModelFlagsCallback = purego.NewCallback(func(ptr unsafe.Pointer, rawIndex DosQModelIndex, result *int) {
		model := *(*IQAbstractItemModel)(ptr)
		index := NewQModelIndexFromOther(rawIndex, OwnershipClone)
		*result = int(model.Flags(index))
	})

	qModelHeaderDataCallback = purego.NewCallback(func(ptr unsafe.Pointer, section int, orientation int, role int, result DosQVariant) {
		model := *(*IQAbstractItemModel)(ptr)
		qvar := model.HeaderData(section, orientation, role)
		if qvar != nil {
			dos.QVariantAssign(result, qvar.vptr)
			qvar.Delete()
		}
	})

	qModelIndexCallback = purego.NewCallback(func(ptr unsafe.Pointer, row int, column int, rawParent DosQModelIndex, result DosQModelIndex) {
		model := *(*IQAbstractItemModel)(ptr)
		parent := NewQModelIndexFromOther(rawParent, OwnershipClone)
		index := model.Index(row, column, parent)
		dos.QModelIndexAssign(result, index.vptr)
	})

	qModelParentCallback = purego.NewCallback(func(ptr unsafe.Pointer, childIndex DosQModelIndex, result DosQModelIndex) {
		model := *(*IQAbstractItemModel)(ptr)
		child := NewQModelIndexFromOther(childIndex, OwnershipClone)
		index := model.Parent(child)
		dos.QModelIndexAssign(result, index.vptr)
	})

	qModelHasChildrenCallback = purego.NewCallback(func(ptr unsafe.Pointer, parentIndex DosQModelIndex, result *bool) {
		model := *(*IQAbstractItemModel)(ptr)
		parent := NewQModelIndexFromOther(parentIndex, OwnershipClone)
		*result = model.HasChildren(parent)
	})

	qModelCanFetchMoreCallback = purego.NewCallback(func(ptr unsafe.Pointer, parentIndex DosQModelIndex, result *bool) {
		model := *(*IQAbstractItemModel)(ptr)
		parent := NewQModelIndexFromOther(parentIndex, OwnershipClone)
		*result = model.CanFetchMore(parent)
	})

	qModelFetchMoreCallback = purego.NewCallback(func(ptr unsafe.Pointer, parentIndex DosQModelIndex) {
		model := *(*IQAbstractItemModel)(ptr)
		index := NewQModelIndexFromOther(parentIndex, OwnershipClone)
		model.FetchMore(index)
	})
)

type IQAbstractItemModel interface {
	IQObject

	RowCount(index *QModelIndex) int
	ColumnCount(index *QModelIndex) int
	Data(index *QModelIndex, role int) *QVariant
	SetData(index *QModelIndex, value *QVariant, role int) bool
	RoleNames() map[int]string
	Flags(index *QModelIndex) QtItemFlag
	HeaderData(section int, orientation int, role int) *QVariant
	Index(row int, column int, parent *QModelIndex) *QModelIndex
	Parent(index *QModelIndex) *QModelIndex
	HasChildren(parent *QModelIndex) bool
	CanFetchMore(parent *QModelIndex) bool
	FetchMore(parent *QModelIndex)
	HasIndex(row int, column int, parent *QModelIndex) bool
	BeginInsertRows(parentIndex *QModelIndex, first int, last int)
	EndInsertRows()
	BeginRemoveRows(parentIndex *QModelIndex, first int, last int)
	EndRemoveRows()
	BeginInsertColumns(parentIndex *QModelIndex, first int, last int)
	EndInsertColumns()
	BeginRemoveColumns(parentIndex *QModelIndex, first int, last int)
	EndRemoveColumns()
	BeginResetModel()
	EndResetModel()
	DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)
}

type QAbstractItemModel struct {
	QObject
}

func (model *QAbstractItemModel) StaticMetaObject() *QMetaObject {
	return rootAbstractItemModelMetaObject
}

func (model *QAbstractItemModel) RowCount(index *QModelIndex) int {
	return 0
}

func (model *QAbstractItemModel) ColumnCount(index *QModelIndex) int {
	return 1
}

func (model *QAbstractItemModel) Data(index *QModelIndex, role int) *QVariant {
	return nil
}

func (model *QAbstractItemModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	return false
}

func (model *QAbstractItemModel) RoleNames() map[int]string {
	return nil
}

func (model *QAbstractItemModel) Flags(index *QModelIndex) QtItemFlag {
	return QtItemFlagNone
}

func (model *QAbstractItemModel) HeaderData(section int, orientation int, role int) *QVariant {
	return nil
}

func (model *QAbstractItemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	panic("not implemented")
}

func (model *QAbstractItemModel) Parent(index *QModelIndex) *QModelIndex {
	panic("not implemented")
}

func (model *QAbstractItemModel) HasChildren(parent *QModelIndex) bool {
	return dos.QAbstractItemModelHasChildren(DosQAbstractItemModel(model.vptr), parent.vptr)
}

func (model *QAbstractItemModel) CanFetchMore(parent *QModelIndex) bool {
	return dos.QAbstractItemModelCanFetchMore(DosQAbstractItemModel(model.vptr), parent.vptr)
}

func (model *QAbstractItemModel) FetchMore(parent *QModelIndex) {
	dos.QAbstractItemModelFetchMore(DosQAbstractItemModel(model.vptr), parent.vptr)
}

func (model *QAbstractItemModel) OnSlotCalled(slotName string, arguments []*QVariant) {
	fmt.Println("ignore QAbstractItemModel slot:", slotName)
}

func (model *QAbstractItemModel) Setup() {
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
	model.vptr = DosQObject(dos.QAbstractItemModelCreate(unsafe.Pointer(model), model.StaticMetaObject().vptr, qObjectCallback, qAIMCallbacks))
}

func (model *QAbstractItemModel) HasIndex(row int, column int, parent *QModelIndex) bool {
	return dos.QAbstractItemModelHasIndex(DosQAbstractItemModel(model.vptr), row, column, parent.vptr)
}

func (model *QAbstractItemModel) BeginInsertRows(parentIndex *QModelIndex, first int, last int) {
	dos.QAbstractItemModelBeginInsertRows(DosQAbstractItemModel(model.vptr), parentIndex.vptr, first, last)
}

func (model *QAbstractItemModel) EndInsertRows() {
	dos.QAbstractItemModelEndInsertRows(DosQAbstractItemModel(model.vptr))
}

func (model *QAbstractItemModel) BeginRemoveRows(parentIndex *QModelIndex, first int, last int) {
	dos.QAbstractItemModelBeginRemoveRows(DosQAbstractItemModel(model.vptr), parentIndex.vptr, first, last)
}

func (model *QAbstractItemModel) EndRemoveRows() {
	dos.QAbstractItemModelEndRemoveRows(DosQAbstractItemModel(model.vptr))
}

func (model *QAbstractItemModel) BeginInsertColumns(parentIndex *QModelIndex, first int, last int) {
	dos.QAbstractItemModelBeginInsertColumns(DosQAbstractItemModel(model.vptr), parentIndex.vptr, first, last)
}

func (model *QAbstractItemModel) EndInsertColumns() {
	dos.QAbstractItemModelEndInsertColumns(DosQAbstractItemModel(model.vptr))
}

func (model *QAbstractItemModel) BeginRemoveColumns(parentIndex *QModelIndex, first int, last int) {
	dos.QAbstractItemModelBeginRemoveColumns(DosQAbstractItemModel(model.vptr), parentIndex.vptr, first, last)
}

func (model *QAbstractItemModel) EndRemoveColumns() {
	dos.QAbstractItemModelEndRemoveColumns(DosQAbstractItemModel(model.vptr))
}

func (model *QAbstractItemModel) BeginResetModel() {
	dos.QAbstractItemModelBeginResetModel(DosQAbstractItemModel(model.vptr))
}

func (model *QAbstractItemModel) EndResetModel() {
	dos.QAbstractItemModelEndResetModel(DosQAbstractItemModel(model.vptr))
}

func (model *QAbstractItemModel) DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	dos.QAbstractItemModelDataChanged(DosQAbstractItemModel(model.vptr), topLeft.vptr, bottomRight.vptr, sliceToPtr(nil, roles), len(roles))
}
