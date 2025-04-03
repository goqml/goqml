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
	dosSignals := make([]DosSignalDefinition, len(signals))
	for _, signal := range signals {
		dosSignals = append(dosSignals, signal.ToDos())
	}

	dosSlots := make([]DosSlotDefinition, len(slots))
	for _, slot := range slots {
		dosSlots = append(dosSlots, slot.ToDos())
	}

	dosProperties := make([]DosPropertyDefinition, len(properties))
	for _, property := range properties {
		dosProperties = append(dosProperties, property.ToDos())
	}

	meta.vptr = dos.QMetaObjectCreate(
		super.vptr,
		className,
		sliceToPtr(dosSignals),
		sliceToPtr(dosSlots),
		sliceToPtr(dosProperties),
	)
}

func MakeSlot() *SlotDefinition {
	return nil
}
