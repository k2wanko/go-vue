package render

import "github.com/dop251/goja"

type (
	ComponentCacher interface {
		Get(string) ([]byte, error)
		Set(string, []byte) error
	}
)

func applyComponentCache(cc ComponentCacher, vm *goja.Runtime) {
	o := vm.NewObject()
	o.Set("get", func(fc goja.FunctionCall) (v goja.Value) {
		if len(fc.Arguments) == 0 {
			return goja.Undefined()
		}

		key := fc.Argument(0).String()

		b, err := cc.Get(key)
		if err != nil {
			return vm.NewGoError(err)
		}

		if b == nil {
			return goja.Undefined()
		}

		return vm.ToValue(string(b))
	})
	o.Set("set", func(fc goja.FunctionCall) (v goja.Value) {
		if len(fc.Arguments) < 2 {
			return goja.Undefined()
		}

		k := fc.Argument(0).String()
		val := fc.Argument(1).String()
		err := cc.Set(k, []byte(val))
		if err != nil {
			return vm.NewGoError(err)
		}

		return goja.Undefined()
	})
	vm.Set("__ComponentCache__", o)
}
