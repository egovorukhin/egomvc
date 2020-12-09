package controllers

import (
	"fmt"
	"github.com/egovorukhin/egomvc/webserver"
	"net/http"
)

type Index webserver.Controller

func (a *Index) New(path string) *webserver.Controller {
	path = webserver.CheckPath(path, a)
	a.Name = webserver.SetControllerName(a, path)
	a.Description = "Контроллер манипулирует данными о пользователе"
	return a.Controller().NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get))
}

func (a Index) Controller() *webserver.Controller {
	controller := webserver.Controller(a)
	return &controller
}

func (a Index) Set(name, description string, secure bool, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, secure, routes)
}

func (a Index) Get(w http.ResponseWriter, r *http.Request) {

	fmt.Println(a)

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
