package config

import (
	"github.com/gorilla/sessions"
	"ws/internal/util/jsonReader"
)

type serverInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type sessionInfo struct {
	EncryptionKey string `json:"encryptionKey"`
	MaxAge        int    `json:"maxAge"`
}

type configuration struct {
	ServerInfo  serverInfo  `json:"serverInfo"`
	SessionInfo sessionInfo `json:"sessionInfo"`
}

var Configuration = configuration{}

var Store *sessions.CookieStore

func init() {
	jsonReader.ReadAndConvert("./config/config.json", &Configuration)
	setSessionStore()
}

func setSessionStore() {
	Store = sessions.NewCookieStore([]byte(Configuration.SessionInfo.EncryptionKey))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   Configuration.SessionInfo.MaxAge,
		HttpOnly: true,
	}
}
