package util

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
)

func GetFuncName(fn reflect.Value) (string, error) {
	if fn.Kind() != reflect.Func {
		return "", errors.New("fn must be reflect.Func")
	}
	s := runtime.FuncForPC(fn.Pointer()).Name()
	list := strings.Split(s, ".")
	if len(list) == 0 {
		return "", errors.New("get func name error")
	}
	return list[len(list)-1], nil
}