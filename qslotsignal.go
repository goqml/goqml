package goqml

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/ebitengine/purego"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type ConnectionType int

const (
	ConnectionTypeAuto           ConnectionType = 0
	ConnectionTypeDirect         ConnectionType = 1
	ConnectionTypeQueued         ConnectionType = 2
	ConnectionTypeBlockingQueued ConnectionType = 3
	ConnectionTypeUnique         ConnectionType = 0x80
)

type ParameterDefinition struct {
	Name     string
	MetaType QMetaType
}

func (d *ParameterDefinition) toDos() DosParameterDefinition {
	return DosParameterDefinition{name: stringToCharPtr(d.Name), metaType: int32(d.MetaType)}
}

type SlotDefinition struct {
	Name        string
	RetMetaType QMetaType
	Params      []*ParameterDefinition
}

func (d *SlotDefinition) toDos() DosSlotDefinition {
	parameters := make([]DosParameterDefinition, len(d.Params))
	for i, param := range d.Params {
		parameters[i] = param.toDos()
	}
	return DosSlotDefinition{
		name:            stringToCharPtr(d.Name),
		returnMetaType:  int32(d.RetMetaType),
		parametersCount: int32(len(parameters)),
		parameters:      sliceToPtr(parameters),
	}
}

type SignalDefinition struct {
	Name   string
	Params []*ParameterDefinition
}

func (d *SignalDefinition) toDos() DosSignalDefinition {
	parameters := make([]DosParameterDefinition, len(d.Params))
	for i, param := range d.Params {
		parameters[i] = param.toDos()
	}
	return DosSignalDefinition{
		name:            stringToCharPtr(d.Name),
		parametersCount: int32(len(parameters)),
		parameters:      sliceToPtr(parameters),
	}
}

type PropertyDefinition struct {
	Name     string
	MetaType QMetaType
	Getter   string
	Setter   string
	Emitter  string
}

func (d *PropertyDefinition) toDos() DosPropertyDefinition {
	return DosPropertyDefinition{
		name:             stringToCharPtr(d.Name),
		propertyMetaType: int32(d.MetaType),
		readSlot:         stringToCharPtr(d.Getter),
		writeSlot:        stringToCharPtr(d.Setter),
		notifySignal:     stringToCharPtr(d.Emitter),
	}
}

func toQVariantSequence(qs DosQVariantArray, length int, takeOwnership Ownership) []*QVariant {
	var result []*QVariant
	qSlice := unsafe.Slice((*uintptr)(qs), length)
	for i := 0; i < length; i++ {
		result = append(result, NewQVariantFrom(DosQVariant(qSlice[i]), takeOwnership))
	}
	return result
}

func qobjectCallback[T IQObjectReal](_ purego.CDecl, ptr unsafe.Pointer, slotNamePtr DosQVariant, dosArgumentsLength int, dosArguments DosQVariantArray) uintptr {
	obj := *(*T)(ptr)

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
	return 0
}

var qObjectCallbackCache cmap.ConcurrentMap[uintptr, uintptr] = cmap.NewWithCustomShardingFunction[uintptr, uintptr](func(key uintptr) uint32 {
	return uint32(key)
})

func getQObjectCallback[T IQObjectReal]() uintptr {
	funcValue := reflect.ValueOf(qobjectCallback[T])
	funcPtr := funcValue.Pointer()

	if cb, ok := qObjectCallbackCache.Get(funcPtr); ok {
		return cb
	}

	cb := purego.NewCallback(qobjectCallback[T])
	qObjectCallbackCache.Set(funcPtr, cb)
	return cb
}

func Connect[S any, Sender IQObjectPtr[S], R any, Recevier IQObjectPtr[R]](
	sender Sender,
	signalName string,
	receiver Recevier,
	slotName string,
) *QMetaObjectConnection {
	return ConnectWithType(sender, signalName, receiver, slotName, ConnectionTypeAuto)
}

func ConnectWithType[S any, Sender IQObjectPtr[S], R any, Recevier IQObjectPtr[R]](
	sender Sender,
	signalName string,
	receiver Recevier,
	slotName string,
	connectionType ConnectionType,
) *QMetaObjectConnection {
	vptr := dos.QObjectConnectStatic(sender.getVPtr(), signalName, receiver.getVPtr(), slotName, int32(connectionType))
	return NewQMetaObjectConnection(vptr)
}

var qLambdaCallbackCache cmap.ConcurrentMap[uintptr, any] = cmap.NewWithCustomShardingFunction[uintptr, any](func(key uintptr) uint32 {
	return uint32(key)
})

var qLambdaCallback = purego.NewCallback(func(_ purego.CDecl, ptr uintptr, dosArgumentsLength int, dosArguments DosQVariantArray) uintptr {
	arguments := toQVariantSequence(dosArguments, dosArgumentsLength, OwnershipClone)
	defer func() {
		for _, qvar := range arguments {
			qvar.Delete()
		}
	}()

	if fn, ok := qLambdaCallbackCache.Get(ptr); ok {
		ApplyAndAssignQVariants(fn, append([]*QVariant{nil}, arguments...))
	}
	return 0
})

func ConnectFunc[S any, Sender IQObjectPtr[S]](sender Sender, signalName string, fn any) *QMetaObjectConnection {
	return ConnectFuncWithType(sender, signalName, fn, ConnectionTypeAuto)
}

func ConnectFuncWithType[S any, Sender IQObjectPtr[S]](sender Sender, signalName string, fn any, connectionType ConnectionType) *QMetaObjectConnection {
	funcType := reflect.TypeOf(fn)
	if funcType.Kind() != reflect.Func {
		panic("not a function")
	}
	if funcType.NumOut() != 0 {
		panic("only zero return value is supported")
	}

	funcValue := reflect.ValueOf(fn)
	funcPtr := funcValue.Pointer()

	qLambdaCallbackCache.Set(funcPtr, fn)
	vptr := dos.QObjectConnectLambdaStatic(sender.getVPtr(), signalName, DosQObjectConnectLambdaCallback(qLambdaCallback), funcPtr, int32(connectionType))
	return NewQMetaObjectConnection(vptr)
}

func MakeSignal(name string, paramTypes ...string) string {
	return fmt.Sprintf("2%s(%s)", name, strings.Join(paramTypes, ","))
}

func MakeSlot(name string, paramTypes ...string) string {
	return fmt.Sprintf("1%s(%s)", name, strings.Join(paramTypes, ","))
}
