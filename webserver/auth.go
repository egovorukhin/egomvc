package webserver

import "net/http"

//BasicAuth - авторизация способом Basic=
//auth bool - флаг для отправки заголовка авторизации в браузере
func (Controller) BasicAuth(w http.ResponseWriter, r *http.Request, f func(username, password string) error) bool {
	username, password, ok := r.BasicAuth()
	if ok {
		err := f(username, password)
		if err == nil {
			return true
		}
	}
	w.Header().Add("WWW-Authenticate", `Basic realm="EgoMvc"`)
	w.WriteHeader(http.StatusUnauthorized)

	return false
}

func (Controller) FormAuth(r *http.Request, f func(username, password string) error) (string, error) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	err := f(username, password)
	if err != nil {
		return "", err
	}
	return username, nil
}
