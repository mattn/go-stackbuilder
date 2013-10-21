package stackbuilder

import (
	"net/http"
	"reflect"
)

func NewStackBuilder(hfs ...interface{}) http.Handler {
	curr := reflect.ValueOf(http.Handler(http.DefaultServeMux))
	for _, hf := range hfs {
		rv := reflect.ValueOf(hf)
		curr = rv.Call([]reflect.Value{curr})[0]
	}
	return curr.Interface().(http.Handler)
}
