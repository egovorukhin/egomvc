package egomvc

import (
	"github.com/egovorukhin/egomvc/controllers"
	"github.com/egovorukhin/egomvc/webserver"
	"testing"
)

func Test(t *testing.T) {
	//webserver.Init(controllers.Init())
	webserver.SetSessionSaveFunc(func(session webserver.Session) error {
		return nil
	})
	webserver.SetSessionLoadFunc(func(session *webserver.Session) error {
		return nil
	})
	webserver.InitTest(5, controllers.Init())
}
