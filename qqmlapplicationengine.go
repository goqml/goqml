package goqml

import (
	"fmt"
	"os"
	"path/filepath"
	"unsafe"
)

type QQmlApplicationEngine struct {
	vptr unsafe.Pointer
}

func NewQQmlApplicationEngine() *QQmlApplicationEngine {
	var engine QQmlApplicationEngine
	engine.Setup()
	return &engine
}

func (engine *QQmlApplicationEngine) Setup() {
	engine.vptr = dos.QQmlApplicationEngineCreate()
}

func (engine *QQmlApplicationEngine) Load(filename string) {
	dirname, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	dos.QQmlApplicationEngineLoad(engine.vptr, filepath.Join(dirname, filename))
}

func (engine *QQmlApplicationEngine) LoadUrl(url *QUrl) {
	dos.QQmlApplicationEngineLoadUrl(engine.vptr, url.vptr)
}

func (engine *QQmlApplicationEngine) LoadData(data string) {
	dos.QQmlApplicationEngineLoadData(engine.vptr, data)
}

func (engine *QQmlApplicationEngine) addImportPath(path string) {
	dos.QQmlApplicationEngineAddImportPath(engine.vptr, path)
}

func (engine *QQmlApplicationEngine) SetRootContextProperty(name string, value *QVariant) {
	ctx := dos.QQmlApplicationEngineContext(engine.vptr)
	dos.QQmlContextSetContextProperty(ctx, name, value.vptr)
}

func (engine *QQmlApplicationEngine) Delete() {
	dos.QQmlApplicationEngineDelete(engine.vptr)
}
