package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	mux *http.ServeMux
	handlers map[string]map[string]http.HandlerFunc
}

func (router *Router) SetHeaderJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}
func (router *Router) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "method not found",
	})
}

func (router *Router) Run(addr string) {
	for path := range router.handlers{
		router.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			router.SetHeaderJson(&w)

			if handler, ok := router.handlers[r.URL.Path][r.Method]; ok {
				handler(w,r)
				return
			}

			re := regexp.MustCompile(`/\d+`) // /users/1
			match := re.FindString(r.URL.Path) // /1
			if match != "" { 
				path := strings.Replace(r.URL.Path, match, ``, -1) + "/" // /users/
				if handler, ok := router.handlers[path][r.Method]; ok {
					handler(w,r)
					return
				}
			}
			router.NotFound(w,r)
		})
	}
	http.ListenAndServe(addr, router.mux)
}

func (route *Router) Get(path string, handler http.HandlerFunc) {
	path = strings.Replace(path, ":id", ``, -1)
	if _, ok := route.handlers[path]; !ok {
		route.handlers[path] = map[string]http.HandlerFunc{}
	}
	route.handlers[path]["GET"] = handler
}

func (route *Router) Post(path string, handler http.HandlerFunc) {
	if _, ok := route.handlers[path]; !ok {
		route.handlers[path] = map[string]http.HandlerFunc{}
	}
	route.handlers[path]["POST"] = handler
}
func (route *Router) Put(path string, handler http.HandlerFunc) {
	path = strings.Replace(path, ":id", ``, -1)
	if _, ok := route.handlers[path]; !ok {
		route.handlers[path] = map[string]http.HandlerFunc{}
	}
	route.handlers[path]["PUT"] = handler
}
func (route *Router) Delete(path string, handler http.HandlerFunc) {
	path = strings.Replace(path, ":id", ``, -1)
	if _, ok := route.handlers[path]; !ok {
		route.handlers[path] = map[string]http.HandlerFunc{}
	}
	route.handlers[path]["DELETE"] = handler
}
