package gvar

import "github.com/ilylx/gconv"

// Struct maps value of <v> to <pointer>.
// The parameter <pointer> should be a pointer to a struct instance.
// The parameter <mapping> is used to specify the key-to-attribute mapping rules.
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Struct(v.Val(), pointer, mapping...)
}

// Structs converts and returns <v> as given struct slice.
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Structs(v.Val(), pointer, mapping...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter
// <pointer> to implement the converting.
// It calls function Struct if <pointer> is type of *struct/**struct to do the converting.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct to do the converting.
func (v *Var) Scan(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Scan(v.Val(), pointer, mapping...)
}
