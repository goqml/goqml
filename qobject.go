package goqml

import (
	"fmt"
	"reflect"
	"unsafe"
)

type (
	Ownership int
)

const (
	OwnershipTake Ownership = iota
	OwnershipClone
)

var (
	RootMetaObject = NewQObjectMetaObject()
)

type IQObjectReal interface {
	getVPtr() DosQObject
	setVPtr(vptr DosQObject)
	setOwned(owned bool)
	StaticMetaObject() *QMetaObject
	OnSlotCalled(slotName string, arguments []*QVariant)
}

func ptrOfIQObjectReal(obj IQObjectReal) unsafe.Pointer {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		panic("obj must be a pointer type")
	}
	return unsafe.Pointer(val.Pointer())
}

type IQObjectPtr[T any] interface {
	IQObjectReal
	*T
}

type QObject[T IQObjectReal] struct {
	vptr  DosQObject
	owner bool
}

func (obj *QObject[T]) StaticMetaObject() *QMetaObject {
	return RootMetaObject
}

func (obj *QObject[T]) Setup(inst T, meta *QMetaObject) {
	obj.owner = true
	obj.vptr = dos.QObjectCreate(unsafe.Pointer(&inst), meta.vptr, DosQObjectCallBack(getQObjectCallback[T]()))
}

func (obj *QObject[T]) Delete() {
	if obj.vptr == nil || !obj.owner {
		return
	}
	dos.QObjectDelete(obj.vptr)
	obj.vptr = nil
}

func (obj *QObject[T]) getVPtr() DosQObject {
	return obj.vptr
}

func (obj *QObject[T]) setVPtr(vptr DosQObject) {
	obj.vptr = vptr
}

func (obj *QObject[T]) setOwned(owned bool) {
	obj.owner = owned
}

func (obj *QObject[T]) Emit(signalName string, arguments ...*QVariant) {
	dosArguments := []DosQVariant{}
	for _, argument := range arguments {
		dosArguments = append(dosArguments, argument.vptr)
	}
	dos.QObjectSignalEmit(obj.vptr, signalName, len(dosArguments), DosQVariantArray(sliceToPtr(dosArguments)))
}

func (obj *QObject[T]) DeleteLater() {
	if !obj.owner || obj.vptr == nil {
		return
	}
	dos.QObjectDeleteLater(obj.vptr)
	obj.vptr = nil
}

func (obj *QObject[T]) OnSlotCalled(slotName string, arguments []*QVariant) {
	fmt.Println("ignore QObject slot:", slotName)
}

// let qObjectStaticMetaObjectInstance = newQObjectMetaObject()

// proc staticMetaObject*(c: type QObject): QMetaObject =
//   ## Return the metaObject of QObject
//   qObjectStaticMetaObjectInstance

// proc staticMetaObject*(self: QObject): QMetaObject =
//   ## Return the metaObject of QObject
//   qObjectStaticMetaObjectInstance

// proc objectName*(self: QObject): string =
//   ## Return the QObject name
//   var str = dos_qobject_objectName(self.vptr)
//   result = $str
//   dos_chararray_delete(str)

// proc setObjectName*(self: QObject, name: string) =
//   ## Sets the Qobject name
//   dos_qobject_setObjectName(self.vptr, name.cstring)

// proc `objectName=`*(self: QObject, name: string) =
//   ## Sets the Qobject name
//   self.setObjectName(name)

// method metaObject*(self: QObject): QMetaObject {.base.} =
//   ## Return the metaObject
//   QObject.staticMetaObject

// proc SLOT*(signature: string): string =
//   ## Decorate the signature for being a slot
//   "1" & signature

// proc SIGNAL*(signature: string): string =
//   ## Decorates the signature for being a signal
//   "2" & signature

// proc connect*(typ: type QObject, sender: QObject, senderFunc: string, receiver: QObject, receiverFunc: string, connectionType: ConnectionType = ConnectionType.AutoConnection): QMetaObjectConnection =
//   ## String based QObject connections
//   let conn: DosQMetaObjectConnection = dos_qobject_connect_static(sender.vptr, senderFunc.cstring, receiver.vptr, receiverFunc.cstring, connectionType.cint)
//   result = QMetaObjectConnection.new(conn)

// proc connectRawLambda*[T](typ: type QObject, sender: QObject, senderFunc: string, context: QObject, p: T, connectionType: ConnectionType = ConnectionType.AutoConnection): QMetaObjectConnection =
//   let id = LambdaInvoker.instance.add(p)
//   let conn = dos_qobject_connect_lambda_with_context_static(sender.vptr, senderFunc.cstring, context.vptr, lambdaCallback, cast[pointer](id), connectionType.cint)
//   result = QMetaObjectConnection.new(conn)

// proc connectRawLambda*[T](typ: type QObject, sender: QObject, senderFunc: string, p: T, connectionType: ConnectionType = ConnectionType.AutoConnection): QMetaObjectConnection =
//   let id = LambdaInvoker.instance.add(p)
//   let conn = dos_qobject_connect_lambda_static(sender.vptr, senderFunc.cstring, lambdaCallback, cast[pointer](id), connectionType.cint)
//   result = QMetaObjectConnection.new(conn)

// macro connect*(typ: type QObject, sender: QObject, senderFuncUntyped: typed, receiverFunc: typed, connectionType: ConnectionType = ConnectionType.AutoConnection): untyped =
//   ## Connect a QObject signal to another QObject signal or slot
//   let senderFunc = findOverload(sender, senderFuncUntyped)
//   if senderFunc.kind == nnkSym and isSignal(getImpl(senderFunc)) and receiverFunc.kind == nnkLambda:
//     let senderFuncSignature = newLit(generateSignature(senderFunc))
//     result = quote do:
//       QObject.connectRawLambda(`sender`, `senderFuncSignature`, `receiverFunc`, `connectionType`)
//   elif senderFunc.kind == nnkSym and isSignal(getImpl(senderFunc)) and isProcOrMethod(getImpl(receiverFunc)):
//     let senderFuncSignature = newLit(generateSignature(senderFunc))
//     result = quote do:
//       QObject.connectRawLambda(`sender`, `senderFuncSignature`, `receiverFunc`, `connectionType`)
//   elif senderFunc.kind == nnkSym and isSignal(getImpl(senderFunc)) and isLambdaSymbol(receiverFunc):
//     let senderFuncSignature = newLit(generateSignature(senderFunc))
//     result = quote do:
//       QObject.connectRawLambda(`sender`, `senderFuncSignature`, `receiverFunc`, `connectionType`)

// macro connect*(typ: type QObject, sender: typed, senderFuncUntyped: typed, receiver: QObject, receiverFunc: typed, connectionType: ConnectionType = ConnectionType.AutoConnection): untyped =
//   ## Connect a QObject signal to another QObject signal or slot
//   let senderFunc = findOverload(sender, senderFuncUntyped)
//   if senderFunc.kind == nnkSym and isSignal(getImpl(senderFunc)) and receiverFunc.kind == nnkSym and isSignalOrSlot(getImpl(receiverFunc)):
//     let senderFuncSignature = newLit(generateSignature(senderFunc))
//     let receiverFuncSignature = newLit(generateSignature(receiverFunc))
//     result = quote do:
//       QObject.connect(`sender`, `senderFuncSignature`, `receiver`, `receiverFuncSignature`, `connectionType`)
//   elif senderFunc.kind == nnkSym and isSignal(getImpl(senderFunc)) and receiverFunc.kind == nnkLambda:
//     let senderFuncSignature = newLit(generateSignature(senderFunc))
//     result = quote do:
//       QObject.connectRawLambda(`sender`, `senderFuncSignature`, `receiver`, `receiverFunc`, `connectionType`)
//   elif senderFunc.kind == nnkSym and isSignal(getImpl(senderFunc)) and isProcOrMethod(getImpl(receiverFunc)):
//     let senderFuncSignature = newLit(generateSignature(senderFunc))
//     result = quote do:
//       QObject.connectRawLambda(`sender`, `senderFuncSignature`, `receiver`, `receiverFunc`, `connectionType`)
//   elif senderFunc.kind == nnkSym and isSignal(getImpl(senderFunc)) and isLambdaSymbol(receiverFunc):
//     let senderFuncSignature = newLit(generateSignature(senderFunc))
//     result = quote do:
//       QObject.connectRawLambda(`sender`, `senderFuncSignature`, `receiver`, `receiverFunc`, `connectionType`)
//   else:
//     assert(false)

// proc disconnect*(typ: type QObject, sender: QObject, senderFunc: string, receiver: QObject, receiverFunc: string) =
//   ## Disconnect a qobject signal/slot connection
//   dos_qobject_disconnect_static(sender.vptr, senderFunc.cstring, receiver.vptr, receiverFunc.cstring)

// proc disconnect*(self: QObject, senderFunc: string, receiver: QObject, receiverFunc: string) =
//   ## Disconnect a qobject signal/slot connection
//   QObject.disconnect(self, senderFunc, receiver, receiverFunc)

// proc disconnect*(typ: type QObject, connection: QMetaObjectConnection) =
//   ## Disconnect a qobject signal/slot connection
//   dos_qobject_disconnect_with_connection_static(connection.vptr)

// proc disconnect*(self: QObject, connection: QMetaObjectConnection) =
//   ## Disconnect a qobject signal/slot connection
//   QObject.disconnect(connection)

// proc emit*(qobject: QObject, signalName: string, arguments: openarray[QVariant] = []) =
//   ## Emit the signal with the given name and values
//   var dosArguments: seq[DosQVariant] = @[]
//   for argument in arguments:
//     dosArguments.add(argument.vptr)
//   let dosNumArguments = dosArguments.len.cint
//   let dosArgumentsPtr: ptr DosQVariant = if dosArguments.len > 0: dosArguments[0].unsafeAddr else: nil
//   dos_qobject_signal_emit(qobject.vptr, signalName.cstring, dosNumArguments, cast[ptr DosQVariantArray](dosArgumentsPtr))

// method onSlotCalled*(self: QObject, slotName: string, arguments: openarray[QVariant]) {.base.} =
//   ## Called from the dotherside library when a slot is called from Qml.
//   discard()

// proc qobjectCallback(qobjectPtr: pointer, slotNamePtr: DosQVariant, dosArgumentsLength: cint, dosArguments: ptr DosQVariantArray) {.cdecl, exportc.} =
//   ## Called from the dotherside library for invoking a slot
//   let qobject = cast[QObject](qobjectPtr)
//   GC_ref(qobject)
//   # Retrieve slot name
//   let slotName = newQVariant(slotNamePtr, Ownership.Clone) # Don't take ownership but clone
//   defer: slotName.delete
//   # Retrieve arguments
//   let arguments = toQVariantSequence(dosArguments, dosArgumentsLength, Ownership.Clone) # Don't take ownership but clone
//   defer: arguments.delete
//   # Forward to args to the slot
//   qobject.onSlotCalled(slotName.stringVal, arguments)
//   # Update the slot return value
//   dos_qvariant_assign(dosArguments[0], arguments[0].vptr)
//   GC_unref(qobject)

// proc deleteLater*(self: QObject) =
//   debugMsg("QObject", "deleteLater")
//   ## Delete a QObject
//   if not self.owner or self.vptr.isNil:
//     return
//   dos_qobject_deleteLater(self.vptr)
//   self.vptr.resetToNil

// proc objectNameChanged*(self: QObject, objectName: string) {.signal.} =
//   ## Emit the object name changed signal
//   self.emit("objectNameChanged", [newQVariant(objectName)])
