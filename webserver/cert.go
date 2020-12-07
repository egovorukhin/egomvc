package webserver

import (
	"github.com/egovorukhin/egologger"
	"os"
	"path/filepath"
)

type Certificate struct {
	Path string `yaml:"path"`
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
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

	log := egologger.New(c.Check, logFilename)

	app, err := os.Executable()
	if err != nil {
		log.Error(err)
		return err
	}

	c.Path = filepath.Join(filepath.Dir(app), c.Path)

	//Проверяем директорию
	if _, err := os.Stat(c.Path); os.IsNotExist(err) {
		log.Error(err)
		return err
	}

	//Проверяем наличие cert.pem
	if _, err := os.Stat(filepath.Join(c.Path, c.Cert)); os.IsNotExist(err) {
		log.Error(err)
		return err
	}

	//Проверяем наличие key.pem
	if _, err := os.Stat(filepath.Join(c.Path, c.Key)); os.IsNotExist(err) {
		log.Error(err)
		return err
	}

	return nil
}
