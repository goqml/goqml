package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/shapled/goqml"
)

func TestParseStructs(t *testing.T) {
	// 测试用例
	testCases := []struct {
		name     string
		input    string
		expected []*StructDef
	}{
		{
			name: "Simple struct with QObject",
			input: `
package main

import "github.com/shapled/goqml"

type MyStruct struct {
	goqml.QObject

	// @goqml.property
	MyProperty int
}

// @goqml.slot
func (*MyStruct) MySlot(x int) {}

// @goqml.signal
func (*MyStruct) MySignal(y string)
`,
			expected: []*StructDef{
				{
					Name:       "MyStruct",
					ParentType: "goqml.QObject",
					Slots: []*SlotDef{
						{
							StructName: "MyStruct",
							MethodName: "MySlot",
							Name:       "MySlot",
							Params:     []*ParamDef{{Name: "x", Type: "int"}},
						},
					},
					Signals: []*SignalDef{
						{
							StructName: "MyStruct",
							FieldName:  "MySignal",
							Name:       "MySignal",
							Params:     []*ParamDef{{Name: "y", Type: "string"}},
						},
					},
					Properties: []*PropertyDef{
						{
							StructName: "MyStruct",
							Name:       "MyProperty",
							Type:       goqml.QMetaTypeInt,
							Getter: &PropertyAccessor{
								Name:              "getMyProperty",
								AnnotationType:    PropertyAnnotationTypeField,
								FieldOrMethodName: "MyProperty",
							},
							Setter: &PropertyAccessor{
								Name:              "setMyProperty",
								AnnotationType:    PropertyAnnotationTypeField,
								FieldOrMethodName: "MyProperty",
							},
							Emitter: &PropertyAccessor{
								Name:              "MyPropertyChanged",
								AnnotationType:    PropertyAnnotationTypeField,
								FieldOrMethodName: "MyProperty",
							},
						},
					},
				},
			},
		},
		{
			name: "Struct with custom property names",
			input: `
package main

import "github.com/shapled/goqml"

type CustomStruct struct {
	goqml.QObject
	
	// @goqml.property("customPropertyName")
	MyProperty float64
}
	
// @goqml.slot("customSlotName")
func (*CustomStruct) MySlot(x int, y string) {}

// @goqml.signal("customSignalName")
func (*CustomStruct) MySignal(z bool)
`,
			expected: []*StructDef{
				{
					Name:       "CustomStruct",
					ParentType: "goqml.QObject",
					Slots: []*SlotDef{
						{
							StructName: "CustomStruct",
							MethodName: "MySlot",
							Name:       "customSlotName",
							Params:     []*ParamDef{{Name: "x", Type: "int"}, {Name: "y", Type: "string"}},
						},
					},
					Signals: []*SignalDef{
						{
							StructName: "CustomStruct",
							FieldName:  "MySignal",
							Name:       "customSignalName",
							Params:     []*ParamDef{{Name: "z", Type: "bool"}},
						},
					},
					Properties: []*PropertyDef{
						{
							StructName: "CustomStruct",
							Name:       "customPropertyName",
							Type:       goqml.QMetaTypeFloat,
							Getter: &PropertyAccessor{
								Name:              "getCustomPropertyName",
								AnnotationType:    PropertyAnnotationTypeField,
								FieldOrMethodName: "MyProperty",
							},
							Setter: &PropertyAccessor{
								Name:              "setCustomPropertyName",
								AnnotationType:    PropertyAnnotationTypeField,
								FieldOrMethodName: "MyProperty",
							},
							Emitter: &PropertyAccessor{
								Name:              "customPropertyNameChanged",
								AnnotationType:    PropertyAnnotationTypeField,
								FieldOrMethodName: "MyProperty",
							},
						},
					},
				},
			},
		},
		{
			name: "Struct with property getter, setter, and emitter",
			input: `
package main

import "github.com/shapled/goqml"

type PropertyStruct struct {
	goqml.QObject
}

// @goqml.property("p1").getter
func (s *PropertyStruct) GetP1() int {
	return 0
}

// @goqml.property("p1").setter
func (s *PropertyStruct) SetP1(value int) {}

// @goqml.property("p1").emitter
func (s *PropertyStruct) P1Changed(value int)
`,
			expected: []*StructDef{
				{
					Name:       "PropertyStruct",
					ParentType: "goqml.QObject",
					Properties: []*PropertyDef{
						{
							StructName: "PropertyStruct",
							Name:       "p1",
							Type:       goqml.QMetaTypeInt,
							Getter: &PropertyAccessor{
								Name:              "GetP1",
								AnnotationType:    PropertyAnnotationTypeMethod,
								FieldOrMethodName: "GetP1",
							},
							Setter: &PropertyAccessor{
								Name:              "SetP1",
								AnnotationType:    PropertyAnnotationTypeMethod,
								FieldOrMethodName: "SetP1",
							},
							Emitter: &PropertyAccessor{
								Name:              "P1Changed",
								AnnotationType:    PropertyAnnotationTypeMethod,
								FieldOrMethodName: "P1Changed",
							},
						},
					},
				},
			},
		},
		{
			name: "Struct with named getter, setter, and emitter",
			input: `
package main

import "github.com/shapled/goqml"

type NamedPropertyStruct struct {
	goqml.QObject
}

// @goqml.property("p1").getter("customGetter")
func (s *NamedPropertyStruct) CustomGetter() int {
	return 0
}

// @goqml.property("p1").setter("customSetter")
func (s *NamedPropertyStruct) CustomSetter(value int) {}

// @goqml.property("p1").emitter("customEmitter")
func (s *NamedPropertyStruct) CustomEmitter(value int)
`,
			expected: []*StructDef{
				{
					Name:       "NamedPropertyStruct",
					ParentType: "goqml.QObject",
					Properties: []*PropertyDef{
						{
							StructName: "NamedPropertyStruct",
							Name:       "p1",
							Type:       goqml.QMetaTypeInt,
							Getter: &PropertyAccessor{
								Name:              "customGetter",
								AnnotationType:    PropertyAnnotationTypeMethod,
								FieldOrMethodName: "CustomGetter",
							},
							Setter: &PropertyAccessor{
								Name:              "customSetter",
								AnnotationType:    PropertyAnnotationTypeMethod,
								FieldOrMethodName: "CustomSetter",
							},
							Emitter: &PropertyAccessor{
								Name:              "customEmitter",
								AnnotationType:    PropertyAnnotationTypeMethod,
								FieldOrMethodName: "CustomEmitter",
							},
						},
					},
				},
			},
		},
		{
			name: "Struct with only getter",
			input: `
package main

import "github.com/shapled/goqml"

type GetterStruct struct {
	goqml.QObject
}

// @goqml.property("p1").getter
func (s *GetterStruct) GetP1() int {
	return 0
}
`,
			expected: []*StructDef{
				{
					Name:       "GetterStruct",
					ParentType: "goqml.QObject",
					Properties: []*PropertyDef{
						{
							StructName: "GetterStruct",
							Name:       "p1",
							Type:       goqml.QMetaTypeInt,
							Getter: &PropertyAccessor{
								Name:              "GetP1",
								AnnotationType:    PropertyAnnotationTypeMethod,
								FieldOrMethodName: "GetP1",
							},
							Setter:  nil,
							Emitter: nil,
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "", tc.input, parser.ParseComments)
			if err != nil {
				t.Fatalf("Failed to parse input: %v", err)
			}

			result := parseStructs(file)

			diff := getStructDefDiff(result, tc.expected)
			if diff != "" {
				t.Errorf("parseStructs() mismatch:\n%s", diff)
			}
		})
	}
}

func getStructDefDiff(a, b []*StructDef) string {
	var diffs []string
	for i := range a {
		if a[i].Name != b[i].Name {
			diffs = append(diffs, fmt.Sprintf("Struct Name mismatch: got %s, want %s", a[i].Name, b[i].Name))
		}
		if a[i].ParentType != b[i].ParentType {
			diffs = append(diffs, fmt.Sprintf("Parent Name mismatch in struct %s: got %s, want %s", a[i].Name, a[i].ParentType, b[i].ParentType))
		}
		if slotDiff := getSlotDiff(a[i].Slots, b[i].Slots); slotDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Slots mismatch in struct %s:\n%s", a[i].Name, slotDiff))
		}
		if signalDiff := getSignalDiff(a[i].Signals, b[i].Signals); signalDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Signals mismatch in struct %s:\n%s", a[i].Name, signalDiff))
		}
		if propertyDiff := getPropertyDiff(a[i].Properties, b[i].Properties); propertyDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Properties mismatch in struct %s:\n%s", a[i].Name, propertyDiff))
		}
	}
	if len(diffs) > 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}

func getSlotDiff(a, b []*SlotDef) string {
	var diffs []string
	for i := range a {
		if a[i].StructName != b[i].StructName {
			diffs = append(diffs, fmt.Sprintf("Slot StructName mismatch: got %s, want %s", a[i].StructName, b[i].StructName))
		}
		if a[i].MethodName != b[i].MethodName {
			diffs = append(diffs, fmt.Sprintf("Slot MethodName mismatch: got %s, want %s", a[i].MethodName, b[i].MethodName))
		}
		if paramDiff := getParamDiff(a[i].Params, b[i].Params); paramDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Slot Params mismatch:\n%s", paramDiff))
		}
	}
	if len(diffs) > 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}

func getSignalDiff(a, b []*SignalDef) string {
	var diffs []string
	for i := range a {
		if a[i].StructName != b[i].StructName {
			diffs = append(diffs, fmt.Sprintf("Signal StructName mismatch: got %s, want %s", a[i].StructName, b[i].StructName))
		}
		if a[i].FieldName != b[i].FieldName {
			diffs = append(diffs, fmt.Sprintf("Signal MethodName mismatch: got %s, want %s", a[i].FieldName, b[i].FieldName))
		}
		if paramDiff := getParamDiff(a[i].Params, b[i].Params); paramDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Signal Params mismatch:\n%s", paramDiff))
		}
	}
	if len(diffs) > 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}

func getPropertyDiff(a, b []*PropertyDef) string {
	var diffs []string
	for i := range a {
		if a[i].StructName != b[i].StructName {
			diffs = append(diffs, fmt.Sprintf("Property StructName mismatch: got %s, want %s", a[i].StructName, b[i].StructName))
		}
		if a[i].Name != b[i].Name {
			diffs = append(diffs, fmt.Sprintf("Property Name mismatch: got %s, want %s", a[i].Name, b[i].Name))
		}
		if a[i].Type != b[i].Type {
			diffs = append(diffs, fmt.Sprintf("Property Type mismatch: got %s, want %s", a[i].Type, b[i].Type))
		}
		if accessorDiff := getAccessorDiff(a[i].Getter, b[i].Getter, "Getter"); accessorDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Property Getter mismatch:\n%s", accessorDiff))
		}
		if accessorDiff := getAccessorDiff(a[i].Setter, b[i].Setter, "Setter"); accessorDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Property Setter mismatch:\n%s", accessorDiff))
		}
		if accessorDiff := getAccessorDiff(a[i].Emitter, b[i].Emitter, "Emitter"); accessorDiff != "" {
			diffs = append(diffs, fmt.Sprintf("Property Emitter mismatch:\n%s", accessorDiff))
		}
	}
	if len(diffs) > 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}

func getAccessorDiff(a, b *PropertyAccessor, accessorType string) string {
	if a == nil && b == nil {
		return ""
	}
	if a != nil && b != nil {
		var diffs []string
		if a.Name != b.Name {
			diffs = append(diffs, fmt.Sprintf("%s Name mismatch: got %s, want %s", accessorType, a.Name, b.Name))
		}
		if a.AnnotationType != b.AnnotationType {
			diffs = append(diffs, fmt.Sprintf("%s AnnotationType mismatch: got %v, want %v", accessorType, a.AnnotationType, b.AnnotationType))
		}
		if a.FieldOrMethodName != b.FieldOrMethodName {
			diffs = append(diffs, fmt.Sprintf("%s FieldOrMethodName mismatch: got %s, want %s", accessorType, a.FieldOrMethodName, b.FieldOrMethodName))
		}
		if len(diffs) > 0 {
			return strings.Join(diffs, "\n")
		}
		return ""
	}
	return fmt.Sprintf("%s mismatch: one is nil, the other is not", accessorType)
}

func getParamDiff(a, b []*ParamDef) string {
	var diffs []string
	for i := range a {
		if i >= len(b) {
			diffs = append(diffs, fmt.Sprintf("Param index %d: expected param is missing", i))
			continue
		}
		if a[i].Name != b[i].Name {
			diffs = append(diffs, fmt.Sprintf("Param Name mismatch at index %d: got %s, want %s", i, a[i].Name, b[i].Name))
		}
		if a[i].Type != b[i].Type {
			diffs = append(diffs, fmt.Sprintf("Param Type mismatch at index %d: got %s, want %s", i, a[i].Type, b[i].Type))
		}
	}
	if len(b) > len(a) {
		for i := len(a); i < len(b); i++ {
			diffs = append(diffs, fmt.Sprintf("Param index %d: unexpected param found: %s %s", i, b[i].Name, b[i].Type))
		}
	}
	if len(diffs) > 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}
