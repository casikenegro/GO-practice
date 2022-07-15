package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, existPath := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, existPath, methodExist
}
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, errPath, errMethod := r.FindHandler(request.URL.Path, request.Method)

	if !errPath || !errMethod {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}
