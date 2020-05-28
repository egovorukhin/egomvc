package webserver

import "net/http"

//BasicAuth - авторизация способом Basic=
//auth bool - флаг для отправки заголовка авторизации в браузере
func BasicAuth(w http.ResponseWriter, r *http.Request, f func(username, password string) error) bool {
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

func FormAuth(r *http.Request, f func(username, password string) error) (string, bool) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	err := f(username, password)
	if err != nil {
		return "", false
	}
	return username, true
}
