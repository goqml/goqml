package goqml

type QVariant struct {
	vptr DosQVariant
}

func NewQVariantInt(v int) *QVariant {
	return &QVariant{vptr: DosQVariantCreateInt(v)}
}

func NewQVariantString(v string) *QVariant {
	return &QVariant{vptr: DosQVariantCreateString(v)}
}

func NewQVariantBool(v bool) *QVariant {
	return &QVariant{vptr: DosQVariantCreateBool(v)}
}

func NewQVariantFloat(v float32) *QVariant {
	return &QVariant{vptr: DosQVariantCreateFloat(v)}
}

func NewQVariantDouble(v float64) *QVariant {
	return &QVariant{vptr: DosQVariantCreateDouble(v)}
}

func NewQVariantFrom(value DosQVariant, takeOwnership Ownership) *QVariant {
	switch takeOwnership {
	case OwnershipTake:
		return &QVariant{vptr: value}
	case OwnershipClone:
		return &QVariant{vptr: DosQVariantCreateQVariant(value)}
	default:
		panic("invalid ownership")
	}
}

func NewQVariantQObject(obj IQObject) *QVariant {
	return &QVariant{vptr: DosQVariantCreateQObject(obj.getVPtr())}
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
	case *QObject:
		return NewQVariantQObject(value)
	default:
		panic("invalid type")
	}
}

func (qvar *QVariant) Delete() {
	DosQVariantDelete(qvar.vptr)
}

func (qvar *QVariant) StringVal() string {
	ptr := DosQVariantToString(qvar.vptr)
	defer DosCharArrayDelete(ptr)
	return charPtrToString(ptr)
}

func (qvar *QVariant) SetStringVal(value string) {
	DosQVariantSetString(qvar.vptr, value)
}

func (qvar *QVariant) IntVal() int {
	return DosQVariantToInt(qvar.vptr)
}

func (qvar *QVariant) SetIntVal(value int) {
	DosQVariantSetInt(qvar.vptr, value)
}

func (qvar *QVariant) BoolVal() bool {
	return DosQVariantToBool(qvar.vptr)
}

func (qvar *QVariant) SetBoolVal(value bool) {
	DosQVariantSetBool(qvar.vptr, value)
}

func (qvar *QVariant) FloatVal() float32 {
	return DosQVariantToFloat(qvar.vptr)
}

func (qvar *QVariant) SetFloatVal(value float32) {
	DosQVariantSetFloat(qvar.vptr, value)
}

func (qvar *QVariant) DoubleVal() float64 {
	return DosQVariantToDouble(qvar.vptr)
}

func (qvar *QVariant) SetDoubleVal(value float64) {
	DosQVariantSetDouble(qvar.vptr, value)
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
	default:
		panic("invalid type")
	}
}
