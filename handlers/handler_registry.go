package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RequestHandlerRegistry []*RequestHandler

type RequestHandler struct {
	Path    string
	Methods []string
	Handler func(w http.ResponseWriter, r *http.Request)
}

type DocumentRequestHandler struct {
	Path          string
	IndexHandler  func(w http.ResponseWriter, r *http.Request)
	CreateHandler func(w http.ResponseWriter, r *http.Request)
	ReadHandler   func(w http.ResponseWriter, r *http.Request)
	UpdateHandler func(w http.ResponseWriter, r *http.Request)
	DeleteHandler func(w http.ResponseWriter, r *http.Request)
}

func (drh DocumentRequestHandler) RegisterHandlers() {
	if drh.IndexHandler != nil {
		GetRegistry().Add(&RequestHandler{Path: drh.Path, Methods: []string{"GET"}, Handler: drh.IndexHandler})
	}
	if drh.CreateHandler != nil {
		GetRegistry().Add(&RequestHandler{Path: drh.Path, Methods: []string{"POST"}, Handler: drh.CreateHandler})
	}
	if drh.ReadHandler != nil {
		GetRegistry().Add(&RequestHandler{Path: drh.Path + "/{id}", Methods: []string{"GET"}, Handler: drh.ReadHandler})
	}
	if drh.UpdateHandler != nil {
		GetRegistry().Add(&RequestHandler{Path: drh.Path + "/{id}", Methods: []string{"PUT"}, Handler: drh.UpdateHandler})
	}
	if drh.DeleteHandler != nil {
		GetRegistry().Add(&RequestHandler{Path: drh.Path + "/{id}", Methods: []string{"DELETE"}, Handler: drh.DeleteHandler})
	}
}

var registry = make(RequestHandlerRegistry, 0)

func GetRegistry() *RequestHandlerRegistry {
	if registry == nil {
		registry = make([]*RequestHandler, 0)
	}
	return &registry
}

func (r *RequestHandlerRegistry) Add(handlers ...*RequestHandler) {
	*r = append(*r, handlers...)
}

func (r *RequestHandlerRegistry) Get() []*RequestHandler {
	return *r
}

func (r *RequestHandlerRegistry) Clear() {
	*r = make([]*RequestHandler, 0)
}

func (r *RequestHandlerRegistry) AddHandler(path string, methods []string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Add(&RequestHandler{Path: path, Methods: methods, Handler: handler})
}

func (r *RequestHandlerRegistry) RegisterHandlers(router *mux.Router) {
	for _, handler := range r.Get() {
		router.HandleFunc(handler.Path, handler.Handler).Methods(handler.Methods...)
	}
}
