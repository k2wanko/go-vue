package render

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
)

type (
	// Renderer rendering from component.
	Renderer struct {
		// Program is javascript code.
		// Example:
		//   TODO: example code.
		//   module.exports = function (context) {
		//      var res = context.res
		//   }
		//
		Path string

		// TODO:
		// Component *vue.Vue

		Data  map[string]interface{}
		Cache ComponentCacher

		b *bytes.Buffer
	}
)

// Read implements io#Reader
func (r *Renderer) Read(p []byte) (n int, err error) {
	if r.b == nil {
		r.b, err = r.render()
		if err != nil {
			return
		}
	}

	return r.b.Read(p)
}

// Reset reset instance.
func (r *Renderer) Reset() {
	if r.b != nil {
		r.b.Reset()
	}
}

// Render returns html bytes.
func (r *Renderer) Render() (b []byte, err error) {
	var buf *bytes.Buffer
	buf, err = r.render()
	if err != nil {
		return
	}
	b = buf.Bytes()
	buf.Reset()
	return
}

func (r *Renderer) render() (buf *bytes.Buffer, err error) {
	if r.Path == "" {
		return nil, errors.New("Path is empty")
	}

	var b bytes.Buffer
	loop := eventloop.NewEventLoop()
	loop.Run(func(vm *goja.Runtime) {
		mr := new(require.Registry)
		rm := mr.Enable(vm)
		err = applyProcess(vm)
		if err != nil {
			return
		}

		if r.Cache != nil {
			applyComponentCache(r.Cache, vm)
		}

		var v goja.Value
		v, err = rm.Require(r.Path)

		if jserr, ok := err.(*goja.Exception); ok {
			err = errors.New(jserr.Value().String())
			return
		} else if err != nil {
			return
		}

		renderFunc, ok := goja.AssertFunction(v.ToObject(vm))
		if !ok {
			err = errors.New("not function render")
			return
		}

		jsCtx := vm.NewObject()

		resObj := vm.NewObject()
		resObj.Set("write", func(c goja.FunctionCall) goja.Value {
			data := c.Argument(0).String()
			b.Write([]byte(data))
			return goja.Undefined()
		})
		resObj.Set("end", func(c goja.FunctionCall) goja.Value {
			data := c.Argument(0).String()
			b.Write([]byte(data))
			return goja.Undefined()
		})
		resObj.Set("error", func(c goja.FunctionCall) goja.Value {
			if len(c.Arguments) == 0 {
				return goja.Undefined()
			}
			err = fmt.Errorf("renderFunc: %#v", c.Argument(0).String())
			return goja.Undefined()
		})

		var ctxData map[string]interface{}
		if r.Data != nil {
			ctxData = r.Data
			if v, ok := r.Data["res"]; ok {
				err = fmt.Errorf("'res' is the reserved name. v = %v", v)
				return
			}
		} else {
			ctxData = make(map[string]interface{})
		}

		ctxData["res"] = resObj

		for k, v := range ctxData {
			jsCtx.Set(k, v)
		}

		_, err = renderFunc(nil, jsCtx)
		if jserr, ok := err.(*goja.Exception); ok {
			err = errors.New(jserr.Value().String())
			return
		} else if err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	buf = &b

	return
}

func applyProcess(vm *goja.Runtime) (err error) {
	process := vm.NewObject()
	env := vm.NewObject()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		env.Set(pair[0], pair[1])
	}
	process.Set("browser", true)
	process.Set("env", env)
	setTimeout, ok := goja.AssertFunction(vm.Get("setTimeout"))
	if !ok {
		err = errors.New("setTimeout is not fcuntion")
		return
	}
	process.Set("nextTick", func(c goja.FunctionCall) goja.Value {
		if len(c.Arguments) == 0 {
			return vm.NewGoError(errors.New("arguments is 0"))
		}
		setTimeout(nil, c.Argument(0))
		return goja.Undefined()
	})
	vm.Set("process", process)
	return
}
