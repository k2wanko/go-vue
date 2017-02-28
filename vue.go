// Package vue exports explicit bindings for the Vue Javascript Library.
// Target version 2.2.1
package vue

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jsbuiltin"
)

var (
	Config *config

	vue *js.Object
)

func init() {
	require := js.Global.Get("require")
	if js.Undefined != require && jsbuiltin.TypeOf(require) == "function" {
		vue = require.Invoke("vue")
	} else {
		vue = js.Global.Get("Vue")
	}

	if vue != nil {
		Config = &config{Object: vue.Get("config")}
	}
}

type (
	// Vue is component instance.
	Vue struct {
		*js.Object
		El          *js.Object `js:"$el"`
		Data        *js.Object `js:"$data"`
		Options     *js.Object `js:"$options"`
		Parent      *js.Object `js:"$parent"`
		Root        *js.Object `js:"$root"`
		Children    *js.Object `js:"$children"`
		Refs        *js.Object `js:"$refs"`
		Slots       *js.Object `js:"$slots"`
		ScopedSlots *js.Object `js:"$scopedSlots"`
		VNode       *js.Object `js:"$vnode"`
		IsServer    bool       `js:"$isServer"`
	}

	config struct {
		*js.Object
		Silent                bool                   `js:"silent"`
		OptionMergeStrategies map[string]*js.Object  `js:"optionMergeStrategies"`
		Devtools              bool                   `js:"devtools"`
		ErrorHandler          *js.Object             `js:"errorHandler"`
		IgnoredElements       []string               `js:"ignoredElements"`
		KeyCodes              map[string]interface{} `js:"keyCodes"`
	}
)

func Directive(id string, def *js.Object) *js.Object {
	return vue.Call("directive", id, def)
}

func Component(id string, def *ComponentOptions) *js.Object {
	return vue.Call("component", id, def)
}

func Filter(id string, def *ComponentOptions) *js.Object {
	return vue.Call("filter", id, def)
}

func New(options *ComponentOptions) *Vue {
	return &Vue{Object: vue.New(options)}
}

func FromObject(o *js.Object) *Vue {
	return &Vue{Object: o}
}

func (v *Vue) Mount(elQuery string) {
	v.Call("$mount", elQuery)
}

func (v *Vue) ForceUpdate() {
	v.Call("$forceUpdate")
}

func (v *Vue) Destroy() {
	v.Call("$destroy")
}

func (v *Vue) Set(o *js.Object, key string, val interface{}) {
	v.Call("$set", o, key, val)
}

func (v *Vue) Delete(o *js.Object, key string) {
	v.Call("$delete", o, key)
}

func (v *Vue) Watch(expOrFn, cb, options *js.Object) func() {
	unwatch := v.Call("$watch", expOrFn, cb, options)
	return func() {
		unwatch.Invoke()
	}
}

func (v *Vue) On(event string, f *js.Object) *Vue {
	v.Call("$on", event, f)
	return v
}

func (v *Vue) Once(event string, f *js.Object) *Vue {
	v.Call("$once", event, f)
	return v
}

func (v *Vue) Off(event string, f *js.Object) *Vue {
	v.Call("$off", event, f)
	return v
}

func (v *Vue) Emit(event string, args ...interface{}) *Vue {
	v.Call("$emit", append([]interface{}{event}, args...)...)
	return v
}

func (v *Vue) NextTick(f *js.Object) {
	v.Call("$nextTick", f)
}

func (v *Vue) CreateElement(tag, data, children *js.Object) *js.Object {
	return v.Call("$createElement", tag, data, children)
}

func newObject() *js.Object {
	return js.Global.Get("Object").New()
}
