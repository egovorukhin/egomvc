package controllers

import (
	"github.com/egovorukhin/egomvc/webserver"
	"net/http"
)

type Info webserver.Controller

func (a Info) New(path string) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return webserver.Controller(a).
		SetName(a, path).
		SetDescription("Информация о контроллерах").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Возвращаем все контроллеры", a.Get))
}

func (a Info) Set(name, description string, secure bool, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, secure, routes)
}

func (a Info) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Info) Get(writer http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		w := webserver.GetSecureControllers()
		webserver.Page(a, writer, "", w)
		return
	}
	webserver.Page(a, writer, "", webserver.GetControllers())
}
