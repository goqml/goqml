package goqml

type SignalDefinition struct {
	name       string
	parameters []ParameterDefinition
}

func (d *SignalDefinition) ToDos() DosSignalDefinition {
	parameters := make([]DosParameterDefinition, len(d.parameters))
	for i, param := range d.parameters {
		parameters[i] = param.ToDos()
	}
	return DosSignalDefinition{
		name:            stringToCharPtr(d.name),
		parametersCount: len(parameters),
		parameters:      sliceToPtr(parameters),
	}
}
