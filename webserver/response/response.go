package response

import (
	"encoding/json"
	"encoding/xml"
	"github.com/egovorukhin/egomvc/webserver"
	"html/template"
	"net/http"
	"path"
)

//Формат передавамых данных
type FormatBody int

const (
	JSON FormatBody = iota
	XML
)

//Глобальная переменная формата передаваемых данных
var formatBody FormatBody

type Code int

const (
	OK Code = iota
	ERROR
)

//Структура ответа
type Response struct {
	Writer http.ResponseWriter
	Body   Body
}

type Body struct {
	Code    Code        `json:"code" xml:"code"`
	Message interface{} `json:"message" xml:"message"`
}

func SetFormatBody(f FormatBody) {
	formatBody = f
}

//Инициализируем ответ для Json и Xml
func Init(w http.ResponseWriter, code Code, message interface{}) Response {
	return Response{
		Writer: w,
		Body: Body{
			Code:    code,
			Message: message,
		},
	}
}

//Возвращаем html страницу
//Используется для страниц Views, рендеринг страниц
func View(i interface{}, w http.ResponseWriter, pageName string, data interface{}) error {
	if pageName == "" {
		pageName = webserver.CheckPath("", i) + ".html"
	}

	tmpl, err := template.ParseFiles(
		path.Join("views/share", "layout.html"),
		path.Join("views", pageName))
	if err != nil {
		err = View(i, w, "share/error.html", err.Error())
		if err != nil {
			return err
		}
		return err
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		return err
	}

	return nil
}

//Отдаём страницы которые находяться в папке www.
//Т.е. используя функцию Page мы из модели MVC убираем views, а так же static
//в том виде который используется для MVC. Такой подход был реализован для проектов на React JS.
//Собираем проект React App с помощью npm или yarn, копируем все содержимое
//каталога build в каталог www вашего проекта и все будет работать. не забудьте создать
//контроллер Index, и добавить все возможные пути (react-router-dom) это очень важно.
func Page(i interface{}, w http.ResponseWriter, pageName string) error {
	if pageName == "" {
		pageName = webserver.CheckPath("", i)
	}
	pageName += ".html"
	page, err := template.ParseFiles(path.Join("www", pageName))
	if err != nil {
		return err
	}
	err = page.Execute(w, nil)
	if err != nil {
		return err
	}
	return nil
}

//Зполняем code = 0 и отправляем в json или xml
//в соответсвии с formatBody
func Ok(w http.ResponseWriter, message interface{}) error {
	if message == nil {
		message = "OK"
	}
	r := Init(w, OK, message)
	if formatBody == JSON {
		return r.Json()
	}

	return r.Xml()
}

//Зполняем code = 0 и отправляем в json или xml
//в соответсвии с formatBody
func Error(w http.ResponseWriter, message interface{}) error {
	if message == nil {
		message = "Error"
	}
	r := Init(w, ERROR, message)
	if formatBody == JSON {
		return r.Json()
	}

	return r.Xml()
}

//Отправляем ответ в формате Json
func (r Response) Json() error {
	r.Writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(r.Writer).Encode(r.Body)
	if err != nil {
		_, err = r.Writer.Write([]byte(err.Error()))
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

//Отправляем ответ в формате Xml
func (r Response) Xml() error {
	r.Writer.Header().Add("Content-Type", "application/xml")
	err := xml.NewEncoder(r.Writer).Encode(r.Body)
	if err != nil {
		_, err = r.Writer.Write([]byte(err.Error()))
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
