package render

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestRenderer(t *testing.T) {
	r := &Renderer{Path: "./testdata/index.js"}

	res, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("res = \n%s", res)

	if want := "hello world"; 0 > strings.Index(string(res), want) {
		t.Errorf("Not has '%s'\nres = %s", want, res)
	}
}
