package gconv

import (
	"github.com/ilylx/gconv/internal/gerror"
	"reflect"
)

// Scan automatically calls Struct or Structs function according to the type of parameter
// <pointer> to implement the converting.
// It calls function Struct if <pointer> is type of *struct/**struct to do the converting.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct to do the converting.
func Scan(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	t := reflect.TypeOf(pointer)
	k := t.Kind()
	if k != reflect.Ptr {
		return gerror.Newf("params should be type of pointer, but got: %v", k)
	}
	switch t.Elem().Kind() {
	case reflect.Array, reflect.Slice:
		return Structs(params, pointer, mapping...)
	default:
		return Struct(params, pointer, mapping...)
	}
}
