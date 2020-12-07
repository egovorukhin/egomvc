package controllers

import (
	"github.com/egovorukhin/egomvc/webserver"
	"net/http"
)

type Index webserver.Controller

func (a Index) New(path string) *webserver.Controller {
	path = webserver.CheckPath(path, a)
	controller := a.Controller()
	return controller.
		SetName(a, path).
		SetDescription("Контроллер манипулируем данными о пользователе").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get))
}

func (a Index) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Index) Set(name, description string, secure bool, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, secure, routes)
}

func (a Index) Get(w http.ResponseWriter, r *http.Request) {

	session, err := webserver.VerifySessionRedirect(w, r, "/login", http.StatusMovedPermanently)
	if err != nil {
		return
	}
	/*
		session, err := webserver.VerifySession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}*/

	webserver.Page(a, w, "", session.Username)
}
