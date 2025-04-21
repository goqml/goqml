package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/shapled/goqml"
)

func TestGenerateCodeContent(t *testing.T) {
	// 测试输入的结构体定义
	structs := []*StructDef{
		{
			Name:       "TestStruct",
			ParentName: "goqml.QObject",
			Slots: []*SlotDef{
				{
					Name:       "TestSlot",
					MethodName: "TestSlot",
					Params:     []*ParamDef{{Name: "x", Type: "int"}},
				},
			},
			Signals: []*SignalDef{
				{
					Name:       "TestSignal",
					MethodName: "TestSignal",
					Params:     []*ParamDef{{Name: "y", Type: "string"}},
				},
			},
			Properties: []*PropertyDef{
				{
					Name: "TestProperty",
					Type: goqml.QMetaTypeInt,
					Getter: &PropertyAccessor{
						Name:              "GetTestProperty",
						AnnotationType:    PropertyAnnotationTypeMethod,
						FieldOrMethodName: "GetTestProperty",
					},
					Setter: &PropertyAccessor{
						Name:              "SetTestProperty",
						AnnotationType:    PropertyAnnotationTypeMethod,
						FieldOrMethodName: "SetTestProperty",
					},
					Emitter: &PropertyAccessor{
						Name:              "TestPropertyChanged",
						AnnotationType:    PropertyAnnotationTypeMethod,
						FieldOrMethodName: "TestPropertyChanged",
					},
				},
			},
		},
	}

	// 调用 generateCodeContent 生成代码内容
	goContent, asmContent := generateCodeContent("testpkg", structs)

	fmt.Println(goContent)
	fmt.Println(asmContent)

	// 验证生成的 Go 代码内容
	expectedGoKeywords := []string{
		"package testpkg",
		"func (s *TestStruct) goqmlTestSignal(y string)",
		"func (s *TestStruct) goqmlTestPropertyChanged(v int)",
		"var staticTestStructQMetaObject = goqml.NewQMetaObject",
		"case \"TestSlot\":",
		"case \"GetTestProperty\":",
	}

	for _, keyword := range expectedGoKeywords {
		if !strings.Contains(goContent, keyword) {
			t.Errorf("Generated Go code does not contain expected keyword: %s", keyword)
		}
	}

	// 验证生成的汇编代码内容
	expectedAsmKeywords := []string{
		"TEXT ·TestStruct·TestSignal(SB), NOSPLIT, $0-16",
		"CALL ·TestStruct·goqmlTestSignal(SB)",
		"TEXT ·TestStruct·TestPropertyChanged(SB), NOSPLIT, $0-16",
		"CALL ·TestStruct·goqmlTestPropertyChanged(SB)",
	}

	for _, keyword := range expectedAsmKeywords {
		if !strings.Contains(asmContent, keyword) {
			t.Errorf("Generated assembly code does not contain expected keyword: %s", keyword)
		}
	}
}
