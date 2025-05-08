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
	engine.vptr = DosQQmlApplicationEngineCreate()
}

func (engine *QQmlApplicationEngine) Load(filename string) {
	dirname, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	DosQQmlApplicationEngineLoad(engine.vptr, filepath.Join(dirname, filename))
}

func (engine *QQmlApplicationEngine) LoadUrl(url *QUrl) {
	DosQQmlApplicationEngineLoadUrl(engine.vptr, url.vptr)
}

func (engine *QQmlApplicationEngine) LoadData(data string) {
	DosQQmlApplicationEngineLoadData(engine.vptr, data)
}

func (engine *QQmlApplicationEngine) addImportPath(path string) {
	DosQQmlApplicationEngineAddImportPath(engine.vptr, path)
}

func (engine *QQmlApplicationEngine) SetRootContextProperty(name string, value *QVariant) {
	ctx := DosQQmlApplicationEngineContext(engine.vptr)
	DosQQmlContextSetContextProperty(ctx, name, value.vptr)
}

func (engine *QQmlApplicationEngine) Delete() {
	DosQQmlApplicationEngineDelete(engine.vptr)
}
