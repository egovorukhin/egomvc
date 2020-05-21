package database

import "errors"

func Authorization(username, password string) error {
	if username == "yegor" && password == "Qq123456" {
		return nil
	}
	return errors.New("Не верное имя пользователя или пароль!")
}

func Verify(id string) bool {
	return true
}
