package controllers

import (
	"github.com/egovorukhin/egomvc/webserver"
)

func Init() []*webserver.Controller {
	return webserver.SetControllers(

		//webserver.NewController(Info{}, ""),
		//webserver.NewSecureController(Info{}, ""),

		webserver.NewController(&Index{}, "/"),
		webserver.NewSecureController(&Index{}, "/"),

		//Вход в систему
		//webserver.NewSecureController(Login{}, ""),
		//Выход из системы
		//webserver.NewSecureController(Logout{}, ""),

		//API
		//webserver.NewController(api.Product{}, ""),
	)
}
