package webserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"path/filepath"
	"time"
)

type Https Protocol

func GetHttps() Https {
	return ws.Https
}

func (h *Https) Init() error {

	//Порт
	addr := ":8099"
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
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	if err := ws.Certificate.Check(); err != nil {
		return err
	}

	//В горутине запускаем слушатель
	go h.ListenAsync()

	h.Started = true

	return nil
}

func (h Https) ListenAsync() {
	cert := filepath.Join(ws.Certificate.Path, ws.Certificate.Cert)
	key := filepath.Join(ws.Certificate.Path, ws.Certificate.Key)
	if err := h.Server.ListenAndServeTLS(cert, key); err != http.ErrServerClosed {
		fmt.Println(err)
		//logger.TraceFileName(h, h.ListenAsync, err, "webserver")
	}
}

func (h Https) InitRoutes() *mux.Router {

	//Инициализируем роутер
	router := mux.NewRouter()
	root := ""
	if ws.Root != nil {
		root = *ws.Root
	}

	//Присоединяем к путям директорию static
	static := http.FileServer(http.Dir(path.Join(".", root, "static")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	//Инициализируем Rest API, но обязательно перед этим в пакете controllers
	//добавьте объекты своих структур. Смотрите патерн проектирования EGoMVC
	GetSecureControllers().SetRouter(router)

	return router
}

func (h Https) Redirect(w http.ResponseWriter, r *http.Request, url string, code int) bool {
	if r.TLS == nil {
		redirect(w, r, "https", h.Port, url, code)
		return true
	}
	return false
}
