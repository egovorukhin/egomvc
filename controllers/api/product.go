package api

import (
	"github.com/egovorukhin/egomvc/src/database"
	"github.com/egovorukhin/egomvc/webserver"
	"github.com/egovorukhin/egomvc/webserver/response"
	"net/http"
)

type Product webserver.Controller

func (a Product) New(path string, secure bool) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return a.Controller().
		SetName(a, path).
		SetSecure(secure).
		SetDescription("Контроллер манипулируем данными о пользователе").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get))
}

func (a Product) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Product) Set(name, description string, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, routes)
}

func (a Product) Get(w http.ResponseWriter, r *http.Request) {

	err := a.Controller().BasicAuth(w, r, database.Authorization)
	if err != nil {
		response.Error(w, err.Error())
		return
	}

	response.Ok(w, "Здесь куча продуктов")
}
