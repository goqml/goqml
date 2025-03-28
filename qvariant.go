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

func (qvar *QVariant) Delete() {
	dos.QVariantDelete(qvar.vptr)
}
