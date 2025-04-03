package goqml

type QVariant struct {
	vptr DosQVariant
}

func NewQVariantInt(v int) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateInt(v)}
}

func NewQVariantString(v string) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateString(v)}
}

func NewQVariantBool(v bool) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateBool(v)}
}

func NewQVariantFloat(v float32) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateFloat(v)}
}

func NewQVariantDouble(v float64) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateDouble(v)}
}

func NewQVariantFrom(value DosQVariant, takeOwnership Ownership) *QVariant {
	switch takeOwnership {
	case OwnershipTake:
		return &QVariant{vptr: value}
	case OwnershipClone:
		return &QVariant{vptr: dos.QVariantCreateQVariant(value)}
	default:
		panic("invalid ownership")
	}
}

func NewQVariantQObject(obj IQObject) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateQObject(obj.qObjectVPtr())}
}

func (qvar *QVariant) Delete() {
	dos.QVariantDelete(qvar.vptr)
}

func (qvar *QVariant) StringVal() string {
	ptr := dos.QVariantToString(qvar.vptr)
	defer dos.CharArrayDelete(ptr)
	return charPtrToString(ptr)
}
