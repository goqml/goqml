package goqml

type QApplication struct {
	deleted bool
}

func NewQApplication() *QApplication {
	app := &QApplication{}
	app.Setup()
	return app
}

func (app *QApplication) Setup() {
	dos.QApplicationCreate()
	app.deleted = false
}

func (app *QApplication) ApplicationDirPath() string {
	ptr := dos.QCoreApplicationApplicationDirPath()
	defer dos.CharArrayDelete(ptr)
	return charPtrToString(ptr)
}

func (app *QApplication) Exec() {
	dos.QApplicationExec()
}

func (app *QApplication) Quit() {
	dos.QApplicationQuit()
}

func (app *QApplication) Delete() {
	if app.deleted {
		return
	}
	dos.QApplicationDelete()
	app.deleted = true
}
