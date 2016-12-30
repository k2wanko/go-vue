package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	vue "github.com/k2wanko/go-vue"
)

type Data struct {
	*js.Object
	Message string `js:"message"`
	Count   int    `js:"count"`
}

func main() {
	data := &Data{Object: js.Global.Get("Object").New()}
	data.Message = "Hello, Vue"
	data.Count = 0

	opts := vue.NewComponentOptions()
	opts.Data = data.Object

	vm := vue.New(opts)

	js.Global.Set("vm", vm)

	vm.Mount("#app")

	go func() {
		for {
			time.Sleep(1 * time.Second)
			data.Count = data.Count + 1
		}
	}()
}
