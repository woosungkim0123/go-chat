package config

import (
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"net/http"
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

var UpgradeConnection = websocket.Upgrader{}

func init() {
	jsonReader.ReadAndConvert("internal/config/config.json", &Configuration)
	setSessionStore()
	setSocketUpgradeConnection("http://localhost:8080")
}

func setSessionStore() {
	Store = sessions.NewCookieStore([]byte(Configuration.SessionConfig.EncryptionKey))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   Configuration.SessionConfig.MaxAge,
		HttpOnly: true,
	}
}

func setSocketUpgradeConnection(allowOrigin string) {
	UpgradeConnection = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == allowOrigin {
				return true
			}
			return false
		},
	}
}
