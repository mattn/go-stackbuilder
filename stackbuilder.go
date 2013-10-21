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

type Builder struct {
	mux *http.ServeMux
}

func (b Builder) Build(hs ...interface{}) http.Handler {
	curr := http.Handler(b.mux)
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

func New(mux *http.ServeMux) Builder {
	return Builder{mux}
}

func Build(hs ...interface{}) http.Handler {
	return New(http.DefaultServeMux).Build(hs...)
}
