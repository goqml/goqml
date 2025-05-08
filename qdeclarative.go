package goqml

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var ctorTable cmap.ConcurrentMap[int, func() IQObject] = cmap.NewWithCustomShardingFunction[int, func() IQObject](func(key int) uint32 {
	return uint32(key)
})

var qObjCache cmap.ConcurrentMap[uintptr, IQObject] = cmap.NewWithCustomShardingFunction[uintptr, IQObject](func(key uintptr) uint32 {
	return uint32(key)
})

var creator = purego.NewCallback(func(_ purego.CDecl, id int, dosQObject DosQObject, qObjectStore QObjectStore, dosQObjectStore DosQObjectStore) uintptr {
	ctor, ok := ctorTable.Get(int(id))
	if !ok {
		fmt.Println("QmlRegisterType: unknown id")
		return 1
	}
	qObject := ctor()
	*qObjectStore = (uintptr)(ptrOfIQObjectReal(qObject))
	*dosQObjectStore = (uintptr)(qObject.getVPtr())
	qObject.setVPtr(dosQObject)
	qObject.setOwned(false)
	return 0
})

var deleter = purego.NewCallback(func(_ purego.CDecl, id int, qObject unsafe.Pointer) uintptr {
	qObjCache.Remove(uintptr(qObject))
	return 0
})

func QmlRegisterType[T any, PT interface {
	IQObject
	*T
}](uri string, major int, minor int, qmlName string, ctor func() PT) int {
	metaObject := (PT)(nil).StaticMetaObject()
	id := int(DosQDeclarativeQmlRegisterType(&DosQmlRegisterType{
		major:            int(major),
		minor:            int(minor),
		uri:              stringToCharPtr(nil, uri),
		qml:              stringToCharPtr(nil, qmlName),
		staticMetaObject: metaObject.vptr,
		createCallback:   creator,
		deleteCallback:   deleter,
	}))
	ctorTable.Set(id, func() IQObject {
		obj := ctor()
		qObjCache.Set(uintptr(obj.getVPtr()), obj)
		return obj
	})
	return id
}

func QmlRegisterSingletonType[T any, PT interface {
	IQObject
	*T
}](uri string, major int, minor int, qmlName string, ctor func() PT) int {
	metaObject := (PT)(nil).StaticMetaObject()
	id := int(DosQDeclarativeQmlRegisterSingletonType(&DosQmlRegisterType{
		major:            int(major),
		minor:            int(minor),
		uri:              stringToCharPtr(nil, uri),
		qml:              stringToCharPtr(nil, qmlName),
		staticMetaObject: metaObject.vptr,
		createCallback:   creator,
		deleteCallback:   deleter,
	}))
	ctorTable.Set(id, func() IQObject {
		obj := ctor()
		qObjCache.Set(uintptr(obj.getVPtr()), obj)
		return obj
	})
	return id
}
