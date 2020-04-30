package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

const Empty = "У данного маршрута нет реализации"

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

func (m Method) String() string {
	return string(m)
}

//Создаём структуру Route (маршрут)
type Route struct {
	Path        string                                   `json:"path"`
	Method      Method                                   `json:"method"`
	Description string                                   `json:"description"`
	Func        func(http.ResponseWriter, *http.Request) `json:"func,omitempty"`
}

//Добавляем маршрут из структуры Route в router *mux.Router из пакета gorilla
func (r Route) SetRouter(router *mux.Router) {
	/*if r.Redirect {
		r.Func = GetHttps().RedirectToHttps
	}*/
	router.HandleFunc(r.Path, r.Func).Methods(r.Method.String())
}

func SetRoute(path string, method Method, description string, Func func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{
		Path:        path,
		Method:      method,
		Description: description,
		Func:        Func,
	}
}
