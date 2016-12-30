package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

type (
	Array struct {
		*js.Object
	}

	CompareFunc func(*js.Object, *js.Object) int
)

func WrapArray(any *js.Object) *Array {
	return &Array{Object: any}
}

func (a *Array) Push(any ...interface{}) int {
	return a.Object.Call("push", any...).Int()
}

func (a *Array) Pop() *js.Object {
	return a.Object.Call("pop")
}

func (a *Array) Unshift(any ...interface{}) int {
	return a.Object.Call("unshift", any...).Int()
}

func (a *Array) Shift() *js.Object {
	return a.Object.Call("shift")
}

func (a *Array) Splice(i, howMany int, any ...interface{}) *js.Object {
	return a.Object.Call("splice", append([]interface{}{i, howMany}, any...)...)
}

func (a *Array) Sort(f CompareFunc) {
	a.Object.Call("sort", f)
}
