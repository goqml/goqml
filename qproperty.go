package goqml

type PropertyDefinition struct {
	name             string
	propertyMetaType QMetaType
	readSlot         string
	writeSlot        string
	notifySignal     string
}

func (d *PropertyDefinition) ToDos() DosPropertyDefinition {
	return DosPropertyDefinition{
		name:             stringToCharPtr(d.name),
		propertyMetaType: int(d.propertyMetaType),
		readSlot:         stringToCharPtr(d.readSlot),
		writeSlot:        stringToCharPtr(d.writeSlot),
		notifySignal:     stringToCharPtr(d.notifySignal),
	}
}
