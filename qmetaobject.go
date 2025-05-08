package goqml

type QMetaObject struct {
	vptr       DosQMetaObject
	signals    []*SignalDefinition
	slots      []*SlotDefinition
	properties []*PropertyDefinition
}

func NewQObjectMetaObject() *QMetaObject {
	return &QMetaObject{vptr: DosQObjectQMetaObject()}
}

func NewQAbstractItemModelMetaObject() *QMetaObject {
	return &QMetaObject{vptr: DosQAbstractItemModelQMetaObject()}
}

func NewQAbstractListModelMetaObject() *QMetaObject {
	return &QMetaObject{vptr: DosQAbstractListModelQMetaObject()}
}

func NewQAbstractTableModelMetaObject() *QMetaObject {
	return &QMetaObject{vptr: DosQAbstractTableModelQMetaObject()}
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
	meta.vptr = DosQMetaObjectCreate(
		super.vptr,
		className,
		signals,
		slots,
		properties,
	)
}
