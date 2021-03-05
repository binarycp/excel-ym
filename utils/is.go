package utils

import (
	"reflect"
	"runtime"
)

func IsWindow() bool {
	return runtime.GOOS == "windows"
}

// 判断是否err类型
var errType = reflect.TypeOf((*error)(nil)).Elem()

func TypeIsError(vars ...interface{}) error {
	for _, v := range vars {
		if v != nil && reflect.TypeOf(v).Implements(errType) {
			return v.(error)
		}
	}
	return nil
}
