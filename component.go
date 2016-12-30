package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

type (
	ComponentOptions struct {
		*js.Object

		// data
		Data      *js.Object              `js:"data"`
		Props     map[string]*PropOptions `js:"props"`
		PropsData map[string]interface{}  `js:"propsData"`
		Computed  map[string]interface{}  `js:"computed"`
		Methods   map[string]interface{}  `js:"methods"`
		Watch     map[string]interface{}  `js:"watch"`

		// DOM
		El              string     `js:"el"`
		Template        string     `js:"template"`
		Render          *js.Object `js:"render"`
		StaticRenderFns *js.Object `js:"staticRenderFns"`

		// lifecycle
		BeforeCreate *js.Object `js:"beforeCreate"`
		Created      *js.Object `js:"created"`
		BeforeMount  *js.Object `js:"beforeMount"`
		Mounted      *js.Object `js:"mounted"`
		BeforeUpdate *js.Object `js:"beforeUpdate"`
		Updated      *js.Object `js:"updated"`

		// assets
		Directives  map[string]*js.Object        `js:"directives"`
		Components  map[string]*ComponentOptions `js:"components"`
		Transitions map[string]*js.Object        `js:"transitions"`
		Filters     map[string]*js.Object        `js:"filters"`

		// misc
		Parent     *Vue              `js:"parent"`
		Mixins     []*js.Object      `js:"mixins"`
		Name       string            `js:"name"`
		Extends    *ComponentOptions `js:"extends"`
		Delimiters []string          `js:"delimiters"`
	}

	PropOptions struct {
		*js.Object
		Type      *js.Object            `js:"type"`
		Default   interface{}           `js:"default"`
		Required  bool                  `js:"required"`
		Validator func(*js.Object) bool `js:"validator"`
	}
)

func NewComponentOptions() *ComponentOptions {
	return &ComponentOptions{Object: newObject()}
}

func NewPropsOptions() *PropOptions {
	return &PropOptions{Object: newObject()}
}
