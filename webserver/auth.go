package webserver

import "net/http"

//BasicAuth - авторизация способом Basic=
//auth bool - флаг для отправки заголовка авторизации в браузере
func BasicAuth(w http.ResponseWriter, r *http.Request, f func(username, password string) error) error {
	username, password, ok := r.BasicAuth()
	var err error
	if ok {
		err = f(username, password)
		if err == nil {
			return nil
		}
	}
	w.Header().Add("WWW-Authenticate", `Basic realm="EgoMvc"`)
	w.WriteHeader(http.StatusUnauthorized)

	if err != nil {
		return err
	}

	return nil
}

func FormAuth(r *http.Request, f func(username, password string) error) (string, error) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	err := f(username, password)
	if err != nil {
		return "", err
	}
	return username, nil
}
