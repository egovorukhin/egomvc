package webserver

import (
	"github.com/egovorukhin/egologger/logger"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"time"
)

type Http Protocol

func GetHttp() Http {
	return ws.Http
}

func (h *Http) Init() {

	//Порт
	addr := ":8098"
	if h.Port != "" {
		addr = ":" + h.Port
	}

	//Таймауты
	read, write := h.Timeout.Get()

	//Инициализируем маршрутизатор
	handle := h.InitRoutes()

	//Инициализируем сервер
	h.Server = &http.Server{
		Addr:           addr,
		Handler:        handle,
		ReadTimeout:    time.Duration(read) * time.Second,
		WriteTimeout:   time.Duration(write) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//В горутине запускаем слушатель
	go h.ListenAsync()

	h.Started = true

}

func (h Http) ListenAsync() {
	err := h.Server.ListenAndServe()
	if err != http.ErrServerClosed {
		logger.TraceFileName(h, h.ListenAsync, err, "webserver")
	}
}

func (Http) InitRoutes() *mux.Router {

	//Инициализируем роутер
	router := mux.NewRouter()
	root := ""
	if ws.Root != nil {
		root = *ws.Root
	}

	//Присоединяем к путям директорию static
	static := http.FileServer(http.Dir(path.Join(".", root, "static")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	//Добавляем все маршруты, но обязательно перед этим в пакете controllers
	//добавьте объекты своих структур. Смотрите патерн проектирования EGoMVC
	GetControllers().SetRouter(router)

	return router
}

func (h Http) Redirect(w http.ResponseWriter, r *http.Request, url string, code int) bool {
	if r.TLS != nil {
		redirect(w, r, "http", h.Port, url, code)
		return true
	}
	return false
}
