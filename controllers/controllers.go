package controllers

import (
	"github.com/egovorukhin/egomvc/controllers/api"
	"github.com/egovorukhin/egomvc/webserver"
)

func Init() []webserver.Controller {
	return webserver.SetControllers(

		webserver.NewController(Info{}, ""),
		webserver.NewSecureController(Info{}, ""),

		webserver.NewSecureController(Index{}, "/"),

		//Вход в систему
		webserver.NewSecureController(Login{}, ""),
		//Выход из системы
		webserver.NewSecureController(Logout{}, ""),

		//API
		webserver.NewController(api.Product{}, ""),
	)
}
