package goqml

type QVariant struct {
	vptr DosQVariant
}

func NewIntQVariant(v int) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateInt(v)}
}

func NewStringQVariant(v string) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateString(v)}
}

func NewBoolQVariant(v bool) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateBool(v)}
}

func NewFloatQVariant(v float32) *QVariant {
	return &QVariant{vptr: dos.QVariantCreateFloat(v)}
}

func (qvar *QVariant) Delete() {
	dos.QVariantDelete(qvar.vptr)
}
