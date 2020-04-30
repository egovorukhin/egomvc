package webserver

import "fmt"

const (
	START   = "start"
	STOP    = "stop"
	RESTART = "restart"
	EXIT    = "exit"
	HELP    = "help"
	CONFIG  = "cfg"
)

func help() string {
	res := START + "- запуск web сервера\n"
	res += STOP + " - остановка web сервера\n"
	res += RESTART + " - перезапуск web сервера\n"
	res += CONFIG + " - конфигурация\n"
	res += EXIT + " - остановка web сервера и выход\n"
	return res
}

func getConfig() string {
	return fmt.Sprintf("http: {%s}\nhttps: {%s}",
		Protocol(ws.Http).String(),
		Protocol(ws.Https).String(),
	)
}
