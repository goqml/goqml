package util

import (
	"reflect"
)

type PinGroup struct {
	pins map[uintptr]any
}

func NewPinGroup() *PinGroup {
	return &PinGroup{
		pins: make(map[uintptr]any),
	}
}

func (pg *PinGroup) Pin(ptr any) uintptr {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Slice && v.Kind() != reflect.Map {
		panic("PinGroup: only pointer, slice, and map types are allowed")
	}

	addr := v.Pointer()
	pg.pins[addr] = ptr
	return addr
}

func (pg *PinGroup) Unpin(ptr any) {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Slice && v.Kind() != reflect.Map {
		panic("PinGroup: only pointer, slice, and map types are allowed")
	}

	addr := v.Pointer()
	delete(pg.pins, addr)
}

func (pg *PinGroup) Pinned(addr uintptr) bool {
	_, exists := pg.pins[addr]
	return exists
}

var pinGroup = NewPinGroup()

func Pin(ptr any) uintptr {
	return pinGroup.Pin(ptr)
}

func Unpin(ptr any) {
	pinGroup.Unpin(ptr)
}

func Pinned(addr uintptr) bool {
	return pinGroup.Pinned(addr)
}
