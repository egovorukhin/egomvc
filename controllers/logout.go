package controllers

import (
	"github.com/egovorukhin/egomvc/webserver"
	"github.com/egovorukhin/egomvc/webserver/response"
	"net/http"
)

type Logout webserver.Controller

func (a Logout) New(path string, secure bool) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return a.Controller().
		SetName(a, path).
		SetSecure(secure).
		SetDescription("Контроллер для выхода из системы").
		NewRoute(webserver.SetRoute(path, webserver.POST, "Выход", a.Post))
}

func (a Logout) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Logout) Set(name, description string, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, routes)
}

func (a Logout) Post(w http.ResponseWriter, r *http.Request) {

	err := webserver.UnAuthorization(w, r)
	if err != nil {
		response.Page(a, w, "/error", err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}
