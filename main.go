package main

import (
	_config "golang-gin/src/config"
	_ports "golang-gin/src/ports"
)

func main() {

	switch _config.SERVER_TYPE {
	case "http":
		_ports.CreateGinServer()
		break
	default:
		_ports.CreateGinServer()
	}

}
