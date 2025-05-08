package goqml

import (
	"unsafe"
)

type QModelIndex struct {
	vptr DosQModelIndex
}

func NewQModelIndex() *QModelIndex {
	index := &QModelIndex{}
	index.Setup()
	return index
}

func NewQModelIndexFromOther(other DosQModelIndex, takeOwnership Ownership) *QModelIndex {
	index := &QModelIndex{}
	index.SetupFromOther(other, takeOwnership)
	return index
}

func (index *QModelIndex) Setup() {
	index.vptr = DosQModelIndexCreate()
}

func (index *QModelIndex) SetupFromOther(other DosQModelIndex, takeOwnership Ownership) {
	switch takeOwnership {
	case OwnershipTake:
		index.vptr = other
	case OwnershipClone:
		index.vptr = DosQModelIndexCreateQModelIndex(other)
	default:
		panic("invalid ownership")
	}
}

func (index *QModelIndex) Row() int {
	return DosQModelIndexRow(index.vptr)
}

func (index *QModelIndex) Column() int {
	return DosQModelIndexColumn(index.vptr)
}

func (index *QModelIndex) IsValid() bool {
	return DosQModelIndexIsValid(index.vptr)
}

func (index *QModelIndex) Data(role int) *QVariant {
	return NewQVariantFrom(DosQModelIndexData(index.vptr, role), OwnershipTake)
}

func (index *QModelIndex) Parent() *QModelIndex {
	return NewQModelIndexFromOther(DosQModelIndexParent(index.vptr), OwnershipTake)
}

func (index *QModelIndex) Child(row int, column int) *QModelIndex {
	return NewQModelIndexFromOther(DosQModelIndexChild(index.vptr, row, column), OwnershipTake)
}

func (index *QModelIndex) Sibling(row int, column int) *QModelIndex {
	return NewQModelIndexFromOther(DosQModelIndexSibling(index.vptr, row, column), OwnershipTake)
}

func (index *QModelIndex) InternalPtr() unsafe.Pointer {
	return DosQModelIndexInternalPointer(index.vptr)
}
