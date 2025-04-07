package goqml

type SlotDefinition QFunc

func NewSlotDefinition(name string, arg any) *SlotDefinition {
	qFunc, err := NewQFunc(name, arg)
	if err != nil {
		panic(err)
	}
	return (*SlotDefinition)(qFunc)
}

func (d *SlotDefinition) ToDos() DosSlotDefinition {
	parameters := make([]DosParameterDefinition, len(d.params))
	for i, param := range d.params {
		parameters[i] = param.ToDos()
	}
	return DosSlotDefinition{
		name:            stringToCharPtr(d.name),
		returnMetaType:  int32(d.retType),
		parametersCount: int32(len(parameters)),
		parameters:      sliceToPtr(parameters),
	}
}
