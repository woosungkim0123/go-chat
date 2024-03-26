package config

import (
	"ws/internal/util/jsonReader"
)

type serverInfo struct {
	Host string
	Port string
}

type configuration struct {
	ServerInfo serverInfo
}

var Configuration = configuration{}

func init() {
	jsonReader.ReadAndConvert("./config/config.json", &Configuration)
}
