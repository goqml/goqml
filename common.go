package goqml

import "reflect"

type (
	Ownership int
	QMetaType int
)

const (
	OwnershipTake Ownership = iota
	OwnershipClone

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

// 根据类型获取对应的 MetaType
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
	return DosParameterDefinition{name: stringToCharPtr(d.Name), metaType: int(d.MetaType)}
}
