package goqml

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/shapled/goqml/util"
)

type (
	Ownership int
)

const (
	OwnershipTake  Ownership = 0
	OwnershipClone Ownership = 1
)

var (
	RootMetaObject = NewQObjectMetaObject()
)

type IQObject interface {
	getVPtr() DosQObject
	setVPtr(vptr DosQObject)
	setOwned(owned bool)
	StaticMetaObject() *QMetaObject
	OnSlotCalled(slotName string, arguments []*QVariant)
}

func ptrOfIQObjectReal(obj IQObject) unsafe.Pointer {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		panic("obj must be a pointer type")
	}
	return unsafe.Pointer(val.Pointer())
}

type IQObjectPtr[T any] interface {
	IQObject
	*T
}

type QObject struct {
	vptr  DosQObject
	owner bool
}

func (obj *QObject) StaticMetaObject() *QMetaObject {
	return RootMetaObject
}

func (obj *QObject) Setup(inst IQObject, meta *QMetaObject) {
	obj.owner = true
	obj.vptr = dos.QObjectCreate(unsafe.Pointer(&inst), meta.vptr, DosQObjectCallBack(qIObjectCallback))
	util.Pin(obj)
}

func (obj *QObject) Delete() {
	util.Unpin(obj)
	if obj.vptr == nil || !obj.owner {
		return
	}
	dos.QObjectDelete(obj.vptr)
	obj.vptr = nil
}

func (obj *QObject) getVPtr() DosQObject {
	return obj.vptr
}

func (obj *QObject) setVPtr(vptr DosQObject) {
	obj.vptr = vptr
}

func (obj *QObject) setOwned(owned bool) {
	obj.owner = owned
}

func (obj *QObject) Emit(signalName string, arguments ...*QVariant) {
	dosArguments := []DosQVariant{}
	for _, argument := range arguments {
		dosArguments = append(dosArguments, argument.vptr)
	}
	dos.QObjectSignalEmit(obj.vptr, signalName, len(dosArguments), DosQVariantArray(sliceToPtr(nil, dosArguments)))
}

func (obj *QObject) DeleteLater() {
	if !obj.owner || obj.vptr == nil {
		return
	}
	dos.QObjectDeleteLater(obj.vptr)
	obj.vptr = nil
}

func (obj *QObject) OnSlotCalled(slotName string, arguments []*QVariant) {
	fmt.Println("ignore QObject slot:", slotName)
}

func qObjectCallback[T interface {
	OnSlotCalled(slotName string, arguments []*QVariant)
}](obj T, slotNamePtr DosQVariant, dosArgumentsLength int, dosArguments DosQVariantArray) {
	slotName := NewQVariantFrom(slotNamePtr, OwnershipClone)
	defer slotName.Delete()

	arguments := toQVariantSequence(dosArguments, dosArgumentsLength, OwnershipClone)
	defer func() {
		for _, qvar := range arguments {
			qvar.Delete()
		}
	}()

	obj.OnSlotCalled(slotName.StringVal(), arguments)

	dosArgs := unsafe.Slice((*uintptr)(dosArguments), dosArgumentsLength)
	dos.QVariantAssign(DosQVariant(dosArgs[0]), arguments[0].vptr)
}

var qIObjectCallback = purego.NewCallback(func(_ purego.CDecl, ptr unsafe.Pointer, slotNamePtr DosQVariant, dosArgumentsLength int, dosArguments DosQVariantArray) uintptr {
	obj := *(*IQObject)(ptr)
	qObjectCallback(obj, slotNamePtr, dosArgumentsLength, dosArguments)
	return 0
})
