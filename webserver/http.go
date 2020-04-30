package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Http Protocol

func GetHttp() Http {
	return ws.Http
}

func (h *Http) Init() error {

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

	return nil
}

func (h Http) ListenAsync()  {
	if err := h.Server.ListenAndServe(); err != http.ErrServerClosed {

	}
}

func (Http) InitRoutes() *mux.Router {

	//Инициализируем роутер
	router := mux.NewRouter()

	//Присоединяем к путям директорию static
	static := http.FileServer(http.Dir("./static"))
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
