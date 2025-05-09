package main

import "github.com/shapled/goqml"

type (
	PropertyAnnotationType int
)

const (
	PropertyAnnotationTypeField PropertyAnnotationType = iota
	PropertyAnnotationTypeMethod
)

type StructDef struct {
	Name       string
	ParentType string
	ParentName string
	Slots      []*SlotDef
	Signals    []*SignalDef
	Properties []*PropertyDef
}

type SlotDef struct {
	StructName string
	MethodName string
	Name       string
	Params     []*ParamDef
	ReturnType string
}

type SignalDef struct {
	StructName string
	FieldName  string
	Name       string
	Params     []*ParamDef
}

type PropertyAccessor struct {
	Name              string
	AnnotationType    PropertyAnnotationType
	FieldOrMethodName string
}

func (pa *PropertyAccessor) NameOrEmpty() string {
	if pa == nil {
		return ""
	}
	return pa.Name
}

type PropertyDef struct {
	StructName string
	Name       string
	Type       goqml.QMetaType
	Getter     *PropertyAccessor
	Setter     *PropertyAccessor
	Emitter    *PropertyAccessor
}

type ParamDef struct {
	Name string
	Type string
}
