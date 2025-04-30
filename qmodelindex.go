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
	index.vptr = dos.QModelIndexCreate()
}

func (index *QModelIndex) SetupFromOther(other DosQModelIndex, takeOwnership Ownership) {
	switch takeOwnership {
	case OwnershipTake:
		index.vptr = other
	case OwnershipClone:
		index.vptr = dos.QModelIndexCreateQModelIndex(other)
	default:
		panic("invalid ownership")
	}
}

func (index *QModelIndex) Row() int {
	return dos.QModelIndexRow(index.vptr)
}

func (index *QModelIndex) Column() int {
	return dos.QModelIndexColumn(index.vptr)
}

func (index *QModelIndex) IsValid() bool {
	return dos.QModelIndexIsValid(index.vptr)
}

func (index *QModelIndex) Data(role int) *QVariant {
	return NewQVariantFrom(dos.QModelIndexData(index.vptr, role), OwnershipTake)
}

func (index *QModelIndex) Parent() *QModelIndex {
	return NewQModelIndexFromOther(dos.QModelIndexParent(index.vptr), OwnershipTake)
}

func (index *QModelIndex) Child(row int, column int) *QModelIndex {
	return NewQModelIndexFromOther(dos.QModelIndexChild(index.vptr, row, column), OwnershipTake)
}

func (index *QModelIndex) Sibling(row int, column int) *QModelIndex {
	return NewQModelIndexFromOther(dos.QModelIndexSibling(index.vptr, row, column), OwnershipTake)
}

func (index *QModelIndex) InternalPtr() unsafe.Pointer {
	return dos.QModelIndexInternalPointer(index.vptr)
}
