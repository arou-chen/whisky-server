package tcp

import "reflect"

type Method struct {
	Name     string
	Handler  func(in interface{}) (out interface{}, err error)
	In       reflect.Type
	Out      reflect.Type
}

func NewMethod(
	name string,
	handler func(in interface{}) (out interface{}, err error),
	in reflect.Type,
	out reflect.Type) Method {
	return Method{
		Name:	 name,
		Handler: handler,
		In:      in,
		Out:     out,
	}
}