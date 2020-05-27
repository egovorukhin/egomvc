package webserver

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var secureCookie = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

//Структура Controller - типа абстрактный класс,
//её необходио переопределить (типа base или super классы)
type Controller struct {
	Name        string //Имя структуры
	Description string //Описание структуры
	Secure      bool   //https протокол
	Routes      Routes //Маршруты
}

//Массив с маршрутами (Каждый маршрут имеет
//разные методы и функции GET, POST...)
type Routes []*Route

type IController interface {
	New(string) Controller
	Set(string, string, bool, Routes) Controller
}

//Устанавливаем контролеры которые переопределены в маршрутных структурах
//New переопределённая функция в тех структурах типа Info...
func NewController(ic IController, path string) Controller {
	return ic.New(path).setSecure(false)
}

//Устанавливаем контролеры которые переопределены в маршрутных структурах
//New переопределённая функция в тех структурах типа Info...
func NewSecureController(ic IController, path string) Controller {
	return ic.New(path).setSecure(true)
}

//Устанавливаем контролеры которые переопределены в маршрутных структурах
//New переопределённая функция в тех структурах типа Info...
func SetController(ic IController, name, description string, secure bool, routes Routes) Controller {
	return ic.Set(name, description, secure, routes)
}

//Инициализируем Controller
func InitController(name, description string, secure bool, routes Routes) Controller {
	return Controller{
		Name:        name,
		Secure:      secure,
		Description: description,
		Routes:      routes,
	}
}

func (c Controller) setSecure(secure bool) Controller {
	c.Secure = secure
	return c
}

//Создаём Route маршрут для контролера
func (c Controller) NewRoute(route *Route) Controller {
	c.Routes = append(c.Routes, route)
	return c
}

//Можем просто добавить сразу массив из маршрутов
func (c Controller) SetRoutes(routes Routes) Controller {
	c.Routes = routes
	return c
}

//Устанавливаем маршруты для основного роутера
func (c Controller) SetRouter(router *mux.Router) {
	for _, route := range c.Routes {
		route.SetRouter(router)
	}
}

//Формируем имя контролера, не можем использовать c Controller
//потому что reflect.TypeOf(v).String() возвращает имя данного пакета (controller),
//а нужно именно v interface{}
func (c Controller) SetName(v interface{}, path string) Controller {

	//Добавляем маршрут к имени контроллера для того,
	//чтобы можно было создать несколько одинаковых котроллеров с разными путями
	c.Name = reflect.TypeOf(v).String() + path
	return c
}

//Устанавливаем описание контроллера
func (c Controller) SetDescription(s string) Controller {
	c.Description = s
	return c
}

//Извлекаем из адресной строки параметры mux.Vars
func (c Controller) Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

//Проверяем на пустоту путь, если путь пуст то забираем из PkgPath
func CheckPath(path string, v interface{}) string {
	if path == "" {
		path = getPkgPath(v)
	}
	return path
}

//Ищем все после пакета controllers
func getPkgPath(v interface{}) string {
	t := reflect.TypeOf(v)
	pkg := strings.Replace(
		regexp.MustCompile(`controllers(.*)$`).FindString(t.PkgPath()),
		"controllers",
		"",
		-1,
	)
	return strings.Join([]string{pkg, strings.ToLower(t.Name())}, "/")
}

//Статичный map со структурами
//var controllers Controllers

/*
//Возвращаем все контроллеры
func GetControllers() Controllers {
	return controllers
}

//Возвращаем все контроллеры
func GetSecureControllers() Controllers {
	newControllers := Controllers{}
	for _, value := range controllers {
		if value.Secure {
			newControllers[value.Name] = value
		}
	}
	return newControllers
}
*/
//Инициализируем map с Controllers
/*
func SetControllers(values ...Controller) []Controller {
	controllers := Controllers{}
	for _, value := range values {
		controllers[value.Name] = value
	}
	return controllers
}*/

//Заполняем маршруты из map с Controllers
func (controllers Controllers) SetRouter(router *mux.Router) {
	for _, controller := range controllers {
		controller.SetRouter(router)
	}
}

//Redirect - перенаправление на другую ссылку
func (Controller) Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, url, code)
}

/*
//Redirect на https
func (c Controller) SecureRedirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	GetHttps().Redirect(w, r, url, code)
}
*/
