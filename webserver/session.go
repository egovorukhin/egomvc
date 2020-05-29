package webserver

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/egovorukhin/egoconf"
	"io"
	"net/http"
	"strings"
	"time"
)

const cookieName = "SID"

var save func(session Session) error
var load func(sid string) (*Session, error)

type Session struct {
	Id         string
	IpAddress  string
	Username   string
	UserAgent  string
	Authorized bool
	DateTime   time.Time
}

type Sessions []*Session

func SetSessionSaveFunc(f func(Session) error) {
	save = f
}

func SetSessionLoadFunc(f func(string) (*Session, error)) {
	load = f
}

func SetSession(w http.ResponseWriter, r *http.Request, username string) error {
	ip, _ := parseIpAddressPort(r.RemoteAddr)
	session := Session{
		Id: generateId(),
		//Извлекаем ip адрес не учитывая порт
		IpAddress:  ip,
		Username:   username,
		UserAgent:  r.Header.Get("User-Agent"),
		Authorized: true,
		DateTime:   time.Now(),
	}

	if save == nil {
		err := session.SaveToFile()
		if err != nil {
			return err
		}
	} else {
		err := save(session)
		if err != nil {
			return err
		}
	}

	//Устанавливаем Cookie
	cookie := http.Cookie{
		Name:    cookieName,
		Value:   session.Id,
		Expires: time.Now().Add(time.Hour * 6),
	}

	http.SetCookie(w, &cookie)

	return nil
}

func (session Session) SaveToFile() error {

	fn := "config/sessions.json"

	sessions := Sessions{}
	err := egoconf.Load(fn, &sessions)
	if err != nil {
		return err
	}

	sessions = append(sessions, &session)

	err = egoconf.Save(fn, sessions)
	if err != nil {
		return err
	}

	return nil
}

func UnAuthorization(w http.ResponseWriter, r *http.Request) error {

	//Находим сессию
	session, err := VerifySession(r)
	if err != nil {
		return err
	}

	//Ставим флаг НеАвторизован
	session.Authorized = false

	//Сохраняем результат
	if save == nil {
		err := session.SaveToFile()
		if err != nil {
			return err
		}
	} else {
		err := save(*session)
		if err != nil {
			return err
		}
	}

	//Очищаем cookie
	cookie := http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)

	return nil
}

//Проверка сессии в хранилище (файл, бд и т.д.)
func VerifySession(r *http.Request) (*Session, error) {

	//Получаем значение sid из cookie
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	//Ищем сессию в хранилище
	session := &Session{}
	if load == nil {
		session, err = loadSessionFromFile(cookie.Value)
		if err != nil {
			return nil, err
		}
	} else {
		session, err = load(cookie.Value)
		if err != nil {
			return nil, err
		}
	}

	if session == nil {
		return nil, errors.New("Сессия не найдена")
	}

	ip, _ := parseIpAddressPort(r.RemoteAddr)
	if session.IpAddress == ip && !session.Authorized {
		return nil, errors.New("Ip адрес указанный в сессии не соответствует удаленному адресу")
	}

	return session, nil
}

func (ss Sessions) Get(sid string) *Session {
	for _, s := range ss {
		if s.Id == sid {
			return s
		}
	}
	return nil
}

func loadSessionFromFile(sid string) (*Session, error) {
	sessions := Sessions{}
	err := egoconf.Load("config/sessions.json", &sessions)
	if err != nil {
		return nil, err
	}

	return sessions.Get(sid), nil
}

func generateId() string {

	//Объявляем срез байт длиной 32 элемента
	b := make([]byte, 32)

	//Рандомно срез заполняем байтами
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}

	//Преобразуем в Base64
	return base64.URLEncoding.EncodeToString(b)
}

func parseIpAddressPort(address string) (string, string) {
	s := strings.Index(address, ":")
	ip := address[:s]
	port := address[s+1:]
	return ip, port
}
