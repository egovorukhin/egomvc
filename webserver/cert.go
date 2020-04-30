package webserver

import (
	"os"
	"path/filepath"
)

type Certificate struct {
	Path string `json:"path" xml:"path"`
	Cert string `json:"cert" xml:"cert"`
	Key  string `json:"key" xml:"key"`
}

/*
var certificate *Certificate

func SetCertificate(cert Certificate) {
	certificate = &cert
}

func GetCertificate() *Certificate {
	return certificate
}
*/
func (c *Certificate) Check() error {

	app, err := os.Executable()
	if err != nil {
		return err
	}

	c.Path = filepath.Join(filepath.Dir(app), c.Path)

	//Проверяем директорию
	if _, err := os.Stat(c.Path); os.IsNotExist(err) {
		return err
	}

	//Проверяем наличие cert.pem
	if _, err := os.Stat(filepath.Join(c.Path, c.Cert)); os.IsNotExist(err) {
		return err
	}

	//Проверяем наличие key.pem
	if _, err := os.Stat(filepath.Join(c.Path, c.Key)); os.IsNotExist(err) {
		return err
	}

	return nil
}
