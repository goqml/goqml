package goqml

import "github.com/goqml/goqml/util"

type QMetaObject struct {
	vptr       DosQMetaObject
	signals    []*SignalDefinition
	slots      []*SlotDefinition
	properties []*PropertyDefinition
}

func NewQObjectMetaObject() *QMetaObject {
	return &QMetaObject{vptr: dos.QObjectQMetaObject()}
}

func NewQAbstractItemModelMetaObject() *QMetaObject {
	return &QMetaObject{vptr: dos.QAbstractItemModelQMetaObject()}
}

func NewQAbstractListModelMetaObject() *QMetaObject {
	return &QMetaObject{vptr: dos.QAbstractListModelQMetaObject()}
}

func NewQAbstractTableModelMetaObject() *QMetaObject {
	return &QMetaObject{vptr: dos.QAbstractTableModelQMetaObject()}
}

func NewQMetaObject(
	super *QMetaObject,
	className string,
	signals []*SignalDefinition,
	slots []*SlotDefinition,
	properties []*PropertyDefinition,
) *QMetaObject {
	for _, property := range properties {
		if property.Getter != "" {
			slots = append(slots, &SlotDefinition{
				Name:        property.Getter,
				RetMetaType: property.MetaType,
				Params:      []*ParameterDefinition{},
			})
		}
		if property.Setter != "" {
			slots = append(slots, &SlotDefinition{
				Name:        property.Setter,
				RetMetaType: QMetaTypeVoid,
				Params:      []*ParameterDefinition{{MetaType: property.MetaType, Name: "value"}},
			})
		}
		if property.Emitter != "" {
			signals = append(signals, &SignalDefinition{
				Name:   property.Emitter,
				Params: []*ParameterDefinition{{MetaType: property.MetaType, Name: "value"}},
			})
		}
	}

	meta := &QMetaObject{signals: signals, slots: slots, properties: properties}
	meta.Setup(super, className, signals, slots, properties)
	return meta
}

func (meta *QMetaObject) Setup(
	super *QMetaObject,
	className string,
	signals []*SignalDefinition,
	slots []*SlotDefinition,
	properties []*PropertyDefinition,
) {
	pg := util.NewPinGroup()

	dosSignals := make([]DosSignalDefinition, 0)
	for _, signal := range signals {
		dosSignals = append(dosSignals, signal.toDos(pg))
	}

	dosSlots := make([]DosSlotDefinition, 0)
	for _, slot := range slots {
		dosSlots = append(dosSlots, slot.toDos(pg))
	}

	dosProperties := make([]DosPropertyDefinition, 0)
	for _, property := range properties {
		dosProperties = append(dosProperties, property.toDos(pg))
	}

	meta.vptr = dos.QMetaObjectCreate(
		super.vptr,
		className,
		&DosSignalDefinitions{count: int32(len(dosSignals)), definitions: sliceToPtr(pg, dosSignals)},
		&DosSlotDefinitions{count: int32(len(dosSlots)), definitions: sliceToPtr(pg, dosSlots)},
		&DosPropertyDefinitions{count: int32(len(dosProperties)), definitions: sliceToPtr(pg, dosProperties)},
	)
}
