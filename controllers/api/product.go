package api

/*
type Product webserver.Controller

func (a Product) New(path string) webserver.Controller {
	path = webserver.CheckPath(path, a)
	return a.Controller().
		SetName(a, path).
		SetDescription("Контроллер манипулируем данными о пользователе").
		NewRoute(webserver.SetRoute(path, webserver.GET, "Доступные пользователи", a.Get))
}

func (a Product) Controller() webserver.Controller {
	return webserver.Controller(a)
}

func (a Product) Set(name, description string, secure bool, routes webserver.Routes) webserver.Controller {
	return webserver.InitController(name, description, secure, routes)
}

func (a Product) Get(w http.ResponseWriter, r *http.Request) {

	err := webserver.BasicAuth(w, r, database.Authorization)
	if err != nil {
		webserver.OK(w).Json(err.Error())
		return
	}

	webserver.OK(w).Json("Куча продуктов")
}*/
