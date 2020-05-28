package controllers

import (
	"github.com/egovorukhin/egomvc/src/database"
	"github.com/egovorukhin/egomvc/webserver"
	"net/http"
)

type Login webserver.Controller

func (a Login) New(path string) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return a.Controller().
		SetName(a, path).
		SetDescription("Контроллер манипулируем данными о пользователе").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get)).
		NewRoute(webserver.SetRoute(path, webserver.POST, "Доступные пользователи", a.Post))
}

func (a Login) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Login) Set(name, description string, secure bool, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, secure, routes)
}

func (a Login) Get(w http.ResponseWriter, r *http.Request) {
	webserver.Page(a, w, "", nil)
}

func (a Login) Post(w http.ResponseWriter, r *http.Request) {

	username, auth := webserver.FormAuth(r, database.Authorization)
	if !auth {
		webserver.Error(w, "Не верный имя пользователя или пароль")
		return
	}

	//Session
	err := webserver.SetSession(w, r, username)
	if err != nil {
		webserver.Error(w, err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
