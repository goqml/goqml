package goqml

import (
	"fmt"
	"reflect"
)

type PropertyDefinition struct {
	name             string
	propertyMetaType QMetaType
	read             *SlotDefinition
	write            *SlotDefinition
	notify           *SignalDefinition
}

func NewPropertyDefinition[T any](name string, read func() T, write func(T), notify func(T)) *PropertyDefinition {
	var readFunc, writeFunc *SlotDefinition
	var notifyFunc *SignalDefinition
	var t T

	if read != nil {
		readFunc = NewSlotDefinition(fmt.Sprintf("read_%s", name), read)
	}
	if write != nil {
		writeFunc = NewSlotDefinition(fmt.Sprintf("write_%s", name), write)
	}
	if notify != nil {
		notifyFunc = NewSignalDefinition(fmt.Sprintf("notify_%s", name), notify)
	}

	return &PropertyDefinition{
		name:             name,
		propertyMetaType: getMetaType(reflect.TypeOf(t)),
		read:             readFunc,
		write:            writeFunc,
		notify:           notifyFunc,
	}
}

func (d *PropertyDefinition) names() (string, string, string) {
	readFuncName, writeFuncName, notifyFuncName := "", "", ""
	if d.read != nil {
		readFuncName = d.read.name
	}
	if d.write != nil {
		writeFuncName = d.write.name
	}
	if d.notify != nil {
		notifyFuncName = d.notify.name
	}
	return readFuncName, writeFuncName, notifyFuncName
}

func (d *PropertyDefinition) ToDos() DosPropertyDefinition {
	readSlot, writeSlot, notifySignal := d.names()
	return DosPropertyDefinition{
		name:             stringToCharPtr(d.name),
		propertyMetaType: int32(d.propertyMetaType),
		readSlot:         stringToCharPtr(readSlot),
		writeSlot:        stringToCharPtr(writeSlot),
		notifySignal:     stringToCharPtr(notifySignal),
	}
}
