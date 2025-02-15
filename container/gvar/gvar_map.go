package gvar

import "github.com/ilylx/gconv"

// Map converts and returns <v> as map[string]interface{}.
func (v *Var) Map(tags ...string) map[string]interface{} {
	return gconv.Map(v.Val(), tags...)
}

// MapStrAny is like function Map, but implements the interface of MapStrAny.
func (v *Var) MapStrAny() map[string]interface{} {
	return v.Map()
}

// MapStrStr converts and returns <v> as map[string]string.
func (v *Var) MapStrStr(tags ...string) map[string]string {
	return gconv.MapStrStr(v.Val(), tags...)
}

// MapStrVar converts and returns <v> as map[string]Var.
func (v *Var) MapStrVar(tags ...string) map[string]*Var {
	m := v.Map(tags...)
	if len(m) > 0 {
		vMap := make(map[string]*Var, len(m))
		for k, v := range m {
			vMap[k] = New(v)
		}
		return vMap
	}
	return nil
}

// MapDeep converts and returns <v> as map[string]interface{} recursively.
func (v *Var) MapDeep(tags ...string) map[string]interface{} {
	return gconv.MapDeep(v.Val(), tags...)
}

// MapDeep converts and returns <v> as map[string]string recursively.
func (v *Var) MapStrStrDeep(tags ...string) map[string]string {
	return gconv.MapStrStrDeep(v.Val(), tags...)
}

// MapStrVarDeep converts and returns <v> as map[string]*Var recursively.
func (v *Var) MapStrVarDeep(tags ...string) map[string]*Var {
	m := v.MapDeep(tags...)
	if len(m) > 0 {
		vMap := make(map[string]*Var, len(m))
		for k, v := range m {
			vMap[k] = New(v)
		}
		return vMap
	}
	return nil
}

// Maps converts and returns <v> as map[string]string.
// See gconv.Maps.
func (v *Var) Maps(tags ...string) []map[string]interface{} {
	return gconv.Maps(v.Val(), tags...)
}

// MapToMap converts any map type variable <params> to another map type variable <pointer>.
// See gconv.MapToMap.
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMap(v.Val(), pointer, mapping...)
}

// MapToMaps converts any map type variable <params> to another map type variable <pointer>.
// See gconv.MapToMaps.
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}
