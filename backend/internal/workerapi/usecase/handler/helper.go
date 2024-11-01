package handler

import (
	"reflect"

	"github.com/spf13/cast"
)

func castVariableType(varType string, value interface{}) (interface{}, error) {
	switch varType {
	case reflect.String.String():
		return cast.ToStringE(value)
	case reflect.Int32.String():
		return cast.ToInt32E(value)
	case reflect.Int64.String():
		return cast.ToInt64E(value)
	case reflect.Float32.String():
		return cast.ToFloat32E(value)
	case reflect.Float64.String():
		return cast.ToFloat64E(value)
	case reflect.Bool.String():
		return cast.ToBoolE(value)
	}

	return nil, nil
}
