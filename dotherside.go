package goqml

/*
#cgo LDFLAGS: -ldotherside

#include "DOtherSide/DOtherSide.h"
#include "DOtherSide/DOtherSideTypes.h"

#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/shapled/goqml/util"
)

type (
	QObjectStore             *uintptr
	DosQMetaObject           unsafe.Pointer
	DosQObject               unsafe.Pointer
	DosQObjectStore          *uintptr
	DosQUrl                  unsafe.Pointer
	DosQHashIntByteArray     unsafe.Pointer
	DosQVariant              unsafe.Pointer
	DosQVariantArray         *unsafe.Pointer
	DosQMetaObjectConnection unsafe.Pointer
	DosQModelIndex           unsafe.Pointer
	DosQAbstractItemModel    unsafe.Pointer
	DosQAbstractTableModel   unsafe.Pointer
	DosQAbstractListModel    unsafe.Pointer

	DosQmlRegisterType struct {
		major            int
		minor            int
		uri              unsafe.Pointer
		qml              unsafe.Pointer
		staticMetaObject DosQMetaObject
		createCallback   uintptr
		deleteCallback   uintptr
	}

	DosQObjectCallBack uintptr // func(purego.CDecl, unsafe.Pointer, DosQVariant, int, DosQVariantArray) uintptr

	DosParameterDefinition struct {
		name     unsafe.Pointer
		metaType int
	}

	DosSignalDefinition struct {
		name            unsafe.Pointer
		parametersCount int
		parameters      unsafe.Pointer // []DosParameterDefinition
	}

	DosSlotDefinition struct {
		name            unsafe.Pointer
		returnMetaType  int
		parametersCount int
		parameters      unsafe.Pointer // []DosParameterDefinition
	}

	DosPropertyDefinition struct {
		name             unsafe.Pointer
		propertyMetaType int
		readSlot         unsafe.Pointer
		writeSlot        unsafe.Pointer
		notifySignal     unsafe.Pointer
	}

	DosSignalDefinitions struct {
		count       int
		definitions unsafe.Pointer
	}

	DosSlotDefinitions struct {
		count       int
		definitions unsafe.Pointer
	}

	DosPropertyDefinitions struct {
		count       int
		definitions unsafe.Pointer
	}

	DosRowCountCallback     uintptr // proc(nimmodel: NimQAbstractItemModel, rawIndex: DosQModelIndex, result: var cint) {.cdecl.}
	DosColumnCountCallback  uintptr // proc(nimmodel: NimQAbstractItemModel, rawIndex: DosQModelIndex, result: var cint) {.cdecl.}
	DosDataCallback         uintptr // proc(nimmodel: NimQAbstractItemModel, rawIndex: DosQModelIndex, role: cint, result: DosQVariant) {.cdecl.}
	DosSetDataCallback      uintptr // proc(nimmodel: NimQAbstractItemModel, rawIndex: DosQModelIndex, value: DosQVariant, role: cint, result: var bool) {.cdecl.}
	DosRoleNamesCallback    uintptr // proc(nimmodel: NimQAbstractItemModel, result: DosQHashIntByteArray) {.cdecl.}
	DosFlagsCallback        uintptr // proc(nimmodel: NimQAbstractItemModel, index: DosQModelIndex, result: var cint) {.cdecl.}
	DosHeaderDataCallback   uintptr // proc(nimmodel: NimQAbstractItemModel, section: cint, orientation: cint, role: cint, result: DosQVariant) {.cdecl.}
	DosIndexCallback        uintptr // proc(nimmodel: NimQAbstractItemModel, row: cint, column: cint, parent: DosQModelIndex, result: DosQModelIndex) {.cdecl.}
	DosParentCallback       uintptr // proc(nimmodel: NimQAbstractItemModel, child: DosQModelIndex, result: DosQModelIndex) {.cdecl.}
	DosHasChildrenCallback  uintptr // proc(nimmodel: NimQAbstractItemModel, parent: DosQModelIndex, result: var bool) {.cdecl.}
	DosCanFetchMoreCallback uintptr // proc(nimmodel: NimQAbstractItemModel, parent: DosQModelIndex, result: var bool) {.cdecl.}
	DosFetchMoreCallback    uintptr // proc(nimmodel: NimQAbstractItemModel, parent: DosQModelIndex) {.cdecl.}

	DosQAbstractItemModelCallbacks struct {
		RowCount     DosRowCountCallback
		ColumnCount  DosColumnCountCallback
		Data         DosDataCallback
		SetData      DosSetDataCallback
		RoleNames    DosRoleNamesCallback
		Flags        DosFlagsCallback
		HeaderData   DosHeaderDataCallback
		Index        DosIndexCallback
		Parent       DosParentCallback
		HasChildren  DosHasChildrenCallback
		CanFetchMore DosCanFetchMoreCallback
		FetchMore    DosFetchMoreCallback
	}

	DosQObjectConnectLambdaCallback uintptr // func(_ purego.CDecl, uintptr, int, DosQVariantArray) uintptr
)

// CharArray
func DosCharArrayDelete(ptr unsafe.Pointer) {
	C.dos_chararray_delete((*C.char)(ptr))
}

// QCoreApplication
func DosQCoreApplicationApplicationDirPath() unsafe.Pointer {
	ptr := C.dos_qcoreapplication_application_dir_path()
	return unsafe.Pointer(ptr)
}

// QApplication
func DosQApplicationCreate() {
	C.dos_qapplication_create()
}

func DosQApplicationExec() {
	C.dos_qapplication_exec()
}

func DosQApplicationQuit() {
	C.dos_qapplication_quit()
}

func DosQApplicationDelete() {
	C.dos_qapplication_delete()
}

// QGuiApplication
func DosQGuiApplicationCreate() {
	C.dos_qguiapplication_create()
}

func DosQGuiApplicationExec() {
	C.dos_qguiapplication_exec()
}

func DosQGuiApplicationQuit() {
	C.dos_qguiapplication_quit()
}

func DosQGuiApplicationDelete() {
	C.dos_qguiapplication_delete()
}

// QQmlContext
func DosQQmlContextSetContextProperty(ctx unsafe.Pointer, name string, value DosQVariant) {
	s1 := C.CString(name)
	defer C.free(unsafe.Pointer(s1))
	C.dos_qqmlcontext_setcontextproperty(ctx, s1, unsafe.Pointer(value))
}

// QQmlApplicationEngine
func DosQQmlApplicationEngineCreate() unsafe.Pointer {
	return C.dos_qqmlapplicationengine_create()
}

func DosQQmlApplicationEngineLoad(engine unsafe.Pointer, path string) {
	C.dos_qqmlapplicationengine_load(engine, C.CString(path))
}

func DosQQmlApplicationEngineLoadUrl(engine unsafe.Pointer, url DosQUrl) {
	C.dos_qqmlapplicationengine_load_url(engine, unsafe.Pointer(url))
}

func DosQQmlApplicationEngineLoadData(engine unsafe.Pointer, data string) {
	C.dos_qqmlapplicationengine_load_data(engine, C.CString(data))
}

func DosQQmlApplicationEngineAddImportPath(engine unsafe.Pointer, path string) {
	C.dos_qqmlapplicationengine_add_import_path(engine, C.CString(path))
}

func DosQQmlApplicationEngineContext(engine unsafe.Pointer) unsafe.Pointer {
	return C.dos_qqmlapplicationengine_context(engine)
}

func DosQQmlApplicationEngineDelete(engine unsafe.Pointer) {
	C.dos_qqmlapplicationengine_delete(engine)
}

// QVariant
func DosQVariantCreate() DosQVariant {
	return DosQVariant(C.dos_qvariant_create())
}

func DosQVariantCreateInt(value int) DosQVariant {
	return DosQVariant(C.dos_qvariant_create_int(C.int(value)))
}

func DosQVariantCreateBool(value bool) DosQVariant {
	return DosQVariant(C.dos_qvariant_create_bool(C.bool(value)))
}

func DosQVariantCreateString(value string) DosQVariant {

	return DosQVariant(C.dos_qvariant_create_string(C.CString(value)))
}

func DosQVariantCreateQObject(obj DosQObject) DosQVariant {
	return DosQVariant(C.dos_qvariant_create_qobject(unsafe.Pointer(obj)))
}

func DosQVariantCreateQVariant(variant DosQVariant) DosQVariant {
	return DosQVariant(C.dos_qvariant_create_qvariant(unsafe.Pointer(variant)))
}

func DosQVariantCreateFloat(value float32) DosQVariant {
	return DosQVariant(C.dos_qvariant_create_float(C.float(value)))
}

func DosQVariantCreateDouble(value float64) DosQVariant {
	return DosQVariant(C.dos_qvariant_create_double(C.double(value)))
}

func DosQVariantDelete(variant DosQVariant) {
	C.dos_qvariant_delete(unsafe.Pointer(variant))
}

func DosQVariantIsNull(variant DosQVariant) bool {
	return bool(C.dos_qvariant_isnull(unsafe.Pointer(variant)))
}

func DosQVariantToInt(variant DosQVariant) int {
	return int(C.dos_qvariant_toInt(unsafe.Pointer(variant)))
}

func DosQVariantToBool(variant DosQVariant) bool {
	return bool(C.dos_qvariant_toBool(unsafe.Pointer(variant)))
}

func DosQVariantToString(variant DosQVariant) unsafe.Pointer {
	return unsafe.Pointer(C.dos_qvariant_toString(unsafe.Pointer(variant)))
}

func DosQVariantToDouble(variant DosQVariant) float64 {
	return float64(C.dos_qvariant_toDouble(unsafe.Pointer(variant)))
}

func DosQVariantToFloat(variant DosQVariant) float32 {
	return float32(C.dos_qvariant_toFloat(unsafe.Pointer(variant)))
}

func DosQVariantSetInt(variant DosQVariant, value int) {
	C.dos_qvariant_setInt(unsafe.Pointer(variant), C.int(value))
}

func DosQVariantSetBool(variant DosQVariant, value bool) {
	C.dos_qvariant_setBool(unsafe.Pointer(variant), C.bool(value))
}

func DosQVariantSetString(variant DosQVariant, value string) {
	str := C.CString(value)
	defer C.free(unsafe.Pointer(str))
	C.dos_qvariant_setString(unsafe.Pointer(variant), str)
}

func DosQVariantAssign(variant DosQVariant, other DosQVariant) {
	C.dos_qvariant_assign(unsafe.Pointer(variant), unsafe.Pointer(other))
}

func DosQVariantSetFloat(variant DosQVariant, value float32) {
	C.dos_qvariant_setFloat(unsafe.Pointer(variant), C.float(value))
}

func DosQVariantSetDouble(variant DosQVariant, value float64) {
	C.dos_qvariant_setDouble(unsafe.Pointer(variant), C.double(value))
}

func DosQVariantSetQObject(variant DosQVariant, obj DosQObject) {
	C.dos_qvariant_setQObject(unsafe.Pointer(variant), unsafe.Pointer(obj))
}

// QObject
func DosQObjectQMetaObject() DosQMetaObject {
	return DosQMetaObject(C.dos_qobject_qmetaobject())
}

func DosQObjectCreate(inst unsafe.Pointer, vptr DosQMetaObject, callback DosQObjectCallBack) DosQObject {
	return DosQObject(C.dos_qobject_create(inst, unsafe.Pointer(vptr), (*[0]byte)(unsafe.Pointer(callback))))
}

func DosQObjectObjectName(obj DosQObject) string {
	str := C.dos_qobject_objectName(unsafe.Pointer(obj))
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

func DosQObjectSetObjectName(obj DosQObject, name string) {
	str := C.CString(name)
	defer C.free(unsafe.Pointer(str))
	C.dos_qobject_setObjectName(unsafe.Pointer(obj), str)
}

func DosQObjectSignalEmit(obj DosQObject, signal string, argc int, argv DosQVariantArray) {
	str := C.CString(signal)
	defer C.free(unsafe.Pointer(str))
	C.dos_qobject_signal_emit(unsafe.Pointer(obj), str, C.int(argc), (*unsafe.Pointer)(argv))
}

func DosQObjectConnectStatic(sender DosQObject, signal string, receiver DosQObject, method string, connectionType int) DosQMetaObjectConnection {
	str1 := C.CString(signal)
	defer C.free(unsafe.Pointer(str1))
	str2 := C.CString(method)
	defer C.free(unsafe.Pointer(str2))
	return DosQMetaObjectConnection(C.dos_qobject_connect_static(unsafe.Pointer(sender), str1, unsafe.Pointer(receiver), str2, C.DosQtConnectionType(connectionType)))
}

func DosQObjectConnectLambdaStatic(sender DosQObject, signal string, callback DosQObjectConnectLambdaCallback, context uintptr, connectionType int) DosQMetaObjectConnection {
	str := C.CString(signal)
	defer C.free(unsafe.Pointer(str))
	return DosQMetaObjectConnection(C.dos_qobject_connect_lambda_static(unsafe.Pointer(sender), str, (*[0]byte)(unsafe.Pointer(callback)), unsafe.Pointer(context), C.DosQtConnectionType(connectionType)))
}

func DosQObjectConnectLambdaWithContextStatic(sender DosQObject, signal string, receiver DosQObject, context unsafe.Pointer, data unsafe.Pointer, connectionType int) DosQMetaObjectConnection {
	str := C.CString(signal)
	defer C.free(unsafe.Pointer(str))
	return DosQMetaObjectConnection(C.dos_qobject_connect_lambda_with_context_static(unsafe.Pointer(sender), str, unsafe.Pointer(receiver), (*[0]byte)(context), data, C.DosQtConnectionType(connectionType)))
}

func DosQObjectDisconnectStatic(sender DosQObject, signal string, receiver DosQObject, method string) {
	str1 := C.CString(signal)
	defer C.free(unsafe.Pointer(str1))
	str2 := C.CString(method)
	defer C.free(unsafe.Pointer(str2))
	C.dos_qobject_disconnect_static(unsafe.Pointer(sender), str1, unsafe.Pointer(receiver), str2)
}

func DosQObjectDisconnectWithConnectionStatic(connection DosQMetaObjectConnection) {
	C.dos_qobject_disconnect_with_connection_static(unsafe.Pointer(connection))
}

func DosQObjectDelete(obj DosQObject) {
	C.dos_qobject_delete(unsafe.Pointer(obj))
}

func DosQObjectDeleteLater(obj DosQObject) {
	C.dos_qobject_deleteLater(unsafe.Pointer(obj))
}

func DosSignalMacro(signal string) string {
	str := C.CString(signal)
	defer C.free(unsafe.Pointer(str))
	return C.GoString(C.dos_signal_macro(str))
}

func DosSlotMacro(slot string) string {
	str := C.CString(slot)
	defer C.free(unsafe.Pointer(str))
	return C.GoString(C.dos_slot_macro(str))
}

// QMetaObject::Connection
func DosQMetaObjectConnectionDelete(connection DosQMetaObjectConnection) {
	C.dos_qmetaobject_connection_delete(unsafe.Pointer(connection))
}

// QAbstractItemModel
func DosQAbstractItemModelQMetaObject() DosQMetaObject {
	return DosQMetaObject(C.dos_qabstractitemmodel_qmetaobject())
}

// QMetaObject
func DosQMetaObjectCreate(parentMetaObject DosQMetaObject, className string, signals []*SignalDefinition, slots []*SlotDefinition, properties []*PropertyDefinition) DosQMetaObject {
	str := C.CString(className)
	defer C.free(unsafe.Pointer(str))
	return DosQMetaObject(C.dos_qmetaobject_create(unsafe.Pointer(parentMetaObject), str, signals, slots, properties))
}

func DosQMetaObjectDelete(metaObject DosQMetaObject) {
	C.dos_qmetaobject_delete(unsafe.Pointer(metaObject))
}

func DosQMetaObjectInvokeMethod(obj DosQObject, callback unsafe.Pointer, callbackData unsafe.Pointer, connectionType int) bool {
	return bool(C.dos_qmetaobject_invoke_method(unsafe.Pointer(obj), (*[0]byte)(callback), callbackData, C.DosQtConnectionType(connectionType)))
}

// QUrl
func DosQUrlCreate(url string, mode int) DosQUrl {
	str := C.CString(url)
	defer C.free(unsafe.Pointer(str))
	return DosQUrl(C.dos_qurl_create(str, C.int(mode)))
}

func DosQUrlDelete(url DosQUrl) {
	C.dos_qurl_delete(unsafe.Pointer(url))
}

func DosQUrlToString(url DosQUrl) unsafe.Pointer {
	return unsafe.Pointer(C.dos_qurl_to_string(unsafe.Pointer(url)))
}

// QQuickView
func DosQQuickViewCreate() unsafe.Pointer {
	return C.dos_qquickview_create()
}

func DosQQuickViewDelete(view unsafe.Pointer) {
	C.dos_qquickview_delete(view)
}

func DosQQuickViewShow(view unsafe.Pointer) {
	C.dos_qquickview_show(view)
}

func DosQQuickViewSource(view unsafe.Pointer) string {
	str := C.dos_qquickview_source(view)
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

func DosQQuickViewSetSource(view unsafe.Pointer, source string) {
	str := C.CString(source)
	defer C.free(unsafe.Pointer(str))
	C.dos_qquickview_set_source(view, str)
}

// QHash<int, QByteArray>
func DosQHashIntByteArrayCreate() DosQHashIntByteArray {
	return DosQHashIntByteArray(C.dos_qhash_int_qbytearray_create())
}

func DosQHashIntByteArrayDelete(hash DosQHashIntByteArray) {
	C.dos_qhash_int_qbytearray_delete(unsafe.Pointer(hash))
}

func DosQHashIntByteArrayInsert(hash DosQHashIntByteArray, key int, value string) {
	str := C.CString(value)
	defer C.free(unsafe.Pointer(str))
	C.dos_qhash_int_qbytearray_insert(unsafe.Pointer(hash), C.int(key), str)
}

func DosQHashIntByteArrayValue(hash DosQHashIntByteArray, key int) string {
	str := C.dos_qhash_int_qbytearray_value(unsafe.Pointer(hash), C.int(key))
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

// QModelIndex
func DosQModelIndexCreate() DosQModelIndex {
	return DosQModelIndex(C.dos_qmodelindex_create())
}

func DosQModelIndexCreateQModelIndex(index DosQModelIndex) DosQModelIndex {
	return DosQModelIndex(C.dos_qmodelindex_create_qmodelindex(unsafe.Pointer(index)))
}

func DosQModelIndexDelete(index DosQModelIndex) {
	C.dos_qmodelindex_delete(unsafe.Pointer(index))
}

func DosQModelIndexRow(index DosQModelIndex) int {
	return int(C.dos_qmodelindex_row(unsafe.Pointer(index)))
}

func DosQModelIndexColumn(index DosQModelIndex) int {
	return int(C.dos_qmodelindex_column(unsafe.Pointer(index)))
}

func DosQModelIndexIsValid(index DosQModelIndex) bool {
	return bool(C.dos_qmodelindex_isValid(unsafe.Pointer(index)))
}

func DosQModelIndexData(index DosQModelIndex, role int) DosQVariant {
	return DosQVariant(C.dos_qmodelindex_data(unsafe.Pointer(index), C.int(role)))
}

func DosQModelIndexParent(index DosQModelIndex) DosQModelIndex {
	return DosQModelIndex(C.dos_qmodelindex_parent(unsafe.Pointer(index)))
}

func DosQModelIndexChild(index DosQModelIndex, row int, column int) DosQModelIndex {
	return DosQModelIndex(C.dos_qmodelindex_child(unsafe.Pointer(index), C.int(row), C.int(column)))
}

func DosQModelIndexSibling(index DosQModelIndex, row int, column int) DosQModelIndex {
	return DosQModelIndex(C.dos_qmodelindex_sibling(unsafe.Pointer(index), C.int(row), C.int(column)))
}

func DosQModelIndexAssign(index DosQModelIndex, other DosQModelIndex) {
	C.dos_qmodelindex_assign(unsafe.Pointer(index), unsafe.Pointer(other))
}

func DosQModelIndexInternalPointer(index DosQModelIndex) unsafe.Pointer {
	return C.dos_qmodelindex_internalPointer(unsafe.Pointer(index))
}

// QAbstractItemModel
func DosQAbstractItemModelCreate(
	userData unsafe.Pointer,
	metaObject DosQMetaObject,
	callbacks uintptr,
	modelCallbacks DosQAbstractItemModelCallbacks,
) DosQAbstractItemModel {
	return DosQAbstractItemModel(C.dos_qabstractitemmodel_create(userData, unsafe.Pointer(metaObject), (*[0]byte)(unsafe.Pointer(callbacks)), modelCallbacks))
}

func DosQAbstractItemModelBeginInsertRows(
	model DosQAbstractItemModel,
	parent DosQModelIndex,
	first int,
	last int,
) {
	C.dos_qabstractitemmodel_beginInsertRows(model, parent, C.int(first), C.int(last))
}

func DosQAbstractItemModelEndInsertRows(model DosQAbstractItemModel) {
	C.dos_qabstractitemmodel_endInsertRows(model)
}

func DosQAbstractItemModelBeginRemoveRows(
	model DosQAbstractItemModel,
	parent DosQModelIndex,
	first int,
	last int,
) {
	C.dos_qabstractitemmodel_beginRemoveRows(model, parent, C.int(first), C.int(last))
}

func DosQAbstractItemModelEndRemoveRows(model DosQAbstractItemModel) {
	C.dos_qabstractitemmodel_endRemoveRows(model)
}

func DosQAbstractItemModelBeginInsertColumns(
	model DosQAbstractItemModel,
	parent DosQModelIndex,
	first int,
	last int,
) {
	C.dos_qabstractitemmodel_beginInsertColumns(model, parent, C.int(first), C.int(last))
}

func DosQAbstractItemModelEndInsertColumns(model DosQAbstractItemModel) {
	C.dos_qabstractitemmodel_endInsertColumns(model)
}

func DosQAbstractItemModelBeginRemoveColumns(
	model DosQAbstractItemModel,
	parent DosQModelIndex,
	first int,
	last int,
) {
	C.dos_qabstractitemmodel_beginRemoveColumns(model, parent, C.int(first), C.int(last))
}

func DosQAbstractItemModelEndRemoveColumns(model DosQAbstractItemModel) {
	C.dos_qabstractitemmodel_endRemoveColumns(model)
}

func DosQAbstractItemModelBeginResetModel(model DosQAbstractItemModel) {
	C.dos_qabstractitemmodel_beginResetModel(model)
}

func DosQAbstractItemModelEndResetModel(model DosQAbstractItemModel) {
	C.dos_qabstractitemmodel_endResetModel(model)
}

func DosQAbstractItemModelDataChanged(
	model DosQAbstractItemModel,
	topLeft DosQModelIndex,
	bottomRight DosQModelIndex,
	roles unsafe.Pointer,
	roleCount int,
) {
	C.dos_qabstractitemmodel_dataChanged(model, topLeft, bottomRight, roles, C.int(roleCount))
}

func DosQAbstractItemModelCreateIndex(
	model DosQAbstractItemModel,
	row int,
	column int,
	internalPointer unsafe.Pointer,
) DosQModelIndex {
	return C.dos_qabstractitemmodel_createIndex(model, C.int(row), C.int(column), internalPointer)
}

func DosQAbstractItemModelHasChildren(model DosQAbstractItemModel, parent DosQModelIndex) bool {
	return bool(C.dos_qabstractitemmodel_hasChildren(model, parent))
}

func DosQAbstractItemModelHasIndex(
	model DosQAbstractItemModel,
	row int,
	column int,
	parent DosQModelIndex,
) bool {
	return bool(C.dos_qabstractitemmodel_hasIndex(model, C.int(row), C.int(column), parent))
}

func DosQAbstractItemModelCanFetchMore(model DosQAbstractItemModel, parent DosQModelIndex) bool {
	return bool(C.dos_qabstractitemmodel_canFetchMore(model, parent))
}

func DosQAbstractItemModelFetchMore(model DosQAbstractItemModel, parent DosQModelIndex) {
	C.dos_qabstractitemmodel_fetchMore(model, parent)
}

// QResource 相关字段转换为独立函数
func DosQResourceRegister(resourcePath string) {
	C.dos_qresource_register(resourcePath)
}

// QDeclarative 相关字段转换为独立函数
func DosQDeclarativeQmlRegisterType(registerType *DosQmlRegisterType) int {
	return C.dos_qdeclarative_qmlregistertype(registerType)
}

func DosQDeclarativeQmlRegisterSingletonType(registerType *DosQmlRegisterType) int {
	return C.dos_qdeclarative_qmlregistersingletontype(registerType)
}

// QAbstractListModel 相关字段转换为独立函数
func DosQAbstractListModelQMetaObject() DosQMetaObject {
	return C.dos_qabstractlistmodel_qmetaobject()
}

func DosQAbstractListModelCreate(
	userData unsafe.Pointer,
	metaObject DosQMetaObject,
	callback uintptr,
	modelCallbacks DosQAbstractItemModelCallbacks,
) DosQAbstractListModel {
	return C.dos_qabstractlistmodel_create(userData, metaObject, C.uintptr_t(callback), modelCallbacks)
}

func DosQAbstractListModelColumnCount(model DosQAbstractListModel, index DosQModelIndex) int {
	return int(C.dos_qabstractlistmodel_columnCount(model, index))
}

func DosQAbstractListModelParent(model DosQAbstractListModel, index DosQModelIndex) DosQModelIndex {
	return C.dos_qabstractlistmodel_parent(model, index)
}

func DosQAbstractListModelIndex(
	model DosQAbstractListModel,
	row int,
	column int,
	parentIndex DosQModelIndex,
) DosQModelIndex {
	return C.dos_qabstractlistmodel_index(model, C.int(row), C.int(column), parentIndex)
}

// QAbstractTableModel 相关字段转换为独立函数
func DosQAbstractTableModelQMetaObject() DosQMetaObject {
	return DosQMetaObject(C.dos_qabstracttablemodel_qmetaobject())
}

func DosQAbstractTableModelCreate(
	userData unsafe.Pointer,
	metaObject DosQMetaObject,
	callback uintptr,
	modelCallbacks DosQAbstractItemModelCallbacks,
) DosQAbstractTableModel {
	return C.dos_qabstracttablemodel_create(userData, unsafe.Pointer(metaObject), C.uintptr_t(callback), modelCallbacks)
}

func DosQAbstractTableModelParent(model DosQAbstractTableModel, index DosQModelIndex) DosQModelIndex {
	return C.dos_qabstracttablemodel_parent(model, index)
}

func DosQAbstractTableModelIndex(
	model DosQAbstractTableModel,
	row int,
	column int,
	parentIndex DosQModelIndex,
) DosQModelIndex {
	return C.dos_qabstracttablemodel_index(model, C.int(row), C.int(column), parentIndex)
}

func charPtrToString(ptr unsafe.Pointer) string {
	if ptr == nil {
		return ""
	}

	data := uintptr(ptr)
	len := 0
	for {
		b := *(*byte)(unsafe.Pointer(data + uintptr(len)))
		if b == 0 {
			break
		}
		len++
	}

	bs := unsafe.Slice((*byte)(ptr), len)
	return string(bs)
}

func stringToCharPtr(pg *util.PinGroup, s string) unsafe.Pointer {
	bs := []byte(s + "\x00")
	pg.Pin(bs)
	return unsafe.Pointer(&bs[0])
}

func sliceToPtr[T any](pg *util.PinGroup, arr []T) unsafe.Pointer {
	if len(arr) == 0 {
		return nil
	}
	if pg != nil {
		pg.Pin(arr)
	}
	return unsafe.Pointer(&arr[0])
}

func ptrArrayIndex(array unsafe.Pointer, index int) unsafe.Pointer {
	elemSize := unsafe.Sizeof(uintptr(0))
	elemPtr := unsafe.Pointer(uintptr(array) + uintptr(index)*elemSize)
	return unsafe.Pointer(*(**int)(elemPtr))
}
