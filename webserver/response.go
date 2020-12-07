package webserver

import (
	"encoding/json"
	"encoding/xml"
	"github.com/egovorukhin/egologger"
	"html/template"
	"net/http"
	"path"
)

const logControllers = "controllers"

type Response struct {
	writer http.ResponseWriter
	code   int
}

const (
	Text = "text/plain"
	Json = "application/json"
	Xml  = "application/xml"
)

//Отправляем ответ в формате Json
func (r Response) Json(i interface{}) {
	r.writer.Header().Add("Content-Type", Json)
	r.writer.WriteHeader(r.code)
	err := json.NewEncoder(r.writer).Encode(i)
	if err != nil {
		_, err = r.writer.Write([]byte(err.Error()))
		if err != nil {
			egologger.Error(Page, logControllers, err.Error())
			return
		}
		egologger.Error(Page, logControllers, err.Error())
	}
}

//Отправляем ответ в формате Xml
func (r Response) Xml(i interface{}) {
	r.writer.Header().Add("Content-Type", Xml)
	r.writer.WriteHeader(r.code)
	err := xml.NewEncoder(r.writer).Encode(i)
	if err != nil {
		_, err = r.writer.Write([]byte(err.Error()))
		if err != nil {
			egologger.Error(Page, logControllers, err.Error())
			return
		}
		egologger.Error(Page, logControllers, err.Error())
	}
}

//Отправляем ответ в формате Xml
func (r Response) Text(s string) {
	r.writer.Header().Add("Content-Type", Text)
	r.writer.WriteHeader(r.code)
	_, err := r.writer.Write([]byte(s))
	if err != nil {
		egologger.Error(Page, logControllers, err.Error())
	}
}

func OK(w http.ResponseWriter) Response {
	return Response{
		writer: w,
		code:   http.StatusOK,
	}
}

func Error(w http.ResponseWriter, code int) Response {
	return Response{
		writer: w,
		code:   code,
	}
}

//Возвращаем html страницу
//Используется для страниц Views, рендеринг страниц
func View(i interface{}, w http.ResponseWriter, pageName string, data interface{}) {
	if pageName == "" {
		pageName = CheckPath("", i) + ".html"
	}

	tmpl, err := template.ParseFiles(
		path.Join("views/share", "layout.html"),
		path.Join("views", pageName))
	if err != nil {
		View(i, w, "share/error.html", err.Error())
		egologger.Error(Page, logControllers, err.Error())
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		egologger.Error(Page, logControllers, err.Error())
	}
}

//Отдаём страницы которые находяться в папке www.
//Т.е. используя функцию Page мы из модели MVC убираем views, а так же static
//в том виде который используется для MVC. Такой подход был реализован для проектов на React JS.
//Собираем проект React App с помощью npm или yarn, копируем все содержимое
//каталога build в каталог www вашего проекта и все будет работать. не забудьте создать
//контроллер Index, и добавить все возможные пути (react-router-dom) это очень важно.
func Page(i interface{}, w http.ResponseWriter, pageName string, data interface{}) {
	if pageName == "" {
		pageName = CheckPath("", i)
	}
	pageName += ".html"

	page, err := template.ParseFiles(path.Join(getRoot(), pageName))
	if err != nil {
		egologger.Error(Page, logControllers, err.Error())
	}
	err = page.Execute(w, data)
	if err != nil {
		egologger.Error(Page, logControllers, err.Error())
	}
}
