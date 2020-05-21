package controllers

import (
	"github.com/egovorukhin/egomvc/src/database"
	"github.com/egovorukhin/egomvc/webserver"
	"github.com/egovorukhin/egomvc/webserver/response"
	"net/http"
)

type Login webserver.Controller

func (a Login) New(path string, secure bool) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return a.Controller().
		SetName(a, path).
		SetSecure(secure).
		SetDescription("Контроллер манипулируем данными о пользователе").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get)).
		NewRoute(webserver.SetRoute(path, webserver.POST, "Доступные пользователи", a.Post))
}

func (a Login) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Login) Set(name, description string, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, routes)
}

func (a Login) Get(w http.ResponseWriter, r *http.Request) {
	response.Page(a, w, "", nil)
}

func (a Login) Post(w http.ResponseWriter, r *http.Request) {

	username, err := a.Controller().FormAuth(r, database.Authorization)
	if err != nil {
		response.Error(w, err.Error())
		return
	}

	//Session
	err = webserver.SetSession(w, r, username)
	if err != nil {
		response.Error(w, err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
