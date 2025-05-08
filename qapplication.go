package goqml

import "runtime"

type QApplication struct {
	deleted bool
}

func NewQApplication() *QApplication {
	runtime.LockOSThread()

	app := &QApplication{}
	app.Setup()
	return app
}

func (app *QApplication) Setup() {
	DosQApplicationCreate()
	app.deleted = false
}

func (app *QApplication) ApplicationDirPath() string {
	ptr := DosQCoreApplicationApplicationDirPath()
	defer DosCharArrayDelete(ptr)
	return charPtrToString(ptr)
}

func (app *QApplication) Exec() {
	DosQApplicationExec()
}

func (app *QApplication) Quit() {
	DosQApplicationQuit()
}

func (app *QApplication) Delete() {
	if app.deleted {
		return
	}
	DosQApplicationDelete()
	app.deleted = true

	runtime.UnlockOSThread()
}
