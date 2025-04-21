package main

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"github.com/shapled/goqml"
)

func parseStructs(node *ast.File) []*StructDef {
	var structs []*StructDef
	structMap := make(map[string]*StructDef)

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if structType, ok := x.Type.(*ast.StructType); ok {
				structDef := &StructDef{Name: x.Name.Name, ParentName: fmt.Sprintf("goqml.QObject[*%s]", x.Name.Name)}
				structMap[x.Name.Name] = structDef

				// Find parent struct and parse properties
				parentFound := false
				for _, field := range structType.Fields.List {
					if !parentFound {
						if len(field.Names) == 0 {
							if ident, ok := field.Type.(*ast.Ident); ok {
								structDef.ParentName = ident.Name
								parentFound = true
							} else if indexExpr, ok := field.Type.(*ast.IndexExpr); ok {
								if ident, ok := indexExpr.X.(*ast.Ident); ok {
									genericType := getTypeName(indexExpr.Index)
									structDef.ParentName = fmt.Sprintf("%s[%s]", ident.Name, genericType)
									parentFound = true
								} else if selectorExpr, ok := indexExpr.X.(*ast.SelectorExpr); ok {
									genericType := getTypeName(indexExpr.Index)
									structDef.ParentName = fmt.Sprintf("%s.%s[%s]", getTypeName(selectorExpr.X), selectorExpr.Sel.Name, genericType)
									parentFound = true
								}
							} else if selectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {
								// 新增：处理 SelectorExpr 类型
								structDef.ParentName = fmt.Sprintf("%s.%s", getTypeName(selectorExpr.X), selectorExpr.Sel.Name)
								parentFound = true
							}
						}
					}
					if field.Doc != nil {
						for _, c := range field.Doc.List {
							// 修改：去除 // 前缀后匹配注解
							commentText := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
							if strings.HasPrefix(commentText, "@goqml.property") {
								structDef.Properties = append(structDef.Properties, parseFieldPropertyDef(commentText, x.Name.Name, field))
							}
						}
					}
				}
			}
		case *ast.FuncDecl:
			if x.Recv != nil && len(x.Recv.List) > 0 {
				if starExpr, ok := x.Recv.List[0].Type.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok {
						if structDef, exists := structMap[ident.Name]; exists {
							if x.Doc != nil {
								for _, c := range x.Doc.List {
									commentText := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
									if strings.HasPrefix(commentText, "@goqml.slot") {
										structDef.Slots = append(structDef.Slots, parseSlotDef(commentText, ident.Name, x))
									} else if strings.HasPrefix(commentText, "@goqml.signal") {
										structDef.Signals = append(structDef.Signals, parseSignalDef(commentText, ident.Name, x))
									} else if strings.HasPrefix(commentText, "@goqml.property") {
										def := parseMethodPropertyDef(commentText, ident.Name, x)
										existingProperty := findPropertyByName(structDef.Properties, def.Name)
										if existingProperty == nil {
											structDef.Properties = append(structDef.Properties, def)
										} else {
											updateProperty(existingProperty, def)
										}
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	for _, structDef := range structMap {
		structs = append(structs, structDef)
	}

	return structs
}

// 新增：查找已存在的 property
func findPropertyByName(properties []*PropertyDef, name string) *PropertyDef {
	for i := range properties {
		if properties[i].Name == name {
			return properties[i]
		}
	}
	return nil
}

func updateProperty(existingProperty *PropertyDef, property *PropertyDef) {
	if existingProperty.Type != property.Type {
		panic("Property type mismatch")
	}
	if property.Getter != nil {
		if existingProperty.Getter != nil && existingProperty.Getter.AnnotationType == PropertyAnnotationTypeMethod {
			panic("Getter already defined")
		}
		existingProperty.Getter = property.Getter
	}
	if property.Setter != nil {
		if existingProperty.Setter != nil && existingProperty.Setter.AnnotationType == PropertyAnnotationTypeMethod {
			panic("Setter already defined")
		}
		existingProperty.Setter = property.Setter
	}
	if property.Emitter != nil {
		if existingProperty.Emitter != nil && existingProperty.Emitter.AnnotationType == PropertyAnnotationTypeMethod {
			panic("Emitter already defined")
		}
		existingProperty.Emitter = property.Emitter
	}
}

func parseSlotDef(comment string, structName string, funcDecl *ast.FuncDecl) *SlotDef {
	re := regexp.MustCompile(`@goqml\.slot\s*(\("?(.*?)"?\))?`)
	match := re.FindStringSubmatch(comment)
	name := funcDecl.Name.Name
	if len(match) > 2 && match[2] != "" {
		name = match[2]
	}

	params := []*ParamDef{}
	returnType := ""

	funcType := funcDecl.Type
	for _, param := range funcType.Params.List {
		paramType := ""
		if ident, ok := param.Type.(*ast.Ident); ok {
			paramType = ident.Name
		}
		for _, paramName := range param.Names {
			params = append(params, &ParamDef{Name: paramName.Name, Type: paramType})
		}
	}

	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		if ident, ok := funcType.Results.List[0].Type.(*ast.Ident); ok {
			returnType = ident.Name
		}
	}

	return &SlotDef{
		StructName: structName,
		MethodName: funcDecl.Name.Name,
		Name:       name,
		Params:     params,
		ReturnType: returnType,
	}
}

func parseSignalDef(comment string, structName string, funcDecl *ast.FuncDecl) *SignalDef {
	re := regexp.MustCompile(`@goqml\.signal\s*(\("(.*?)"\))?`)
	match := re.FindStringSubmatch(comment)
	name := funcDecl.Name.Name
	if len(match) > 2 && match[2] != "" {
		name = match[2]
	}

	params := []*ParamDef{}
	funcType := funcDecl.Type

	for _, param := range funcType.Params.List {
		paramType := ""
		if ident, ok := param.Type.(*ast.Ident); ok {
			paramType = ident.Name
		}
		for _, paramName := range param.Names {
			params = append(params, &ParamDef{Name: paramName.Name, Type: paramType})
		}
	}

	return &SignalDef{
		StructName: structName,
		MethodName: funcDecl.Name.Name,
		Name:       name,
		Params:     params,
	}
}

func parseFieldPropertyDef(comment string, structName string, field *ast.Field) *PropertyDef {
	re := regexp.MustCompile(`@goqml\.property\s*(\("(.*?)"\))?`)
	match := re.FindStringSubmatch(comment)
	fieldName := field.Names[0].Name
	name := fieldName
	if len(match) > 2 && match[2] != "" {
		name = match[2]
	}

	propertyType := goqml.QMetaTypeUnknownType
	if ident, ok := field.Type.(*ast.Ident); ok {
		propertyType = goqml.NewQMetaType(ident.Name)
	}

	return &PropertyDef{
		StructName: structName,
		Name:       name,
		Type:       propertyType,
		Getter: &PropertyAccessor{
			Name:              "get" + strings.Title(name),
			AnnotationType:    PropertyAnnotationTypeField,
			FieldOrMethodName: fieldName,
		},
		Setter: &PropertyAccessor{
			Name:              "set" + strings.Title(name),
			AnnotationType:    PropertyAnnotationTypeField,
			FieldOrMethodName: fieldName,
		},
		Emitter: &PropertyAccessor{
			Name:              name + "Changed",
			AnnotationType:    PropertyAnnotationTypeField,
			FieldOrMethodName: fieldName,
		},
	}
}

func parseMethodPropertyDef(comment string, structName string, funcDecl *ast.FuncDecl) *PropertyDef {
	re := regexp.MustCompile(`@goqml\.property\s*\("?(.*?)"?\)\.(getter|setter|emitter)(\("?(.*?)"?\))?`)
	match := re.FindStringSubmatch(comment)
	if len(match) < 5 {
		panic("Invalid property annotation")
	}

	propertyName := match[1]
	accessorType := match[2]

	accessorName := funcDecl.Name.Name
	if len(match) > 4 && match[4] != "" {
		accessorName = match[4]
	}

	accessor := &PropertyAccessor{
		Name:              accessorName,
		AnnotationType:    PropertyAnnotationTypeMethod,
		FieldOrMethodName: funcDecl.Name.Name,
	}

	def := &PropertyDef{
		StructName: structName,
		Name:       propertyName,
	}

	funcType := funcDecl.Type
	switch accessorType {
	case "getter":
		if funcType.Params != nil && len(funcType.Params.List) != 0 {
			panic("getter must have no parameters")
		}
		if funcType.Results == nil || len(funcType.Results.List) != 1 {
			panic("getter must return one result")
		}
		typeNode := funcType.Results.List[0].Type
		def.Type = getMetaTypeName(typeNode)
		def.Getter = accessor
	case "setter":
		if funcType.Params == nil || len(funcType.Params.List) != 1 {
			panic("setter must have one parameter")
		}
		if funcType.Results != nil && len(funcType.Results.List) != 0 {
			panic("setter must return nothing")
		}
		typeNode := funcType.Params.List[0].Type
		def.Type = getMetaTypeName(typeNode)
		def.Setter = accessor
	case "emitter":
		if funcType.Params == nil || len(funcType.Params.List) != 1 {
			panic("emitter must have one parameter")
		}
		if funcType.Results != nil && len(funcType.Results.List) != 0 {
			panic("emitter must return nothing")
		}
		typeNode := funcType.Params.List[0].Type
		def.Type = getMetaTypeName(typeNode)
		def.Emitter = accessor
	default:
		panic("unsupported property type")
	}

	return def
}

func getTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return t.Sel.Name
	case *ast.StarExpr:
		return "*" + getTypeName(t.X)
	default:
		panic("unsupported type expression")
	}
}

func getMetaTypeName(expr ast.Expr) goqml.QMetaType {
	return goqml.NewQMetaType(getTypeName(expr))
}
