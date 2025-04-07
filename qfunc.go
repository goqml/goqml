package goqml

import (
	"errors"
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

func getMetaType(t reflect.Type) QMetaType {
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

type ParameterDefinition struct {
	Name     string
	MetaType QMetaType
}

func (d *ParameterDefinition) ToDos() DosParameterDefinition {
	return DosParameterDefinition{name: stringToCharPtr(d.Name), metaType: int32(d.MetaType)}
}

type QFunc struct {
	name     string
	callback reflect.Value
	params   []*ParameterDefinition
	retType  QMetaType
}

func NewQFunc(name string, f any) (*QFunc, error) {
	qFunc := &QFunc{name: name, callback: reflect.ValueOf(f)}

	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return nil, errors.New("not a function")
	}

	if t.NumOut() > 1 {
		return nil, errors.New("only zero or one return value is supported")
	}

	for i := 0; i < t.NumIn(); i++ {
		qFunc.params = append(qFunc.params, &ParameterDefinition{
			Name:     fmt.Sprintf("%s_%d", t.In(i).Name(), i),
			MetaType: getMetaType(t.In(i)),
		})
	}

	if t.NumOut() == 0 {
		qFunc.retType = QMetaTypeVoid
	} else {
		qFunc.retType = getMetaType(t.Out(0))
	}

	return qFunc, nil
}

func (qFunc *QFunc) applyQVariants(arguments []*QVariant) error {
	args := make([]reflect.Value, 0)
	for i, param := range qFunc.params {
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
			return fmt.Errorf("unsupported parameter type: %d", param.MetaType)
		}
	}
	results := qFunc.callback.Call(args)
	switch qFunc.retType {
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
		return fmt.Errorf("unsupported return type: %d", qFunc.retType)
	}
	return nil
}
