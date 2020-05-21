package controllers

import (
	"github.com/egovorukhin/egomvc/controllers/api"
	"github.com/egovorukhin/egomvc/webserver"
)

//Устанавливаем все маршруты из структур
func Init() {
	webserver.SetControllers(
		webserver.NewController(Index{}, "/", false),

		//Вход в систему
		webserver.NewController(Login{}, "", false),
		//Выход из системы
		webserver.NewController(Logout{}, "", false),

		//API
		webserver.NewController(api.Product{}, "", false),
	)
}
