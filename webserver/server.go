package webserver

import (
	"fmt"
	"github.com/egovorukhin/egoconf"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var ws WebServer

type WebServer struct {
	Root        *string     `yaml:"root,omitempty"`
	Http        Http        `yaml:"http"`
	Https       Https       `yaml:"https"`
	Certificate Certificate `yaml:"certificate"`
}

func GetWebServer() WebServer {
	return ws
}

func (ws *WebServer) load() error {

	//Загружаем конфигурацию
	err := egoconf.Load("config/webserver.yml", ws)
	if err != nil {
		return err
	}

	return nil
}

func (ws *WebServer) start() string {

	//Проверяем запущен ли сервер
	if (ws.Https.Started) ||
		(ws.Http.Started) {
		return "Web сервер уже запущен"
	}

	//Загружаем конфигурацию
	err := ws.load()
	if err != nil {
		//logger.TraceFileName(ws, ws.start, err, "webserver")
		return err.Error()
	}

	//https
	portHttps := ": не активен"
	if ws.Https.Enabled {
		err = ws.Https.Init()
		if err != nil {
			//logger.TraceFileName(ws, ws.start, err, "webserver")
			return err.Error()
		}
		portHttps = GetHttps().Server.Addr
	}

	//http
	portHttp := ": не активен"
	if ws.Http.Enabled {
		ws.Http.Init()
		portHttp = ws.Http.Server.Addr
	}

	message := fmt.Sprintf("Web сервер запущен [http%s, https%s] время: %s",
		portHttp,
		portHttps,
		getTimeNow(),
	)

	//logger.InfoFileName("ws.start", message, "webserver")

	return message
}

func (ws *WebServer) stop() string {

	ws.Https.Started = false
	ws.Http.Started = false

	message := fmt.Sprintf("Остановка Web сервера [http: %s, https: %s] время: %s",
		Protocol(ws.Http).Shutdown(),
		Protocol(ws.Https).Shutdown(),
		getTimeNow(),
	)

	//logger.InfoFileName("ws.stop", message, "webserver")

	return message
}

func (ws *WebServer) restart() string {

	result := ws.stop() + "\n"
	result += ws.start()

	return result
}

func Init() error {

	ws = WebServer{}

	//Запускаем WebServer
	fmt.Println(ws.start())

	//Крутим в цикле и ждём команды
	for {
		var input string
		_, err := fmt.Fscan(os.Stdin, &input)
		if err != nil {
			return err
		}
		switch strings.ToLower(input) {
		case START:
			fmt.Println(ws.start())
			break
		case STOP:
			fmt.Println(ws.stop())
			break
		case RESTART:
			fmt.Println(ws.restart())
			break
		case EXIT:
			fmt.Println(ws.stop())
			return nil
		case HELP:
			fmt.Println(help())
			break
		case CONFIG:
			fmt.Println(getConfig())
			break
		default:
			fmt.Println("Неизвестная команда! Наберите help для справки.")
			break
		}
	}

}

func redirect(w http.ResponseWriter, r *http.Request, protocol, port, url string, code int) {

	//Если нужно отправить по тому же маршруту, но только по https
	if url == "" {
		host := r.Host[:strings.Index(r.Host, ":")]
		url = path.Join(fmt.Sprintf("%s://%s:%s", protocol, host, port), r.RequestURI)
	}

	//Устанавливаем код перенаправления по умолчанию
	if code == 0 {
		code = http.StatusMovedPermanently
	}

	//Перенаправляем запрос
	http.Redirect(w, r, url, code)
}

func getTimeNow() string {
	t := time.Now()
	return fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d.%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
	)
}
