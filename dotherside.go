package goqml

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/shapled/puregostruct"
)

type (
	DosQMetaObject   unsafe.Pointer
	DosQObject       unsafe.Pointer
	DosQUrl          unsafe.Pointer
	DosQVariant      unsafe.Pointer
	DosQVariantArray unsafe.Pointer // []DosQVariant

	DosQObjectCallBack func(purego.CDecl, unsafe.Pointer, DosQVariant, int, DosQVariantArray) uintptr

	DosParameterDefinition struct {
		name     unsafe.Pointer
		metaType int32
	}

	DosSignalDefinition struct {
		name            unsafe.Pointer
		parametersCount int32
		parameters      unsafe.Pointer // []DosParameterDefinition
	}

	DosSlotDefinition struct {
		name            unsafe.Pointer
		returnMetaType  int32
		parametersCount int32
		parameters      unsafe.Pointer // []DosParameterDefinition
	}

	DosPropertyDefinition struct {
		name             unsafe.Pointer
		propertyMetaType int32
		readSlot         unsafe.Pointer
		writeSlot        unsafe.Pointer
		notifySignal     unsafe.Pointer
	}

	DosSignalDefinitions struct {
		count       int32
		definitions unsafe.Pointer
	}

	DosSlotDefinitions struct {
		count       int32
		definitions unsafe.Pointer
	}

	DosPropertyDefinitions struct {
		count       int32
		definitions unsafe.Pointer
	}
)

type Dos struct {
	// CharArray
	CharArrayDelete func(unsafe.Pointer) `purego:"dos_chararray_delete"`

	// QCoreApplication
	QCoreApplicationApplicationDirPath func() unsafe.Pointer `purego:"dos_qcoreapplication_application_dir_path"`

	// QApplication
	QApplicationCreate func() `purego:"dos_qapplication_create"`
	QApplicationExec   func() `purego:"dos_qapplication_exec"`
	QApplicationQuit   func() `purego:"dos_qapplication_quit"`
	QApplicationDelete func() `purego:"dos_qapplication_delete"`

	// QGuiApplication
	QGuiApplicationCreate func() `purego:"dos_qguiapplication_create"`
	QGuiApplicationExec   func() `purego:"dos_qguiapplication_exec"`
	QGuiApplicationQuit   func() `purego:"dos_qguiapplication_quit"`
	QGuiApplicationDelete func() `purego:"dos_qguiapplication_delete"`

	// QQmlContext
	DosQQmlContextSetContextProperty func(unsafe.Pointer, string, DosQVariant) `purego:"dos_qqmlcontext_setcontextproperty"`

	// QQmlApplicationEngine
	QQmlApplicationEngineCreate        func() unsafe.Pointer               `purego:"dos_qqmlapplicationengine_create"`
	QQmlApplicationEngineLoad          func(unsafe.Pointer, string)        `purego:"dos_qqmlapplicationengine_load"`
	QQmlApplicationEngineLoadUrl       func(unsafe.Pointer, DosQUrl)       `purego:"dos_qqmlapplicationengine_load_url"`
	QQmlApplicationEngineLoadData      func(unsafe.Pointer, string)        `purego:"dos_qqmlapplicationengine_load_data"`
	QQmlApplicationEngineAddImportPath func(unsafe.Pointer, string)        `purego:"dos_qqmlapplicationengine_add_import_path"`
	QQmlApplicationEngineContext       func(unsafe.Pointer) unsafe.Pointer `purego:"dos_qqmlapplicationengine_context"`
	QQmlApplicationEngineDelete        func(unsafe.Pointer)                `purego:"dos_qqmlapplicationengine_delete"`

	// QVariant
	QVariantCreate         func() DosQVariant               `purego:"dos_qvariant_create"`
	QVariantCreateInt      func(int) DosQVariant            `purego:"dos_qvariant_create_int"`
	QVariantCreateBool     func(bool) DosQVariant           `purego:"dos_qvariant_create_bool"`
	QVariantCreateString   func(string) DosQVariant         `purego:"dos_qvariant_create_string"`
	QVariantCreateQObject  func(DosQObject) DosQVariant     `purego:"dos_qvariant_create_qobject"`
	QVariantCreateQVariant func(DosQVariant) DosQVariant    `purego:"dos_qvariant_create_qvariant"`
	QVariantCreateFloat    func(float32) DosQVariant        `purego:"dos_qvariant_create_float"`
	QVariantCreateDouble   func(float64) DosQVariant        `purego:"dos_qvariant_create_double"`
	QVariantDelete         func(DosQVariant)                `purego:"dos_qvariant_delete"`
	QVariantIsNull         func(DosQVariant) bool           `purego:"dos_qvariant_isnull"`
	QVariantToInt          func(DosQVariant) int            `purego:"dos_qvariant_toInt"`
	QVariantToBool         func(DosQVariant) bool           `purego:"dos_qvariant_toBool"`
	QVariantToString       func(DosQVariant) unsafe.Pointer `purego:"dos_qvariant_toString"`
	QVariantToDouble       func(DosQVariant) float64        `purego:"dos_qvariant_toDouble"`
	QVariantToFloat        func(DosQVariant) float32        `purego:"dos_qvariant_toFloat"`
	QVariantSetInt         func(DosQVariant, int)           `purego:"dos_qvariant_setInt"`
	QVariantSetBool        func(DosQVariant, bool)          `purego:"dos_qvariant_setBool"`
	QVariantSetString      func(DosQVariant, string)        `purego:"dos_qvariant_setString"`
	QVariantAssign         func(DosQVariant, DosQVariant)   `purego:"dos_qvariant_assign"`
	QVariantSetFloat       func(DosQVariant, float32)       `purego:"dos_qvariant_setFloat"`
	QVariantSetDouble      func(DosQVariant, float64)       `purego:"dos_qvariant_setDouble"`
	QVariantSetQObject     func(DosQVariant, DosQObject)    `purego:"dos_qvariant_setQObject"`

	// QObject
	QObjectQMetaObject                    func() DosQMetaObject                                                                      `purego:"dos_qobject_qmetaobject"`
	QObjectCreate                         func(unsafe.Pointer, DosQMetaObject, DosQObjectCallBack) DosQObject                        `purego:"dos_qobject_create"`
	QObjectObjectName                     func(DosQObject) string                                                                    `purego:"dos_qobject_objectName"`
	QObjectSetObjectName                  func(DosQObject, string)                                                                   `purego:"dos_qobject_setObjectName"`
	QObjectSignalEmit                     func(DosQObject, string, int, DosQVariantArray)                                            `purego:"dos_qobject_signal_emit"`
	QObjectConnectStatic                  func(DosQObject, string, DosQObject, string, int32) unsafe.Pointer                         `purego:"dos_qobject_connect_static"`
	QObjectConnectLambdaStatic            func(DosQObject, string, unsafe.Pointer, unsafe.Pointer, int32) unsafe.Pointer             `purego:"dos_qobject_connect_lambda_static"`
	QObjectConnectLambdaWithContextStatic func(DosQObject, string, DosQObject, unsafe.Pointer, unsafe.Pointer, int32) unsafe.Pointer `purego:"dos_qobject_connect_lambda_with_context_static"`
	QObjectDisconnectStatic               func(DosQObject, string, unsafe.Pointer, string)                                           `purego:"dos_qobject_disconnect_static"`
	QObjectDisconnectWithConnectionStatic func(unsafe.Pointer)                                                                       `purego:"dos_qobject_disconnect_with_connection_static"`
	QObjectDelete                         func(DosQObject)                                                                           `purego:"dos_qobject_delete"`
	QObjectDeleteLater                    func(DosQObject)                                                                           `purego:"dos_qobject_deleteLater"`
	SignalMacro                           func(string) string                                                                        `purego:"dos_signal_macro"`
	SlotMacro                             func(string) string                                                                        `purego:"dos_slot_macro"`

	// QMetaObject::Connection
	QMetaObjectConnectionDelete func(unsafe.Pointer) `purego:"dos_qmetaobject_connection_delete"`

	// QAbstractItemModel
	QAbstractItemModelQMetaObject func() DosQMetaObject `purego:"dos_qabstractitemmodel_qmetaobject"`

	// QMetaObject
	QMetaObjectCreate       func(DosQMetaObject, string, *DosSignalDefinitions, *DosSlotDefinitions, *DosPropertyDefinitions) DosQMetaObject `purego:"dos_qmetaobject_create"`
	QMetaObjectDelete       func(DosQMetaObject)                                                                                             `purego:"dos_qmetaobject_delete"`
	QMetaObjectInvokeMethod func(DosQObject, unsafe.Pointer, unsafe.Pointer, int) bool                                                       `purego:"dos_qmetaobject_invoke_method"`

	// QUrl
	QUrlCreate   func(string, int) DosQUrl    `purego:"dos_qurl_create"`
	QUrlDelete   func(DosQUrl)                `purego:"dos_qurl_delete"`
	QUrlToString func(DosQUrl) unsafe.Pointer `purego:"dos_qurl_to_string"`

	// QQuickView
	QQuickViewCreate    func() unsafe.Pointer        `purego:"dos_qquickview_create"`
	QQuickViewDelete    func(unsafe.Pointer)         `purego:"dos_qquickview_delete"`
	QQuickViewShow      func(unsafe.Pointer)         `purego:"dos_qquickview_show"`
	QQuickViewSource    func(unsafe.Pointer) string  `purego:"dos_qquickview_source"`
	QQuickViewSetSource func(unsafe.Pointer, string) `purego:"dos_qquickview_set_source"`

	// QHash<int, QByteArra>
	QHashIntByteArrayCreate func() unsafe.Pointer             `purego:"dos_qhash_int_qbytearray_create"`
	QHashIntByteArrayDelete func(unsafe.Pointer)              `purego:"dos_qhash_int_qbytearray_delete"`
	QHashIntByteArrayInsert func(unsafe.Pointer, int, string) `purego:"dos_qhash_int_qbytearray_insert"`
	QHashIntByteArrayValue  func(unsafe.Pointer, int) string  `purego:"dos_qhash_int_qbytearray_value"`

	// QModelIndex
	QModelIndexCreate            func() unsafe.Pointer                             `purego:"dos_qmodelindex_create"`
	QModelIndexCreateQModelIndex func(unsafe.Pointer) unsafe.Pointer               `purego:"dos_qmodelindex_create_qmodelindex"`
	QModelIndexDelete            func(unsafe.Pointer)                              `purego:"dos_qmodelindex_delete"`
	QModelIndexRow               func(unsafe.Pointer) int32                        `purego:"dos_qmodelindex_row"`
	QModelIndexColumn            func(unsafe.Pointer) int32                        `purego:"dos_qmodelindex_column"`
	QModelIndexIsValid           func(unsafe.Pointer) bool                         `purego:"dos_qmodelindex_isValid"`
	QModelIndexData              func(unsafe.Pointer, int32) unsafe.Pointer        `purego:"dos_qmodelindex_data"`
	QModelIndexParent            func(unsafe.Pointer) unsafe.Pointer               `purego:"dos_qmodelindex_parent"`
	QModelIndexChild             func(unsafe.Pointer, int32, int32) unsafe.Pointer `purego:"dos_qmodelindex_child"`
	QModelIndexSibling           func(unsafe.Pointer, int32, int32) unsafe.Pointer `purego:"dos_qmodelindex_sibling"`
	QModelIndexAssign            func(unsafe.Pointer, unsafe.Pointer)              `purego:"dos_qmodelindex_assign"`
	QModelIndexInternalPointer   func(unsafe.Pointer) unsafe.Pointer               `purego:"dos_qmodelindex_internalPointer"`

	// QAbstractItemModel
	QAbstractItemModelCreate             func(unsafe.Pointer, DosQMetaObject, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer `purego:"dos_qabstractitemmodel_create"`
	QAbstractItemModelBeginInsertRows    func(unsafe.Pointer, unsafe.Pointer, int32, int32)                                  `purego:"dos_qabstractitemmodel_beginInsertRows"`
	QAbstractItemModelEndInsertRows      func(unsafe.Pointer)                                                                `purego:"dos_qabstractitemmodel_endInsertRows"`
	QAbstractItemModelBeginRemoveRows    func(unsafe.Pointer, unsafe.Pointer, int32, int32)                                  `purego:"dos_qabstractitemmodel_beginRemoveRows"`
	QAbstractItemModelEndRemoveRows      func(unsafe.Pointer)                                                                `purego:"dos_qabstractitemmodel_endRemoveRows"`
	QAbstractItemModelBeginInsertColumns func(unsafe.Pointer, unsafe.Pointer, int32, int32)                                  `purego:"dos_qabstractitemmodel_beginInsertColumns"`
	QAbstractItemModelEndInsertColumns   func(unsafe.Pointer)                                                                `purego:"dos_qabstractitemmodel_endInsertColumns"`
	QAbstractItemModelBeginRemoveColumns func(unsafe.Pointer, unsafe.Pointer, int32, int32)                                  `purego:"dos_qabstractitemmodel_beginRemoveColumns"`
	QAbstractItemModelEndRemoveColumns   func(unsafe.Pointer)                                                                `purego:"dos_qabstractitemmodel_endRemoveColumns"`
	QAbstractItemModelBeginResetModel    func(unsafe.Pointer)                                                                `purego:"dos_qabstractitemmodel_beginResetModel"`
	QAbstractItemModelEndResetModel      func(unsafe.Pointer)                                                                `purego:"dos_qabstractitemmodel_endResetModel"`
	QAbstractItemModelDataChanged        func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int32)         `purego:"dos_qabstractitemmodel_dataChanged"`
	QAbstractItemModelCreateIndex        func(unsafe.Pointer, int32, int32, unsafe.Pointer) unsafe.Pointer                   `purego:"dos_qabstractitemmodel_createIndex"`
	QAbstractItemModelHasChildren        func(unsafe.Pointer, unsafe.Pointer) bool                                           `purego:"dos_qabstractitemmodel_hasChildren"`
	QAbstractItemModelHasIndex           func(unsafe.Pointer, int, int, unsafe.Pointer) bool                                 `purego:"dos_qabstractitemmodel_hasIndex"`
	QAbstractItemModelCanFetchMore       func(unsafe.Pointer, unsafe.Pointer) bool                                           `purego:"dos_qabstractitemmodel_canFetchMore"`
	QAbstractItemModelFetchMore          func(unsafe.Pointer, unsafe.Pointer)                                                `purego:"dos_qabstractitemmodel_fetchMore"`

	// QResource
	QResourceRegister func(string) `purego:"dos_qresource_register"`

	// QDeclarative
	QDeclarativeQmlRegisterType          func(unsafe.Pointer) int32 `purego:"dos_qdeclarative_qmlregistertype"`
	QDeclarativeQmlRegisterSingletonType func(unsafe.Pointer) int32 `purego:"dos_qdeclarative_qmlregistersingletontype"`

	// QAbstractListModel
	QAbstractListModelQMetaObject func() DosQMetaObject                                                               `purego:"dos_qabstractlistmodel_qmetaobject"`
	QAbstractListModelCreate      func(unsafe.Pointer, DosQMetaObject, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer `purego:"dos_qabstractlistmodel_create"`
	QAbstractListModelColumnCount func(unsafe.Pointer, unsafe.Pointer) int32                                          `purego:"dos_qabstractlistmodel_columnCount"`
	QAbstractListModelParent      func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer                                 `purego:"dos_qabstractlistmodel_parent"`
	QAbstractListModelIndex       func(unsafe.Pointer, int32, int32, unsafe.Pointer) unsafe.Pointer                   `purego:"dos_qabstractlistmodel_index"`

	// QAbstractTableModel
	QAbstractTableModelQMetaObject func() DosQMetaObject                                                               `purego:"dos_qabstracttablemodel_qmetaobject"`
	QAbstractTableModelCreate      func(unsafe.Pointer, DosQMetaObject, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer `purego:"dos_qabstracttablemodel_create"`
	QAbstractTableModelParent      func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer                                 `purego:"dos_qabstracttablemodel_parent"`
	QAbstractTableModelIndex       func(unsafe.Pointer, int32, int32, unsafe.Pointer) unsafe.Pointer                   `purego:"dos_qabstracttablemodel_index"`
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

var cache cmap.ConcurrentMap[uintptr, []byte] = cmap.NewWithCustomShardingFunction[uintptr, []byte](func(key uintptr) uint32 {
	return uint32(key % 8)
})

func stringToCharPtr(s string) unsafe.Pointer {
	bs := []byte(s + "\x00")
	ptr := unsafe.Pointer(&bs[0])
	cache.Set(uintptr(ptr), bs)
	return unsafe.Pointer(&bs[0])
}

func releaseBytes() {
	cache.Clear()
}

func sliceToPtr[T any](arr []T) unsafe.Pointer {
	if len(arr) == 0 {
		return nil
	}
	return unsafe.Pointer(&arr[0])
}

func ptrArrayIndex(array unsafe.Pointer, index int) unsafe.Pointer {
	elemSize := unsafe.Sizeof(uintptr(0))
	elemPtr := unsafe.Pointer(uintptr(array) + uintptr(index)*elemSize)
	return unsafe.Pointer(*(**int)(elemPtr))
}

func getSystemLibrary() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{"libDOtherSide.dll", "DOtherSide.dll"}
	case "linux":
		return []string{"libDOtherSide.so"}
	case "darwin":
		return []string{"libDOtherSide.dylib"}
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

var dos *Dos = func() *Dos {
	var dos Dos
	puregostruct.LoadLibrary(&dos, getSystemLibrary()...)
	return &dos
}()
