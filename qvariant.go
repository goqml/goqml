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
	return &QVariant{vptr: dos.QVariantCreateQObject(obj.getVPtr())}
}

func NewQVariant(value any) *QVariant {
	switch value := value.(type) {
	case int:
		return NewQVariantInt(value)
	case string:
		return NewQVariantString(value)
	case bool:
		return NewQVariantBool(value)
	case float32:
		return NewQVariantFloat(value)
	case float64:
		return NewQVariantDouble(value)
	case IQObject:
		return NewQVariantQObject(value)
	default:
		panic("invalid type")
	}
}

func (qvar *QVariant) Delete() {
	dos.QVariantDelete(qvar.vptr)
}

func (qvar *QVariant) StringVal() string {
	ptr := dos.QVariantToString(qvar.vptr)
	defer dos.CharArrayDelete(ptr)
	return charPtrToString(ptr)
}

func (qvar *QVariant) SetStringVal(value string) {
	dos.QVariantSetString(qvar.vptr, value)
}

func (qvar *QVariant) IntVal() int {
	return dos.QVariantToInt(qvar.vptr)
}

func (qvar *QVariant) SetIntVal(value int) {
	dos.QVariantSetInt(qvar.vptr, value)
}

func (qvar *QVariant) BoolVal() bool {
	return dos.QVariantToBool(qvar.vptr)
}

func (qvar *QVariant) SetBoolVal(value bool) {
	dos.QVariantSetBool(qvar.vptr, value)
}

func (qvar *QVariant) FloatVal() float32 {
	return dos.QVariantToFloat(qvar.vptr)
}

func (qvar *QVariant) SetFloatVal(value float32) {
	dos.QVariantSetFloat(qvar.vptr, value)
}

func (qvar *QVariant) DoubleVal() float64 {
	return dos.QVariantToDouble(qvar.vptr)
}

func (qvar *QVariant) SetDoubleVal(value float64) {
	dos.QVariantSetDouble(qvar.vptr, value)
}

func (qvar *QVariant) SetQVariantVal(value *QVariant) {
	dos.QVariantAssign(qvar.vptr, value.vptr)
}

func (qvar *QVariant) SetVal(value any) {
	switch value := value.(type) {
	case int:
		qvar.SetIntVal(value)
	case string:
		qvar.SetStringVal(value)
	case bool:
		qvar.SetBoolVal(value)
	case float32:
		qvar.SetFloatVal(value)
	case float64:
		qvar.SetDoubleVal(value)
	case *QVariant:
		qvar.SetQVariantVal(value)
	default:
		panic("invalid type")
	}
}
