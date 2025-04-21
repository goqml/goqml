package goqml

type ParameterDefinition struct {
	Name     string
	MetaType QMetaType
}

func (d *ParameterDefinition) toDos() DosParameterDefinition {
	return DosParameterDefinition{name: stringToCharPtr(d.Name), metaType: int32(d.MetaType)}
}

type SlotDefinition struct {
	Name        string
	RetMetaType QMetaType
	Params      []*ParameterDefinition
}

func (d *SlotDefinition) toDos() DosSlotDefinition {
	parameters := make([]DosParameterDefinition, len(d.Params))
	for i, param := range d.Params {
		parameters[i] = param.toDos()
	}
	return DosSlotDefinition{
		name:            stringToCharPtr(d.Name),
		returnMetaType:  int32(d.RetMetaType),
		parametersCount: int32(len(parameters)),
		parameters:      sliceToPtr(parameters),
	}
}

type SignalDefinition struct {
	Name   string
	Params []*ParameterDefinition
}

func (d *SignalDefinition) toDos() DosSignalDefinition {
	parameters := make([]DosParameterDefinition, len(d.Params))
	for i, param := range d.Params {
		parameters[i] = param.toDos()
	}
	return DosSignalDefinition{
		name:            stringToCharPtr(d.Name),
		parametersCount: int32(len(parameters)),
		parameters:      sliceToPtr(parameters),
	}
}

type PropertyDefinition struct {
	Name     string
	MetaType QMetaType
	Getter   string
	Setter   string
	Emitter  string
}

func (d *PropertyDefinition) toDos() DosPropertyDefinition {
	return DosPropertyDefinition{
		name:             stringToCharPtr(d.Name),
		propertyMetaType: int32(d.MetaType),
		readSlot:         stringToCharPtr(d.Getter),
		writeSlot:        stringToCharPtr(d.Setter),
		notifySignal:     stringToCharPtr(d.Emitter),
	}
}
