package goqml

type QMetaObject struct {
	vptr       DosQMetaObject
	signals    []*SignalDefinition
	slots      []*SlotDefinition
	properties []*PropertyDefinition
}

func NewQObjectMetaObject() *QMetaObject {
	return &QMetaObject{vptr: dos.QObjectQMetaObject()}
}

func NewQMetaObject(
	super *QMetaObject,
	className string,
	signals []*SignalDefinition,
	slots []*SlotDefinition,
	properties []*PropertyDefinition,
) *QMetaObject {
	for _, property := range properties {
		if property.read != nil {
			slots = append(slots, property.read)
		}
		if property.write != nil {
			slots = append(slots, property.write)
		}
		if property.notify != nil {
			signals = append(signals, property.notify)
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
	dosSignals := make([]DosSignalDefinition, 0)
	for _, signal := range signals {
		dosSignals = append(dosSignals, signal.ToDos())
	}

	dosSlots := make([]DosSlotDefinition, 0)
	for _, slot := range slots {
		dosSlots = append(dosSlots, slot.ToDos())
	}

	dosProperties := make([]DosPropertyDefinition, 0)
	for _, property := range properties {
		dosProperties = append(dosProperties, property.ToDos())
	}

	meta.vptr = dos.QMetaObjectCreate(
		super.vptr,
		className,
		&DosSignalDefinitions{count: int32(len(dosSignals)), definitions: sliceToPtr(dosSignals)},
		&DosSlotDefinitions{count: int32(len(dosSlots)), definitions: sliceToPtr(dosSlots)},
		&DosPropertyDefinitions{count: int32(len(dosProperties)), definitions: sliceToPtr(dosProperties)},
	)

	releaseBytes()
}

func (meta *QMetaObject) OnSlotCalled(slotName string, arguments []*QVariant) {
	for _, slot := range meta.slots {
		if slot.name == slotName {
			(*QFunc)(slot).applyQVariants(arguments)
			return
		}
	}
}
