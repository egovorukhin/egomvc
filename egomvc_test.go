package egomvc

import (
	"github.com/egovorukhin/egomvc/webserver"
	"testing"
)

func TestInit(t *testing.T) {
	err := webserver.Init()
	if err != nil {
		t.Error(err)
	}
}