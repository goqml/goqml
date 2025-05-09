package goqml

import (
	"fmt"
	"reflect"
)

type QMetaType int

const (
	QMetaTypeUnknownType QMetaType = 0
	QMetaTypeBool        QMetaType = 1
	QMetaTypeInt         QMetaType = 2
	QMetaTypeQString     QMetaType = 10
	QMetaTypeVoidStar    QMetaType = 31
	QMetaTypeFloat       QMetaType = 38
	QMetaTypeQObjectStar QMetaType = 39
	QMetaTypeQVariant    QMetaType = 41
	QMetaTypeVoid        QMetaType = 43
)

func NewQMetaType(t string) QMetaType {
	switch t {
	case "bool":
		return QMetaTypeBool
	case "int":
		return QMetaTypeInt
	case "float32":
		return QMetaTypeFloat
	case "string":
		return QMetaTypeQString
	case "*QObject":
		return QMetaTypeQObjectStar
	case "QVariant":
		return QMetaTypeQVariant
	case "void":
		return QMetaTypeVoid
	default:
		panic("unsupported QMetaType")
	}
}

func (t QMetaType) GoTypeName() string {
	switch t {
	case QMetaTypeBool:
		return "bool"
	case QMetaTypeInt:
		return "int"
	case QMetaTypeFloat:
		return "float32"
	case QMetaTypeQString:
		return "string"
	case QMetaTypeQObjectStar:
		return "*QObject"
	case QMetaTypeQVariant:
		return "QVariant"
	case QMetaTypeVoid:
		return "void"
	default:
		panic("unsupported QMetaType")
	}
}

func (t QMetaType) QVariantGetterName() string {
	switch t {
	case QMetaTypeBool:
		return "BoolVal"
	case QMetaTypeInt:
		return "IntVal"
	case QMetaTypeFloat:
		return "FloatVal"
	case QMetaTypeQString:
		return "StringVal"
	case QMetaTypeQObjectStar:
		return "QObject"
	case QMetaTypeQVariant:
		return "QVariant"
	case QMetaTypeVoid:
		fallthrough
	default:
		panic("unsupported QMetaType")
	}
}

func GetMetaTypeStringFromTypeString(t string) string {
	switch t {
	case "bool":
		return "QMetaTypeBool"
	case "int", "int8", "int16", "int32", "int64":
		return "QMetaTypeInt"
	case "float32", "float64":
		return "QMetaTypeFloat"
	case "string":
		return "QMetaTypeQString"
	case "*QObject":
		return "QMetaTypeQObjectStar"
	case "QVariant":
		return "QMetaTypeQVariant"
	case "void", "":
		return "QMetaTypeVoid"
	default:
		return "QMetaTypeUnknownType"
	}
}

func GetMetaTypeFromReflectType(t reflect.Type) QMetaType {
	switch t.Kind() {
	case reflect.Bool:
		return QMetaTypeBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return QMetaTypeInt
	case reflect.Float32, reflect.Float64:
		return QMetaTypeFloat
	case reflect.String:
		return QMetaTypeQString
	case reflect.Ptr:
		if t.Elem().Kind() == reflect.Struct {
			return QMetaTypeQObjectStar
		}
		return QMetaTypeVoidStar
	case reflect.Interface:
		return QMetaTypeQVariant
	case reflect.Invalid:
		return QMetaTypeVoid
	default:
		return QMetaTypeUnknownType
	}
}

func ApplyAndAssignQVariants(f any, arguments []*QVariant) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("not a function")
	}

	if t.NumOut() > 1 {
		panic("only zero or one return value is supported")
	}

	params := make([]*ParameterDefinition, 0)
	for i := 0; i < t.NumIn(); i++ {
		params = append(params, &ParameterDefinition{
			Name:     fmt.Sprintf("%s_%d", t.In(i).Name(), i),
			MetaType: GetMetaTypeFromReflectType(t.In(i)),
		})
	}

	args := make([]reflect.Value, 0)
	for i, param := range params {
		switch param.MetaType {
		case QMetaTypeBool:
			args = append(args, reflect.ValueOf(arguments[i+1].BoolVal()))
		case QMetaTypeInt:
			args = append(args, reflect.ValueOf(arguments[i+1].IntVal()))
		case QMetaTypeFloat:
			args = append(args, reflect.ValueOf(arguments[i+1].FloatVal()))
		case QMetaTypeQString:
			args = append(args, reflect.ValueOf(arguments[i+1].StringVal()))
		default:
			panic(fmt.Errorf("unsupported parameter type: %d", param.MetaType))
		}
	}

	results := reflect.ValueOf(f).Call(args)
	retType := QMetaTypeVoid
	if t.NumOut() != 0 {
		retType = GetMetaTypeFromReflectType(t.Out(0))
	}

	switch retType {
	case QMetaTypeVoid:
	case QMetaTypeBool:
		arguments[0].SetBoolVal(results[0].Bool())
	case QMetaTypeInt:
		arguments[0].SetIntVal(int(results[0].Int()))
	case QMetaTypeFloat:
		arguments[0].SetFloatVal(float32(results[0].Float()))
	case QMetaTypeQString:
		arguments[0].SetStringVal(results[0].String())
	default:
		panic(fmt.Errorf("unsupported return type: %d", retType))
	}
}
