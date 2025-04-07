package goqml

type SignalDefinition QFunc

func NewSignalDefinition(name string, arg any) *SignalDefinition {
	qFunc, err := NewQFunc(name, arg)
	if err != nil {
		panic(err)
	}
	if qFunc.retType != QMetaTypeVoid {
		panic("SignalDefinition: signals can't return values")
	}
	return (*SignalDefinition)(qFunc)
}

func (d *SignalDefinition) ToDos() DosSignalDefinition {
	parameters := make([]DosParameterDefinition, len(d.params))
	for i, param := range d.params {
		parameters[i] = param.ToDos()
	}
	return DosSignalDefinition{
		name:            stringToCharPtr(d.name),
		parametersCount: int32(len(parameters)),
		parameters:      sliceToPtr(parameters),
	}
}
