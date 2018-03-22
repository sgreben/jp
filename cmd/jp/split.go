package main

import "reflect"

func flatten(in [][]reflect.Value) (out []reflect.Value) {
	for _, a := range in {
		for i := range a {
			out = append(out, a[i])
		}
	}
	return
}

func split(in [][]reflect.Value) (x, y [][]reflect.Value) {
	flat := flatten(in)
	n := len(flat)
	x = [][]reflect.Value{flat[:n/2]}
	y = [][]reflect.Value{flat[n/2:]}
	return
}
