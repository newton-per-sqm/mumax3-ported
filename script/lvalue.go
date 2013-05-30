package script

import (
	"reflect"
)

// read-only value (from script, but mutable from outside)
type reflectROnly struct {
	elem reflect.Value
}

func newReflectROnly(addr interface{}) *reflectROnly {
	return &reflectROnly{reflect.ValueOf(addr).Elem()}
}

func (l *reflectROnly) Eval() interface{} {
	return l.elem.Interface()
}

func (l *reflectROnly) Type() reflect.Type {
	return l.elem.Type()
}

// left-hand value in (single) assign statement
type LValue interface {
	Expr                  // evalutes
	SetValue(interface{}) // assigns a new value
}

// general lvalue implementation using reflect.
// lhs must be settable, e.g. address of something:
// 	var x float64
// 	newReflectLValue(&x)
func newReflectLvalue(addr interface{}) LValue {
	return &reflectLvalue{*newReflectROnly(addr)}
}

type reflectLvalue struct {
	reflectROnly
}

func (l *reflectLvalue) SetValue(rvalue interface{}) {
	l.elem.Set(reflect.ValueOf(rvalue))
}