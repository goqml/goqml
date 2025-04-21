package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/shapled/goqml"
)

const signalTemplate = `
TEXT ·{{.StructName}}·{{.Name}}(SB), NOSPLIT, $0-16
    CALL ·{{.StructName}}·goqml{{.Name}}(SB)
    RET 
`

func generateCode(pkgName string, structs []*StructDef, output string, force bool) {
	goContent, asmContent := generateCodeContent(pkgName, structs)

	if output == "" {
		output = "generated"
	}

	goFile := output + ".go"
	asmFile := output + ".s"

	if _, err := os.Stat(goFile); err == nil && !force {
		fmt.Printf("File %s already exists. Use -f to overwrite.\n", goFile)
		os.Exit(1)
	}

	if _, err := os.Stat(asmFile); err == nil && !force {
		fmt.Printf("File %s already exists. Use -f to overwrite.\n", asmFile)
		os.Exit(1)
	}

	os.WriteFile(goFile, []byte(goContent), 0644)

	// 如果有汇编内容，则生成汇编文件
	if asmContent != "" {
		os.WriteFile(asmFile, []byte(asmContent), 0644)
	}
}

func generateCodeContent(pkgName string, structs []*StructDef) (string, string) {
	tmplSignal := template.Must(template.New("signal").Funcs(template.FuncMap{
		"firstChar": func(s string) string { return strings.ToLower(s[:1]) },
	}).Parse(signalTemplate))

	var goBuilder, asmBuilder strings.Builder
	goBuilder.WriteString("package " + pkgName + "\n\n")
	goBuilder.WriteString("import (\n")
	goBuilder.WriteString("    \"fmt\"\n\n")
	goBuilder.WriteString("    \"github.com/shapled/goqml\"\n")
	goBuilder.WriteString(")\n\n")

	// 添加 #include "textflag.h" 到汇编代码头部
	asmBuilder.WriteString("#include \"textflag.h\"\n\n")

	hasAsmContent := false

	for _, s := range structs {
		// 生成 signal 方法
		for _, signal := range s.Signals {
			signalMethod := fmt.Sprintf("func (s *%s) goqml%s(%s) {\n", s.Name, signal.Name, generateSignalParams(signal))
			signalMethod += fmt.Sprintf("    s.Emit(\"%s\", %s)\n", signal.Name, generateSignalEmitParams(signal))
			signalMethod += "}\n\n"
			goBuilder.WriteString(signalMethod)

			// Plan9 汇编代码生成
			tmplSignal.Execute(&asmBuilder, signal)
			hasAsmContent = true
		}

		// 生成 property 的 Emitter 方法
		for _, prop := range s.Properties {
			if prop.Emitter.AnnotationType == PropertyAnnotationTypeMethod {
				emitterMethod := fmt.Sprintf("func (s *%s) goqml%s(%s) {\n", s.Name, prop.Emitter.FieldOrMethodName, "v "+prop.Type.GoTypeName())
				emitterMethod += fmt.Sprintf("    s.Emit(\"%sChanged\", %s)\n", prop.Name, "v "+prop.Type.GoTypeName())
				emitterMethod += "}\n\n"
				goBuilder.WriteString(emitterMethod)

				// Plan9 汇编代码生成
				tmplSignal.Execute(&asmBuilder, struct {
					StructName string
					Name       string
				}{
					StructName: s.Name,
					Name:       prop.Emitter.FieldOrMethodName,
				})
				hasAsmContent = true
			} else if prop.Emitter.AnnotationType == PropertyAnnotationTypeField {
				emitterMethod := fmt.Sprintf("func (s *%s) goqmlEmit%sSignal(value %s) {\n", s.Name, strings.Title(prop.Name), prop.Type.GoTypeName())
				emitterMethod += fmt.Sprintf("    s.Emit(\"%sChanged\", goqml.NewQVariant(value))\n", prop.Name)
				emitterMethod += "}\n\n"
				goBuilder.WriteString(emitterMethod)
			}
		}

		// 生成 QMetaObject 变量
		goBuilder.WriteString(fmt.Sprintf("var static%sQMetaObject = goqml.NewQMetaObject(\n", s.Name))
		goBuilder.WriteString(fmt.Sprintf("    (*%s)(nil).StaticMetaObject(),\n", s.ParentName))
		goBuilder.WriteString(fmt.Sprintf("    \"%s\",\n", s.Name))
		goBuilder.WriteString("    []*goqml.SignalDefinition{\n")
		for _, signal := range s.Signals {
			goBuilder.WriteString(fmt.Sprintf("        {\n"))
			goBuilder.WriteString(fmt.Sprintf("            Name: \"%s\",\n", signal.Name))
			goBuilder.WriteString(fmt.Sprintf("            Params: []*goqml.ParameterDefinition{\n"))
			for _, param := range signal.Params {
				goBuilder.WriteString(fmt.Sprintf("                {Name: \"%s\", Type: goqml.%s},\n", param.Name, goqml.GetMetaTypeStringFromTypeString(param.Type)))
			}
			goBuilder.WriteString(fmt.Sprintf("            },\n"))
			goBuilder.WriteString(fmt.Sprintf("        },\n"))
		}
		goBuilder.WriteString("    },\n")
		goBuilder.WriteString("    []*goqml.SlotDefinition{\n")
		for _, slot := range s.Slots {
			goBuilder.WriteString(fmt.Sprintf("        {\n"))
			goBuilder.WriteString(fmt.Sprintf("            Name: \"%s\",\n", slot.Name))
			goBuilder.WriteString(fmt.Sprintf("            RetMetaType: goqml.%s,\n", goqml.GetMetaTypeStringFromTypeString(slot.ReturnType)))
			goBuilder.WriteString(fmt.Sprintf("            Params: []*goqml.ParameterDefinition{\n"))
			for _, param := range slot.Params {
				goBuilder.WriteString(fmt.Sprintf("                {Name: \"%s\", Type: goqml.%s},\n", param.Name, goqml.GetMetaTypeStringFromTypeString(param.Type)))
			}
			goBuilder.WriteString(fmt.Sprintf("            },\n"))
			goBuilder.WriteString(fmt.Sprintf("        },\n"))
		}
		goBuilder.WriteString("    },\n")
		goBuilder.WriteString("    []*goqml.PropertyDefinition{\n")
		for _, prop := range s.Properties {
			goBuilder.WriteString(fmt.Sprintf("        {\n"))
			goBuilder.WriteString(fmt.Sprintf("            Name: \"%s\",\n", prop.Name))
			goBuilder.WriteString(fmt.Sprintf("            MetaType: goqml.%s,\n", goqml.GetMetaTypeStringFromTypeString(prop.Type.GoTypeName())))
			goBuilder.WriteString(fmt.Sprintf("            Getter: \"%s\",\n", prop.Getter.Name))
			goBuilder.WriteString(fmt.Sprintf("            Setter: \"%s\",\n", prop.Setter.Name))
			goBuilder.WriteString(fmt.Sprintf("            Emitter: \"%s\",\n", prop.Emitter.Name))
			goBuilder.WriteString(fmt.Sprintf("        },\n"))
		}
		goBuilder.WriteString("    },\n")
		goBuilder.WriteString(")\n\n")

		// 生成 StaticMetaObject 方法
		goBuilder.WriteString(fmt.Sprintf("func (s *%s) StaticMetaObject() *goqml.QMetaObject {\n", s.Name))
		goBuilder.WriteString(fmt.Sprintf("    return static%sQMetaObject\n", s.Name))
		goBuilder.WriteString("}\n\n")

		// 生成 OnSlotCalled 方法
		goBuilder.WriteString(fmt.Sprintf("func (s *%s) OnSlotCalled(slotName string, arguments []*goqml.QVariant) {\n", s.Name))
		goBuilder.WriteString("    switch slotName {\n")
		for _, slot := range s.Slots {
			goBuilder.WriteString(fmt.Sprintf("    case \"%s\":\n", slot.Name))
			goBuilder.WriteString(fmt.Sprintf("        s.%s(%s)\n", slot.MethodName, generateSlotArguments(slot)))
		}
		for _, prop := range s.Properties {
			if prop.Getter != nil {
				goBuilder.WriteString(fmt.Sprintf("    case \"%s\":\n", prop.Getter.Name))
				switch prop.Getter.AnnotationType {
				case PropertyAnnotationTypeMethod:
					goBuilder.WriteString(fmt.Sprintf("        arguments[0].SetVal(s.%s())\n", prop.Getter.FieldOrMethodName))
				case PropertyAnnotationTypeField:
					goBuilder.WriteString(fmt.Sprintf("        arguments[0].SetVal(s.%s)\n", prop.Getter.FieldOrMethodName))
				default:
					panic("invalid property annotation type")
				}
			}
			if prop.Setter != nil {
				goBuilder.WriteString(fmt.Sprintf("    case \"%s\":\n", prop.Setter.Name))
				switch prop.Getter.AnnotationType {
				case PropertyAnnotationTypeMethod:
					goBuilder.WriteString(fmt.Sprintf("        s.%s = arguments[1].To%s()\n", prop.Setter.FieldOrMethodName, strings.Title(prop.Type.GoTypeName())))
					goBuilder.WriteString(fmt.Sprintf("        s.goqml%s(arguments[1].To%s())\n", prop.Emitter.FieldOrMethodName, strings.Title(prop.Type.GoTypeName())))
				case PropertyAnnotationTypeField:
					goBuilder.WriteString(fmt.Sprintf("        v := arguments[1].%s()\n", prop.Type.QVariantGetterName()))
					goBuilder.WriteString(fmt.Sprintf("        if s.%s != v {\n", prop.Setter.FieldOrMethodName))
					goBuilder.WriteString(fmt.Sprintf("            s.%s = v\n", prop.Setter.FieldOrMethodName))
					goBuilder.WriteString(fmt.Sprintf("            s.goqmlEmit%sSignal(v)\n", strings.Title(prop.Name)))
					goBuilder.WriteString(fmt.Sprintf("        }\n"))
				default:
					panic("invalid property annotation type")
				}
			}
		}
		goBuilder.WriteString("    default:\n")
		goBuilder.WriteString("        fmt.Println(\"unknown slot:\", slotName)\n")
		goBuilder.WriteString("    }\n")
		goBuilder.WriteString("}\n\n")
	}

	// 如果没有汇编内容，返回空字符串
	if !hasAsmContent {
		return goBuilder.String(), ""
	}

	return goBuilder.String(), asmBuilder.String()
}

func generateSignalParams(signal *SignalDef) string {
	params := []string{}
	for _, param := range signal.Params {
		params = append(params, fmt.Sprintf("%s %s", param.Name, param.Type))
	}
	return strings.Join(params, ", ")
}

func generateSignalEmitParams(signal *SignalDef) string {
	params := []string{}
	for _, param := range signal.Params {
		params = append(params, fmt.Sprintf("goqml.NewQVariant(%s)", param.Name))
	}
	return strings.Join(params, ", ")
}

func generateQMetaTypes(params []*ParamDef) string {
	types := []string{}
	for _, param := range params {
		types = append(types, fmt.Sprintf("goqml.QMetaType%s", strings.Title(param.Type)))
	}
	if len(types) > 0 {
		return ", " + strings.Join(types, ", ")
	}
	return ""
}

func generateQMetaType(typeName string) string {
	return "goqml." + goqml.GetMetaTypeStringFromTypeString(typeName)
}

func generateSlotArguments(slot *SlotDef) string {
	args := []string{}
	for i, param := range slot.Params {
		args = append(args, fmt.Sprintf("arguments[%d].To%s()", i+1, strings.Title(param.Type)))
	}
	return strings.Join(args, ", ")
}
