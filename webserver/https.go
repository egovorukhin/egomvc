package webserver

import (
	"fmt"
	"github.com/egovorukhin/egologger"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"path/filepath"
	"time"
)

type Https Protocol

func (h *Https) Init(root string) error {

	//Порт
	addr := ":8099"
	if h.Port != "" {
		addr = ":" + h.Port
	}

	//Таймауты
	read, write, idle := h.Timeout.Get()

	//Инициализируем маршрутизатор
	handle := h.InitRoutes(root)

	//Инициализируем сервер
	h.Server = &http.Server{
		Addr:           addr,
		Handler:        handle,
		ReadTimeout:    time.Duration(read) * time.Second,
		WriteTimeout:   time.Duration(write) * time.Second,
		IdleTimeout:    time.Duration(idle) * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	if err := h.Certificate.Check(); err != nil {
		return err
	}

	//В горутине запускаем слушатель
	go h.ListenAsync()

	h.Started = true

	return nil
}

func (h Https) ListenAsync() {
	cert := filepath.Join(h.Certificate.Path, h.Certificate.Cert)
	key := filepath.Join(h.Certificate.Path, h.Certificate.Key)
	if err := h.Server.ListenAndServeTLS(cert, key); err != http.ErrServerClosed {
		egologger.New(h.ListenAsync, logFilename).Error(err)
		fmt.Printf("Https: %s\n", err)
	}
}

func (h Https) InitRoutes(root string) *mux.Router {

	//Инициализируем роутер
	router := mux.NewRouter()

	//Присоединяем к путям директорию static
	static := http.FileServer(http.Dir(path.Join(".", root, "static")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	//Инициализируем Rest API, но обязательно перед этим в пакете controllers
	//добавьте объекты своих структур. Смотрите патерн проектирования EGoMVC
	//GetSecureControllers().SetRouter(router)
	h.Controllers.SetRouter(router)

	return router
}

func (h Https) Redirect(w http.ResponseWriter, r *http.Request, url string, code int) bool {
	if r.TLS == nil {
		redirect(w, r, "https", h.Port, url, code)
		return true
	}
	return false
}
