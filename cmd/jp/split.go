package main

import "reflect"

var indexableKind = map[reflect.Kind]bool{
	reflect.Slice: true,
	reflect.Array: true,
}

func flatten(in [][]reflect.Value) (out []reflect.Value) {
	for _, a := range in {
		for _, v := range a {
			if !v.IsValid() {
				continue
			}
			if indexableKind[v.Kind()] {
				sub := make([]reflect.Value, v.Len())
				for j := 0; j < v.Len(); j++ {
					sub = append(sub, v.Index(j))
				}
				out = append(out, flatten([][]reflect.Value{sub})...)
				continue
			}
			if v.CanInterface() {
				if sub, ok := v.Interface().([]interface{}); ok {
					for i := range sub {
						out = append(out, reflect.ValueOf(sub[i]))
					}
					continue
				}
			}
			out = append(out, v)
		}
	}
	return
}

func split(in [][]reflect.Value) (x, y []reflect.Value) {
	flat := flatten(in)
	n := len(flat)
	x = make([]reflect.Value, 0, n/2)
	y = make([]reflect.Value, 0, n/2)
	for i := range flat {
		if i&1 == 0 {
			x = append(x, flat[i])
		} else {
			y = append(y, flat[i])
		}
	}
	return
}
