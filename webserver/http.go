package webserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"time"
)

type Http Protocol

func (h *Http) Init(root string) {

	//Порт
	addr := ":8098"
	if h.Port != "" {
		addr = ":" + h.Port
	}

	//Таймауты
	read, write := h.Timeout.Get()

	//Инициализируем маршрутизатор
	handle := h.InitRoutes(root)

	//Инициализируем сервер
	h.Server = &http.Server{
		Addr:           addr,
		Handler:        handle,
		ReadTimeout:    time.Duration(read) * time.Second,
		WriteTimeout:   time.Duration(write) * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	//В горутине запускаем слушатель
	go h.ListenAsync()

	h.Started = true

}

func (h Http) ListenAsync() {
	err := h.Server.ListenAndServe()
	if err != http.ErrServerClosed {
		fmt.Printf("Http: %s\n", err)
	}
}

func (h Http) InitRoutes(root string) *mux.Router {

	//Инициализируем роутер
	router := mux.NewRouter()

	//Присоединяем к путям директорию static
	static := http.FileServer(http.Dir(path.Join(".", root, "static")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	//Добавляем все маршруты, но обязательно перед этим в пакете controllers
	//добавьте объекты своих структур. Смотрите патерн проектирования EGoMVC
	//GetControllers().SetRouter(router)
	h.Controllers.SetRouter(router)

	return router
}

func (h Http) Redirect(w http.ResponseWriter, r *http.Request, url string, code int) bool {
	if r.TLS != nil {
		redirect(w, r, "http", h.Port, url, code)
		return true
	}
	return false
}
