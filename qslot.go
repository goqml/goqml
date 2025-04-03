package goqml

type SlotDefinition struct {
	name           string
	returnMetaType QMetaType
	parameters     []ParameterDefinition
}

func (d *SlotDefinition) ToDos() DosSlotDefinition {
	parameters := make([]DosParameterDefinition, len(d.parameters))
	for i, param := range d.parameters {
		parameters[i] = param.ToDos()
	}
	return DosSlotDefinition{
		name:            stringToCharPtr(d.name),
		returnMetaType:  int(d.returnMetaType),
		parametersCount: len(parameters),
		parameters:      sliceToPtr(parameters),
	}
}
