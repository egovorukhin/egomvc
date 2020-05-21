package controllers

import (
	"github.com/egovorukhin/egomvc/webserver"
	"github.com/egovorukhin/egomvc/webserver/response"
	"net/http"
)

type Index webserver.Controller

func (a Index) New(path string, secure bool) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return a.Controller().
		SetName(a, path).
		SetSecure(secure).
		SetDescription("Контроллер манипулируем данными о пользователе").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get))
}

func (a Index) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Index) Set(name, description string, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, routes)
}

func (a Index) Get(w http.ResponseWriter, r *http.Request) {

	session, err := webserver.VerifySession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		return
	}

	response.Page(a, w, "", session.Username)
}
