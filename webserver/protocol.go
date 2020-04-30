package webserver

import (
	"context"
	"fmt"
	"net/http"
)

type Protocol struct {
	Enabled bool    `json:"enabled" xml:"enabled"`
	Port    string  `json:"port" xml:"port"`
	Timeout Timeout `json:"timeout" xml:"timeout"`
	Started bool
	Server  *http.Server
}

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
