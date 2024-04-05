package config

import (
	"github.com/gorilla/sessions"
	"ws/internal/common/jsonReader"
)

type serverConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type sessionConfig struct {
	EncryptionKey string `json:"encryptionKey"`
	MaxAge        int    `json:"maxAge"`
}

type configuration struct {
	ServerConfig  serverConfig  `json:"serverConfig"`
	SessionConfig sessionConfig `json:"sessionConfig"`
}

var Configuration = configuration{}

var Store *sessions.CookieStore

func init() {
	jsonReader.ReadAndConvert("./config/config.json", &Configuration)
	setSessionStore()
}

func setSessionStore() {
	Store = sessions.NewCookieStore([]byte(Configuration.SessionConfig.EncryptionKey))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   Configuration.SessionConfig.MaxAge,
		HttpOnly: true,
	}
}
