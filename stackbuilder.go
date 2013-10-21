package stackbuilder

import (
	"net/http"
	"reflect"
)

type handler struct {
	h http.Handler
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.h.ServeHTTP(w, r)
}

func Build(hs ...interface{}) http.Handler {
	curr := http.Handler(http.DefaultServeMux)
	for _, h := range hs {
		if _, ok := h.(http.Handler); ok {
			curr = &handler{curr}
		} else {
			rv := reflect.ValueOf(h)
			if rv.Kind() == reflect.Func {
				curr = rv.Call([]reflect.Value{reflect.ValueOf(curr)})[0].Interface().(http.Handler)
			}
		}
	}
	return curr
}
