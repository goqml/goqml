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
	DosQObjectDisconnectWithConnectionStatic(obj.vptr)
}

func (obj *QMetaObjectConnection) Delete() {
	obj.Disconnect()
	DosQMetaObjectConnectionDelete(obj.vptr)
}
