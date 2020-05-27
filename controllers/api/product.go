package api

import (
	"github.com/egovorukhin/egomvc/src/database"
	"github.com/egovorukhin/egomvc/webserver"
	"net/http"
)

type Product webserver.Controller

func (a Product) New(path string) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return a.Controller().
		SetName(a, path).
		SetDescription("Контроллер манипулируем данными о пользователе").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get))
}

func (a Product) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Product) Set(name, description string, secure bool, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, secure, routes)
}

func (a Product) Get(w http.ResponseWriter, r *http.Request) {

	if !a.Controller().BasicAuth(w, r, database.Authorization) {
		return
	}

	webserver.Ok(w, "Здесь куча продуктов")
}
