package webserver

import (
	"context"
	"fmt"
	"net/http"
)

type Protocol struct {
	Enabled     bool        `yaml:"enabled"`
	Port        string      `yaml:"port"`
	Timeout     Timeout     `yaml:"timeout"`
	Certificate Certificate `yaml:"certificate"`
	Started     bool
	Server      *http.Server
	Controllers Controllers
}

//map с интерфейсами Router
type Controllers map[string]*Controller

func (p Protocol) Shutdown() string {

	if p.Server == nil {
		return "нет данных"
	}

	err := p.Server.Shutdown(context.TODO())
	if err != nil {
		return err.Error()
	}

	return "успешно"
}

func (p Protocol) String() string {
	return fmt.Sprintf(`"enabled": %t", "port": %s, "timeout.read": %d, "timeout.write": %d, "started": %t`,
		p.Enabled,
		p.Port,
		p.Timeout.Read,
		p.Timeout.Write,
		p.Started,
	)
}

func (controllers Controllers) add(controller *Controller) {
	controllers[controller.Name] = controller
}
