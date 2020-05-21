package egomvc

import (
	"github.com/egovorukhin/egomvc/controllers"
	"github.com/egovorukhin/egomvc/webserver"
	"testing"
)

func TestInit(t *testing.T) {

	controllers.Init()

	webserver.InitTest(5)
}
