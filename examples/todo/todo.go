package main

import (
	"github.com/gopherjs/gopherjs/js"
	vue "github.com/k2wanko/go-vue"
)

type Item struct {
	Title string
}

type Data struct {
	*js.Object
	NewTodo string  `js:"newTodo"`
	Items   []*Item `js:"items"`
}

var data = &Data{Object: js.Global.Get("Object").New()}

func addNewTodo(this *js.Object, args []*js.Object) interface{} {
	if data.NewTodo == "" {
		return nil
	}

	vue.WrapArray(data.Get("items")).Push(&Item{
		Title: data.NewTodo,
	})
	data.NewTodo = ""
	return nil
}

func removeTodo(this *js.Object, args []*js.Object) interface{} {
	if len(args) == 0 {
		return nil
	}
	return vue.WrapArray(data.Get("items")).Splice(args[0].Int(), 1)
}

func mounted(this *js.Object, args []*js.Object) interface{} {
	vue.WrapArray(data.Get("items")).Sort(func(x, y *js.Object) int {
		return int(y.Get("Title").String()[0]) - int(x.Get("Title").String()[0])
	})
	return nil
}

func main() {
	data.NewTodo = ""
	data.Items = []*Item{
		{Title: "Angel Beats!"},
		{Title: "Baka to Test to Shoukanjuu"},
		{Title: "Cowboy Bebop"},
		{Title: "Dragon Ball"},
		{Title: "Evangelion"},
		{Title: "Ghost in the Shell"},
		{Title: "Hunter X Hunter"},
	}
	o := vue.NewComponentOptions()
	o.El = "#app"
	o.Data = data.Object
	o.Methods = js.M{
		"addNewTodo": js.MakeFunc(addNewTodo),
		"removeTodo": js.MakeFunc(removeTodo),
	}
	o.Mounted = js.MakeFunc(mounted)
	vm := vue.New(o)
	js.Global.Set("vm", vm)
}
