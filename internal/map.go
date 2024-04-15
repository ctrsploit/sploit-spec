package internal

import "reflect"

func ReverseMap(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("m must be a map")
	}
	t := reflect.MapOf(v.Type().Elem(), v.Type().Key())
	r := reflect.MakeMap(t)
	for _, k := range v.MapKeys() {
		r.SetMapIndex(v.MapIndex(k), k)
	}
	return r.Interface()
}
