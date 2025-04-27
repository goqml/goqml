package goqml

type QMetaObjectConnection struct {
	vptr DosQMetaObjectConnection
}

func NewQMetaObjectConnection(vptr DosQMetaObjectConnection) *QMetaObjectConnection {
	return &QMetaObjectConnection{
		vptr: vptr,
	}
}

func (obj *QMetaObjectConnection) Disconnect() {
	dos.QObjectDisconnectWithConnectionStatic(obj.vptr)
}

func (obj *QMetaObjectConnection) Delete() {
	obj.Disconnect()
	dos.QMetaObjectConnectionDelete(obj.vptr)
}
